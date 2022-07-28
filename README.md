# 100 Go Mistakes and How to Avoid Them

Source code of 📖 [100 Go Mistakes and How to Avoid Them](https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them), edited by Manning.

![](cover.png)

## Quotes

> "This should be the required reading for all Golang developers before they touch code in Production... It's the Golang equivalent of the legendary 'Effective Java' by Joshua Bloch."

– Neeraj Shah

> "This unique book teaches you good habits by helping you identify bad ones. Harsanyi's writing style is engaging, the examples relevant, and his insights useful. I thought it was a great read, and I think you will too."

– Thad Meyer

> "Learning from mistakes is proven as one of the best ways to learn a subject. This book helps you do just that by demonstrating the most common mistakes people make coming to Go, why most people make them and the proper way to solve the problems."

– Ryan Huber

> "This book explains many subtleties of the Go programming language that may cause errors and provides the reader with advice on how to deal with these situations. The precise explanations and real world examples make it a great addition for those learning Go programming language or looking to advance their mastery of the language."

– Borko Djurkovic

> "Not having this will be the 101st mistake a Go programmer could make."

– Anupam Sengupta

## Table of Contents

### Chapter 1 - Introduction

### Chapter 2 - Code and Project Organization

#### #1: Unintended variable shadowing

Avoiding shadowed variables can help prevent mistakes like referencing the wrong variable or confusing readers.

#### #2: Unnecessary nested code

Avoiding nested levels and keeping the happy path aligned on the left makes building a mental code model easier.

#### #3: Misusing init functions

When initializing variables, remember that init functions have limited error handling and make state handling and testing more complex. In most cases, initializations should be handled as specific functions.

#### #4: Overusing getters and setters

Forcing the use of getters and setters isn’t idiomatic in Go. Being pragmatic and finding the right balance between efficiency and blindly following certain idioms should be the way to go.

#### #5: Interface pollution

Abstractions should be discovered, not created. To prevent unnecessary complexity, create an interface when you need it and not when you foresee needing it, or if you can at least prove the abstraction to be a valid one.

#### #6: Interface on the producer side

Keeping interfaces on the client side avoids unnecessary abstractions.

#### #7: Returning interfaces

To prevent being restricted in terms of flexibility, a function shouldn’t return interfaces but concrete implementations in most cases. Conversely, a function should accept interfaces whenever possible.

#### #8: `any` says nothing

Only use `any` if you need to accept or return any possible type, such as `json.Marshal`. Otherwise, `any` doesn’t provide meaningful information and can lead to compile-time issues by allowing a caller to call methods with any data type.

#### #9: [Being confused about when to use generics](https://teivah.medium.com/when-to-use-generics-in-go-36d49c1aeda)

Relying on generics and type parameters can prevent writing boilerplate code to factor out elements or behaviors. However, do not use type parameters prematurely, but only when you see a concrete need for them. Otherwise, they introduce unnecessary abstractions and complexity.

#### #10: Not being aware of the possible problems with type embedding

Using type embedding can also help avoid boilerplate code; however, ensure that doing so doesn’t lead to visibility issues where some fields should have remained hidden.

#### #11: Not using the functional options pattern

To handle options conveniently and in an API-friendly manner, use the func- tional options pattern.

#### #12: Project misorganization

* Project structure

* Package organization

Following a layout such as project-layout can be a good way to start structuring Go projects, especially if you are looking for existing conventions to standardize a new project.

#### #13: Creating utility packages

Naming is a critical piece of application design. Creating packages such as `common`, `util`, and `shared` doesn’t bring much value for the reader. Refactor such packages into meaningful and specific package names.

#### #14: Ignoring package name collisions

To avoid naming collisions between variables and packages, leading to confusion or perhaps even bugs, use unique names for each one. If this isn’t feasible, use an import alias to change the qualifier to differentiate the package name from the variable name, or think of a better name.

#### #15: Missing code documentation

To help clients and maintainers understand your code’s purpose, document exported elements.

