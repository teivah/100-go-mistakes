---
title: Writing concurrent code that leads to false sharing (#92)
comments: true
hide:
- toc
---

# Writing concurrent code that leads to false sharing

In previous sections, we have discussed the fundamental concepts of CPU caching. We have seen that some specific caches (typically, L1 and L2) aren’t shared among all the logical cores but are specific to a physical core. This specificity has some concrete impacts such as concurrency and the concept of false sharing, which can lead to a significant performance decrease. Let’s look at what false sharing is via an example and then see how to prevent it.

In this example, we use two structs, `Input` and `Result`:

```go
type Input struct {
    a int64
    b int64
}

type Result struct {
    sumA int64
    sumB int64
}
```

The goal is to implement a `count` function that receives a slice of `Input` and computes the following:

* The sum of all the `Input.a` fields into `Result.sumA`
* The sum of all the `Input.b` fields into `Result.sumB`

For the sake of the example, we implement a concurrent solution with one goroutine that computes `sumA` and another that computes `sumB`:

```go
func count(inputs []Input) Result {
    wg := sync.WaitGroup{}
    wg.Add(2)

    result := Result{} // Init the result struct

    go func() {
        for i := 0; i < len(inputs); i++ {
            result.sumA += inputs[i].a // Computes sumA
        }
        wg.Done()
    }()

    go func() {
        for i := 0; i < len(inputs); i++ {
            result.sumB += inputs[i].b // Computes sumB
        }
        wg.Done()
    }()

    wg.Wait()
    return result
}
```

We spin up two goroutines: one that iterates over each a field and another that iterates over each b field. This example is fine from a concurrency perspective. For instance, it doesn’t lead to a data race, because each goroutine increments its own variable. But this example illustrates the false sharing concept that degrades expected performance.

Let’s look at the main memory. Because `sumA` and `sumB` are allocated contiguously, in most cases (seven out of eight), both variables are allocated to the same memory block:

<figure markdown>
  ![](img/false-sharing-1.svg)
  <figcaption>In this example, sumA and sumB are part of the same memory block.</figcaption>
</figure>


Now, let’s assume that the machine contains two cores. In most cases, we should eventually have two threads scheduled on different cores. So if the CPU decides to copy this memory block to a cache line, it is copied twice:

<figure markdown>
  ![](img/false-sharing-2.svg)
  <figcaption>Each block is copied to a cache line on both code 0 and core 1.</figcaption>
</figure>

Both cache lines are replicated because L1D (L1 data) is per core. Recall that in our example, each goroutine updates its own variable: `sumA` on one side, and `sumB` on the other side:

<figure markdown>
  ![](img/false-sharing-3.svg)
  <figcaption>Each goroutine updates its own variable.</figcaption>
</figure>

Because these cache lines are replicated, one of the goals of the CPU is to guarantee cache coherency. For example, if one goroutine updates `sumA` and another reads `sumA` (after some synchronization), we expect our application to get the latest value.

However, our example doesn’t do exactly this. Both goroutines access their own variables, not a shared one. We might expect the CPU to know about this and understand that it isn’t a conflict, but this isn’t the case. When we write a variable that’s in a cache, the granularity tracked by the CPU isn’t the variable: it’s the cache line.

When a cache line is shared across multiple cores and at least one goroutine is a writer, the entire cache line is invalidated. This happens even if the updates are logically independent (for example, `sumA` and `sumB`). This is the problem of false sharing, and it degrades performance.

???+ note

    Internally, a CPU uses the [MESI protocol](https://en.wikipedia.org/wiki/MESI_protocol) to guarantee cache coherency. It tracks each cache line, marking it modified, exclusive, shared, or invalid (MESI).

One of the most important aspects to understand about memory and caching is that sharing memory across cores isn’t real—it’s an illusion. This understanding comes from the fact that we don’t consider a machine a black box; instead, we try to have mechanical sympathy with underlying levels.

So how do we solve false sharing? There are two main solutions.

The first solution is to use the same approach we’ve shown but ensure that `sumA` and `sumB` aren’t part of the same cache line. For example, we can update the `Result` struct to add _padding_ between the fields. Padding is a technique to allocate extra memory. Because an `int64` requires an 8-byte allocation and a cache line 64 bytes long, we need 64 – 8 = 56 bytes of padding:

```go hl_lines="3"
type Result struct {
    sumA int64
    _    [56]byte // Padding
    sumB int64
}
```

The next figure shows a possible memory allocation. Using padding, `sumA` and `sumB` will always be part of different memory blocks and hence different cache lines.

<figure markdown>
  ![](img/false-sharing-4.svg)
  <figcaption>sumA and sumB are part of different memory blocks.</figcaption>
</figure>

If we benchmark both solutions (with and without padding), we see that the padding solution is significantly faster (about 40% on my machine). This is an important improvement that results from the addition of padding between the two fields to prevent false sharing.

The second solution is to rework the structure of the algorithm. For example, instead of having both goroutines share the same struct, we can make them communicate their local result via channels. The result benchmark is roughly the same as with padding.

In summary, we must remember that sharing memory across goroutines is an illusion at the lowest memory levels. False sharing occurs when a cache line is shared across two cores when at least one goroutine is a writer. If we need to optimize an application that relies on concurrency, we should check whether false sharing applies, because this pattern is known to degrade application performance. We can prevent false sharing with either padding or communication.
