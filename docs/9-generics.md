---
title: Being confused about when to use generics (#9)
comments: true
hide:
- toc
---

# Being confused about when to use generics

Generics is a fresh addition to the language. In a nutshell, it allows writing code with types that can be specified later and instantiated when needed. However, it can be pretty easy to be confused about when to use generics and when not to. Throughout this post, we will describe the concept of generics in Go and then delve into common use and misuses.

## Concepts

Consider the following function that extracts all the keys from a `map[string]int` type:

```go
func getKeys(m map[string]int) []string {
    var keys []string
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}
```

What if we would like to use a similar feature for another map type such as a `map[int]string`? Before generics, Go developers had a couple of options: using code generation, reflection, or duplicating code.

For example, we could write two functions, one for each map type, or even try to extend `getKeys` to accept different map types:

```go
func getKeys(m any) ([]any, error) {
    switch t := m.(type) {
    default:
        return nil, fmt.Errorf("unknown type: %T", t)
    case map[string]int:
        var keys []any
        for k := range t {
            keys = append(keys, k)
        }
        return keys, nil
    case map[int]string:
        // Copy the extraction logic
    }
}
```

We can start noticing a couple of issues:

* First, it increases boilerplate code. Indeed, whenever we want to add a case, it will require duplicating the `range` loop.
*  Meanwhile, the function now accepts an empty interface, which means we are losing some of the benefits of Go being a typed language. Indeed, checking whether a type is supported is done at runtime instead of compile-time. Hence, we also need to return an error if the provided type is unknown.
* Last but not least, as the key type can be either `int` or `string`, we are obliged to return a slice of empty interfaces to factor out key types. This approach increases the effort on the caller-side as the client may also have to perform a type check of the keys or extra conversion.

Thanks to generics, we can now refactor this code using type parameters.

Type parameters are generic types we can use with functions and types. For example, the following function accepts a type parameter:

```go
func foo[T any](t T) {
    // ...
}
```

When calling `foo`, we will pass a type argument of any type. Passing a type argument is called instantiation because the work is done at compile time which keeps type safety as part of the core language features and avoids runtime overheads.

Let’s get back to the `getKeys` function and use type parameters to write a generic version that would accept any kind of map:

```go
func getKeys[K comparable, V any](m map[K]V) []K {
  var keys []K <2>
  for k := range m {
    keys = append(keys, k)
  }
  return keys
}
```

To handle the map, we defined two kinds of type parameters. First, the values can be of any type: `V any`. However, in Go, the map keys can’t be of any type. For example, we cannot use slices:

```go
var m map[[]byte]int
```

This code leads to a compilation error: `invalid map key type []byte`. Therefore, instead of accepting any key type, we are obliged to restrict type arguments so that the key type meets specific requirements. Here, being comparable (we can use `==` or `!=`). Hence, we defined `K` as `comparable` instead of `any`.

Restricting type arguments to match specific requirements is called a constraint. A constraint is an interface type that can contain:

* A set of behaviors (methods)
* But also arbitrary type

Let’s see a concrete example for the latter. Imagine we don’t want to accept any `comparable` type for map key type. For instance, we would like to restrict it to either `int` or `string` types. We can define a custom constraint this way:

```go
type customConstraint interface {
   ~int | ~string // Define a custom type that will restrict types to int and string
}

// Change the type parameter K to be custom
func getKeys[K customConstraint, V any](m map[K]V) []K {
   // Same implementation
}
```

First, we define a `customConstraint` interface to restrict the types to be either `int` or `string` using the union operator `|` (we will discuss the use of `~` a bit later). Then, `K` is now a `customConstraint` instead of a `comparable` as before.

Now, the signature of `getKeys` enforces that we can call it with a map of any value type, but the key type has to be an `int` or a `string`. For example, on the caller-side:

```go
m = map[string]int{
   "one":   1,
   "two":   2,
   "three": 3,
}
keys := getKeys(m)
```

Note that Go can infer that `getKeys` is called with a `string` type argument. The previous call was similar to this:

```go
keys := getKeys[string](m)
```

???+ note

    What’s the difference between a constraint using `~int` or `int`? Using `int` restricts it to that type, whereas `~int` restricts all the types whose underlying type is an `int`.

    To illustrate it, let’s imagine a constraint where we would like to restrict a type to any `int` type implementing the `String() string` method:

    ```go
    type customConstraint interface {
       ~int
       String() string
    }
    ```

    Using this constraint will restrict type arguments to custom types like this one:

    ```go
    type customInt int

    func (i customInt) String() string {
       return strconv.Itoa(int(i))
    }
    ```

    As `customInt` is an `int` and implements the `String() string` method, the `customInt` type satisfies the constraint defined.

    However, if we change the constraint to contain an `int` instead of an `~int`, using `customInt` would lead to a compilation error because the `int` type doesn’t implement `String() string`.

Let’s also note the `constraints` package contains a set of common constraints such as `Signed` that includes all the signed integer types. Let’s ensure that a constraint doesn’t already exist in this package before creating a new one.

So far, we have discussed examples using generics for functions. However, we can also use generics with data structures.

For example, we will create a linked list containing values of any type. Meanwhile, we will write an `Add` method to append a node:

```go
type Node[T any] struct { // Use type parameter
   Val  T
   next *Node[T]
}

func (n *Node[T]) Add(next *Node[T]) { // Instantiate type receiver
   n.next = next
}
```

We use type parameters to define `T` and use both fields in `Node`. Regarding the method, the receiver is instantiated. Indeed, because `Node` is generic, it has to follow also the type parameter defined.

One last thing to note about type parameters: they can’t be used on methods, only on functions. For example, the following method wouldn’t compile:

```go
type Foo struct {}

func (Foo) bar[T any](t T) {}
```

```
./main.go:29:15: methods cannot have type parameters
```

Now, let’s delve into concrete cases where we should and shouldn’t use generics.

## Common uses and misuses

So when are generics useful? Let’s discuss a couple of common uses where generics **are** recommended:

* Data structures. For example, we can use generics to factor out the element type if we implement a binary tree, a linked list, or a heap.
* Functions working with slices, maps, and channels of any type. For example, a function to merge two channels would work with any channel type. Hence, we could use type parameters to factor out the channel type:

  ```go
  func merge[T any](ch1, ch2 <-chan T) <-chan T {
      // ...
  }
  ```

* Meanwhile, instead of factoring out a type, we can factor out behaviors. For example, the `sort` package contains functions to sort different slice types such as `sort.Ints` or `sort.Float64s`. Using type parameters, we can factor out the sorting behaviors that rely on three methods, `Len`, `Less`, and `Swap`:

  ```go
  type sliceFn[T any] struct { // Use type parameter
     s       []T
     compare func(T, T) bool // Compare two T elements
  }

  func (s sliceFn[T]) Len() int           { return len(s.s) }
  func (s sliceFn[T]) Less(i, j int) bool { return s.compare(s.s[i], s.s[j]) }
  func (s sliceFn[T]) Swap(i, j int)      { s.s[i], s.s[j] = s.s[j], s.s[i] }
  ```

Conversely, when is it recommended **not** to use generics?

* When just calling a method of the type argument. For example, consider a function that receives an `io.Writer` and call the `Write` method:

  ```go
  func foo[T io.Writer](w T) {
     b := getBytes()
     _, _ = w.Write(b)
  }
  ```

* When it makes our code more complex. Generics are never mandatory, and as Go developers, we have been able to live without them for more than a decade. If writing generic functions or structures we figure out that it doesn’t make our code clearer, we should probably reconsider our decision for this particular use case.

## Conclusion

Though generics can be very helpful in particular conditions, we should be cautious about when to use them and not use them.

In general, when we want to answer when not to use generics, we can find similarities with when not to use interfaces. Indeed, generics introduce a form of abstraction, and we have to remember that unnecessary abstractions introduce complexity.

Let’s not pollute our code with needless abstractions, and let’s focus on solving concrete problems for now. It means that we shouldn’t use type parameters prematurely. Let’s wait until we are about to write boilerplate code to consider using generics.