#### #16: Not using linters

To improve code quality and consistency, use linters and formatters.

### Chapter 3 - Data Types
  
#### #17: Creating confusion with octal literals

When reading existing code, bear in mind that integer literals starting with 0 are octal numbers. Also, to improve readability, make octal integers explicit by prefixing them with `0o`.

#### #18: Neglecting integer overflows

Because integer overflows and underflows are handled silently in Go, you can implement your own functions to catch them.

#### #19: Not understanding floating-points

Making floating-point comparisons within a given delta can ensure that your code is portable. 

When performing addition or subtraction, group the operations with a similar order of magnitude to favor accuracy. Also, perform multiplication and division before addition and subtraction.

#### #20: Not understanding slice length and capacity

Understanding the difference between slice length and capacity should be part of a Go developer’s core knowledge. The slice length is the number of available elements in the slice, whereas the slice capacity is the number of elements in the backing array.

#### #21: Inefficient slice initialization

When creating a slice, initialize it with a given length or capacity if its length is already known. This reduces the number of allocations and improves performance. The same logic goes for maps, and you need to initialize their size.

#### #22: Being confused about nil vs. empty slice

To prevent common confusions such as when using the `encoding/json` or the `reflect` package, you need to understand the difference between nil and empty slices. Both are zero-length, zero-capacity slices, but only a nil slice doesn’t require allocation.

#### #23: Not properly checking if a slice is empty

To check if a slice doesn’t contain any element, check its length. This check works regardless of whether the slice is `nil` or empty. The same goes for maps.

To design unambiguous APIs, you shouldn’t distinguish between nil and empty slices.

#### #24: Not making slice copies correctly

To copy one slice to another using the `copy` built-in function, remember that the number of copied elements corresponds to the minimum between the two slice’s lengths.

#### #25: Unexpected side effects using slice append

Using copy or the full slice expression is a way to prevent `append` from creating conflicts if two different functions use slices backed by the same array. However, only a slice copy prevents memory leaks if you want to shrink a large slice.

#### #26: Slice and memory leaks

Working with a slice of pointers or structs with pointer fields, you can avoid memory leaks by marking as nil the elements excluded by a slicing operation.

#### #27: Inefficient map initialization

