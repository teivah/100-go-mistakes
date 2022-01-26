# 100 Go Mistakes and How to Avoid Them

Source code of [100 Go Mistakes and How to Avoid Them](https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them) 📖, edited by Manning.

## Table of Contents

### Chapter 1 - Introduction

### Chapter 2 - Code and Project Organization

* 1 - Unintended variable shadowing
* 2 - Writing nested code
* 3 - Misusing init functions
* 4 - Always using getters and setters
* 5 - Interface pollution
* 6 - Interface on the producer side
* 7 - Returning interfaces
* 8 - `any` says nothing
* 9 - [Being confused about when to use generics](https://teivah.medium.com/when-to-use-generics-in-go-36d49c1aeda)
* 10 - Not being aware of the possible problems with type embedding
* 11 - Not using the functional options pattern
* 12 - Project misorganization (project structure and package organization)
* 13 - Creating utility packages
* 14 - Ignoring package name collisions
* 15 - Missing code documentation
* 16 - Not using linters

### Chapter 3 - Data Types
  
* 17 - Creating confusion with octal literals
* 18 - Neglecting integer overflows
* 19 - Not understanding floating-points
* 20 - Not understanding slice length and capacity
* 21 - Inefficient slice initialization
* 22 - Being confused about nil vs. empty slice
* 23 - Not properly checking if a slice is empty
* 24 - Not making slice copy correctly
* 25 - Unexpected side-effects using slice append
* 26 - Slice and memory leaks
* 27 - Inefficient map initialization
* 28 - Map and memory leaks
* 29 - Comparing values incorrectly

### Chapter 4 - Control Structures  

* 30 - Ignoring that elements are copied in `range` loops
* 31 - Ignoring how arguments are evaluated in `range` loops (channels and arrays)
* 32 - Ignoring the impacts of using pointer elements in `range` loops
* 33 - Making wrong assumptions during map iterations (ordering and map insert during iteration)
* 34 - Ignoring how the `break` statement work
* 35 - Using `defer` inside a loop

### Chapter 5 - Strings
  
* 36 - Not understanding the concept of rune
* 37 - Inaccurate string iteration
* 38 - Misusing trim functions
* 39 - Under-optimized strings concatenation
* 40 - Useless string conversion
* 41 - Substring and memory leaks

### Chapter 6 - Functions and Methods
  
* 42 - Not knowing which type of receiver to use
* 43 - Never using named result parameters
* 44 - Unintended side-effects with named result parameters
* 45 - Returning a nil receiver
* 46 - Using a filename as a function input
* 47 - Ignoring how defer arguments and receivers are evaluated
  
### Chapter 7 - Error Management

* 48 - Panicking
* 49 - Ignoring when to wrap an error
* 50 - Comparing an error type inaccurately
* 51 - Comparing an error value inaccurately
* 52 - Handling an error twice
* 53 - Not handling an error
* 54 - Not handling defer errors

### Chapter 8 - Concurrency: Foundations
  
* 55 - Mixing concurrency and parallelism
* 56 - Concurrency isn't always faster
* 57 - Being puzzled about when to use channels or mutexes
* 58 - Not understanding race problems (data races vs. race conditions and Go memory model)
* 59 - Not understanding the concurrency impacts of a workload type
* 60 - Misunderstanding Go contexts

### Chapter 9 - Concurrency: Practice

* 61 - Propagating an inappropriate context
* 62 - Starting a goroutine without knowing when to stop it
* 63 - Not being careful with goroutines and loop variables
* 64 - Expecting a deterministic behavior using select and channels
* 65 - Not using notification channels
* 66 - Not using nil channels
* 67 - Being puzzled about a channel size
* 68 - Forgetting about possible side-effects with string formatting (etcd and dead-lock)
* 69 - Creating data races with append
* 70 - Using mutexes inaccurately with slices and maps
* 71 - Misusing `sync.WaitGroup`
* 72 - Forgetting about `sync.Cond`
* 73 - Not using `errgroup`
* 74 - Copying a `sync` type

### Chapter 10 - Standard Library

* 75 - Providing a wrong time duration
* 76 - `time.After` and memory leak
* 77 - JSON handling common mistakes
    * Unexpected behavior because of type embedding
    * JSON and monotonic clock
    * Map of `any`
* 78 - SQL common mistakes
    * Forgetting that `sql.Open` doesn't necessarily establish connections to a DB
    * Forgetting about connections pooling
    * Not using prepared statements
    * Mishandling null values
    * Not handling rows iteration errors
* 79 - Not closing transient resources (HTTP body, `sql.Rows`, and `os.File`)
* 80 - Forgetting the return statement after replying to an HTTP request
* 81 - Using the default HTTP client and server

### Chapter 11 - Testing

* 82 - Not categorizing tests (build tags and short mode)
* 83 - Not enabling the race flag
* 84 - Not using test execution modes (parallel and shuffle)
* 85 - Not using table-driven tests
* 86 - Sleeping in unit tests
* 87 - Not dealing with the time API efficiently
* 88 - Not using testing utility packages (`httptest` and `iotest`)
* 89 - Writing inaccurate benchmarks
    * Not resetting or pausing the timer
    * Making wrong assumptions about micro-benchmarks
    * Not being careful about compiler optimizations
    * Being fooled by the observer effect
* 90 - Not exploring all the Go testing features
    * Code coverage
    * Testing from a different package
    * Utility functions
    * Setup and teardown

### Chapter 12 - Optimizations

* 91 - Not understanding CPU caches
    * CPU architecture
    * Cache line
    * Slice of structs vs. struct of slices
    * Predictability
    * Cache placement policy
* 92 - Writing concurrent code leading to false sharing
* 93 - Not taking into account instruction-level parallelism
* 94 - Not being aware of data alignment
* 95 - Not understanding stack vs. heap
* 96 - Not knowing how to reduce allocations
    * API change
    * Compiler optimizations
    * `sync.Pool`
* 97 - Not relying on inlining
* 98 - Not using Go diagnostics tooling (profiling and execution tracer)
* 99 - Not understanding how the GC works
* 100 - Not understanding the impacts of running Go inside of Docker and Kubernetes