See [#22](#21-inefficient-slice-initialization).

#### #28: Map and memory leaks

A map can always grow in memory, but it never shrinks. Hence, if it leads to some memory issues, you can try different options, such as forcing Go to recreate the map or using pointers.

#### #29: Comparing values incorrectly

To compare types in Go, you can use the == and != operators if two types are comparable: Booleans, numerals, strings, pointers, channels, and structs are composed entirely of comparable types. Otherwise, you can either use `reflect.DeepEqual` and pay the price of reflection or use custom implementations and libraries.

### Chapter 4 - Control Structures  

#### #30: Ignoring that elements are copied in `range` loops
#### #31: Ignoring how arguments are evaluated in `range` loops (channels and arrays)
#### #32: Ignoring the impacts of using pointer elements in `range` loops
#### #33: Making wrong assumptions during map iterations (ordering and map insert during iteration)
#### #34: Ignoring how the `break` statement work
#### #35: Using `defer` inside a loop

### Chapter 5 - Strings
  
#### #36: Not understanding the concept of rune
#### #37: Inaccurate string iteration
#### #38: Misusing trim functions
#### #39: Under-optimized strings concatenation
#### #40: Useless string conversions
#### #41: Substring and memory leaks

### Chapter 6 - Functions and Methods
  
#### #42: Not knowing which type of receiver to use
#### #43: Never using named result parameters
#### #44: Unintended side effects with named result parameters
#### #45: Returning a nil receiver
#### #46: Using a filename as a function input
#### #47: Ignoring how defer arguments and receivers are evaluated (argument evaluation, pointer, and value receivers)
  
### Chapter 7 - Error Management

#### #48: Panicking
#### #49: Ignoring when to wrap an error
#### #50: Comparing an error type inaccurately
#### #51: Comparing an error value inaccurately
#### #52: Handling an error twice
#### #53: Not handling an error
#### #54: Not handling defer errors

### Chapter 8 - Concurrency: Foundations
  
#### #55: Mixing up concurrency and parallelism
#### #56: Thinking concurrency is always faster
#### #57: Being puzzled about when to use channels or mutexes
#### #58: Not understanding race problems (data races vs. race conditions and the Go memory model)
#### #59: Not understanding the concurrency impacts of a workload type
#### #60: Misunderstanding Go contexts

### Chapter 9 - Concurrency: Practice

#### #61: Propagating an inappropriate context
#### #62: Starting a goroutine without knowing when to stop it
#### #63: Not being careful with goroutines and loop variables
#### #64: Expecting a deterministic behavior using select and channels
#### #65: Not using notification channels
#### #66: Not using nil channels
#### #67: Being puzzled about channel size
#### #68: Forgetting about possible side effects with string formatting (etcd data race example and deadlock)
#### #69: Creating data races with append
#### #70: Using mutexes inaccurately with slices and maps
#### #71: Misusing `sync.WaitGroup`
#### #72: Forgetting about `sync.Cond`
#### #73: Not using `errgroup`
#### #74: Copying a `sync` type

### Chapter 10 - Standard Library

#### #75: Providing a wrong time duration
#### #76: `time.After` and memory leaks
#### #77: JSON handling common mistakes
    * Unexpected behavior because of type embedding
    * JSON and the monotonic clock
    * Map of `any`
#### #78: Common SQL mistakes
    * Forgetting that `sql.Open` doesn't necessarily establish connections to a database
    * Forgetting about connections pooling
    * Not using prepared statements
    * Mishandling null values
    * Not handling rows iteration errors
#### #79: Not closing transient resources (HTTP body, `sql.Rows`, and `os.File`)
#### #80: Forgetting the return statement after replying to an HTTP request
#### #81: Using the default HTTP client and server

### Chapter 11 - Testing

#### #82: Not categorizing tests (build tags, environment variables, and short mode)
#### #83: Not enabling the race flag
#### #84: Not using test execution modes (parallel and shuffle)
#### #85: Not using table-driven tests
#### #86: Sleeping in unit tests
#### #87: Not dealing with the time API efficiently
#### #88: Not using testing utility packages (`httptest` and `iotest`)
#### #89: Writing inaccurate benchmarks
    * Not resetting or pausing the timer
    * Making wrong assumptions about micro-benchmarks
    * Not being careful about compiler optimizations
    * Being fooled by the observer effect
#### #90: Not exploring all the Go testing features
    * Code coverage
    * Testing from a different package
    * Utility functions
    * Setup and teardown

### Chapter 12 - Optimizations

#### #91: Not understanding CPU caches
    * CPU architecture
    * Cache line
    * Slice of structs vs. struct of slices
    * Predictability
    * Cache placement policy
#### #92: Writing concurrent code that leads to false sharing
#### #93: Not taking into account instruction-level parallelism
#### #94: Not being aware of data alignment
#### #95: Not understanding stack vs. heap
#### #96: Not knowing how to reduce allocations
    * API change
    * Compiler optimizations
    * `sync.Pool`
#### #97: Not relying on inlining
#### #98: Not using Go diagnostics tooling
    * Profiling (enabling pprof, CPU, heap, goroutines, block, and mutex profiling)
    * Execution tracer
#### #99: Not understanding how the GC works
#### #100: Not understanding the impacts of running Go in Docker and Kubernetes

![](inside-cover.jpg)

## Author

Teiva Harsanyi is a senior software engineer at Docker. He worked in various domains, including insurance, transportation, and safety-critical industries like air traffic management. He is passionate about Go and how to design and implement reliable applications.

**Note:** If you're struggling to afford the book, please DM me [@teivah](https://twitter.com/teivah).
