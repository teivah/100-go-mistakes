# 100 Go Mistakes and How to Avoid Them

Source code and community space of 📖 [100 Go Mistakes and How to Avoid Them](https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them) published by Manning in October 2022. Made by [@teivah](https://twitter.com/teivah).

## Sponsors

_Want your company name to appear here? Go to [❤️ Sponsor](https://github.com/sponsors/teivah), One-time tab, "Sponsor 100 Go Mistakes and How to Avoid Them" tier_.

## 📖 100 Go Mistakes and How to Avoid Them

![](banner.png)

100 Go Mistakes and How to Avoid Them shows you how to replace common programming problems in Go with idiomatic, expressive code. In it, you’ll explore dozens of interesting examples and case studies as you learn to spot mistakes that might appear in your own applications.

### Where to Buy?

* [Manning.com](https://www.manning.com/books/100-go-mistakes-and-how-to-avoid-them)
* Amazon: [.com](https://www.amazon.com/dp/1617299596), [.co.uk](https://www.amazon.co.uk/dp/B0BBSNJR6B), [.de](https://www.amazon.de/dp/B0BBHQD8BQ), [.fr](https://www.amazon.fr/100-Mistakes-How-Avoid-Them/dp/1617299596), [.in](https://www.amazon.in/dp/B0BBHQD8BQ), [.co.jp](https://www.amazon.co.jp/dp/B0BBHQD8BQ), [.es](https://www.amazon.es/dp/B0BBHQD8BQ), [.it](https://www.amazon.it/dp/B0BBHQD8BQ), [.com.br](https://www.amazon.com.br/dp/B0BBHQD8BQ)

## Common Go Mistakes

![](inside-cover.png)

This section contains a summary of the 100 mistakes in the book. Meanwhile, it's also a section open to the community. If you believe that a mistake should be added, please create a [community mistake issue](https://github.com/teivah/100-go-mistakes/issues/new?assignees=&labels=community+mistake&template=community_mistake.md&title=).

* [Code and Project Organization](#code-and-project-organization)
  * [Unintended variable shadowing](#unintended-variable-shadowing-1)
  * [Unnecessary nested code](#unnecessary-nested-code-2)
  * [Misusing init functions](#misusing-init-functions-3)
  * [Overusing getters and setters](#overusing-getters-and-setters-4)
  * [Interface pollution](#interface-pollution-5)
  * [Interface on the producer side](#interface-on-the-producer-side-6)
  * [Returning interfaces](#returning-interfaces-7)
  * [`any` says nothing](#any-says-nothing-8)
  * [Being confused about when to use generics](#being-confused-about-when-to-use-genericshttpsteivahmediumcomwhen-to-use-generics-in-go-36d49c1aeda-9) ([Read the full excerpt 👀](https://teivah.medium.com/when-to-use-generics-in-go-36d49c1aeda))
  * [Not being aware of the possible problems with type embedding](#not-being-aware-of-the-possible-problems-with-type-embedding-10)
  * [Not using the functional options pattern](#not-using-the-functional-options-pattern-11)
  * [Project misorganization](#project-misorganization-project-structure-and-package-organization-12)
  * [Creating utility packages](#creating-utility-packages-13)
  * [Ignoring package name collisions](#ignoring-package-name-collisions-14)
  * [Missing code documentation](#missing-code-documentation-15)
  * [Not using linters](#not-using-linters-16)
* [Data Types](#data-types)
  * [Creating confusion with octal literals](#creating-confusion-with-octal-literals-17)
  * [Neglecting integer overflows](#neglecting-integer-overflows-18)
  * [Not understanding floating-points](#not-understanding-floating-points-19)
  * [Not understanding slice length and capacity](#not-understanding-slice-length-and-capacityhttpsteivahmediumcomslice-length-vs-capacity-in-go-af71a754b7d8-20) ([Read the full excerpt 👀](https://teivah.medium.com/slice-length-vs-capacity-in-go-af71a754b7d8))
  * [Inefficient slice initialization](#inefficient-slice-initialization-21)
  * [Being confused about nil vs. empty slice](#being-confused-about-nil-vs-empty-slice-22)
  * [Not properly checking if a slice is empty](#not-properly-checking-if-a-slice-is-empty-23)
  * [Not making slice copies correctly](#not-making-slice-copies-correctly-24)
  * [Unexpected side effects using slice append](#unexpected-side-effects-using-slice-append-25)
  * [Slice and memory leaks](#slice-and-memory-leaks-26)
  * [Inefficient map initialization](#inefficient-map-initialization-27)
  * [Map and memory leaks](#map-and-memory-leakshttpsteivahmediumcommaps-and-memory-leaks-in-go-a85ebe6e7e69-28) ([Read the full excerpt 👀](https://teivah.medium.com/maps-and-memory-leaks-in-go-a85ebe6e7e69))
  * [Comparing values incorrectly](#comparing-values-incorrectly-29)
* [Control Structures](#control-structures)
  * [Ignoring that elements are copied in `range` loops](#ignoring-that-elements-are-copied-in-range-loops-30)
  * [Ignoring how arguments are evaluated in `range` loops](#ignoring-how-arguments-are-evaluated-in-range-loops-channels-and-arrays-31)
  * [Ignoring the impacts of using pointer elements in `range` loops](#ignoring-the-impacts-of-using-pointer-elements-in-range-loops-32)
  * [Making wrong assumptions during map iterations](#making-wrong-assumptions-during-map-iterations-ordering-and-map-insert-during-iteration-33)
  * [Ignoring how the `break` statement works](#ignoring-how-the-break-statement-works-34)
  * [Using defer inside a loop](#using-defer-inside-a-loop-35)
* [Strings](#strings)
  * [Not understanding the concept of rune](#not-understanding-the-concept-of-rune-36)
  * [Inaccurate string iteration](#inaccurate-string-iteration-37)
  * [Misusing trim functions](#misusing-trim-functions-38)
  * [Under-optimized strings concatenation](#under-optimized-strings-concatenation-39)
  * [Useless string conversions](#useless-string-conversions-40)
  * [Substring and memory leaks](#substring-and-memory-leaks-41)
* [Functions and Methods](#functions-and-methods)
  * [Not knowing which type of receiver to use](#not-knowing-which-type-of-receiver-to-use-42)
  * [Never using named result parameters](#never-using-named-result-parameters-43)
  * [Unintended side effects with named result parameters](#unintended-side-effects-with-named-result-parameters-44)
  * [Returning a nil receiver](#returning-a-nil-receiver-45)
  * [Using a filename as a function input](#using-a-filename-as-a-function-input-46)
  * [Ignoring how `defer` arguments and receivers are evaluated](#ignoring-how-defer-arguments-and-receivers-are-evaluated-argument-evaluation-pointer-and-value-receivers-47)
* [Error Management](#error-management)
  * [Panicking](#panicking-48)
  * [Ignoring when to wrap an error](#ignoring-when-to-wrap-an-error-49)
  * [Comparing an error type inaccurately](#comparing-an-error-type-inaccurately-50)
  * [Comparing an error value inaccurately](#comparing-an-error-value-inaccurately-51)
  * [Handling an error twice](#handling-an-error-twice-52)
  * [Not handling an error](#not-handling-an-error-53)
  * [Not handling `defer` errors](#not-handling-defer-errors-54)
* [Concurrency: Foundations](#concurrency-foundations)
  * [Mixing up concurrency and parallelism](#mixing-up-concurrency-and-parallelism-55)
  * [Thinking concurrency is always faster](#thinking-concurrency-is-always-fasterhttpsteivahmediumcomconcurrency-isnt-always-faster-in-go-de325168907c-56)
  * [Thinking concurrency is always faster](#thinking-concurrency-is-always-fasterhttpsteivahmediumcomconcurrency-isnt-always-faster-in-go-de325168907c-56) ([Read the full excerpt 👀](https://teivah.medium.com/concurrency-isnt-always-faster-in-go-de325168907c))
  * [Being puzzled about when to use channels or mutexes](#being-puzzled-about-when-to-use-channels-or-mutexes-57)
  * [Not understanding race problems](#not-understanding-race-problems-data-races-vs-race-conditions-and-the-go-memory-model-58)
  * [Not understanding the concurrency impacts of a workload type](#not-understanding-the-concurrency-impacts-of-a-workload-type-59)
  * [Misunderstanding Go contexts](#misunderstanding-go-contexts-60)
* [Concurrency: Practice](#concurrency-practice)
  * [Propagating an inappropriate context](#propagating-an-inappropriate-context-61)
  * [Starting a goroutine without knowing when to stop it](#starting-a-goroutine-without-knowing-when-to-stop-it-62)
  * [Not being careful with goroutines and loop variables](#not-being-careful-with-goroutines-and-loop-variables-63)
  * [Expecting a deterministic behavior using select and channels](#expecting-a-deterministic-behavior-using-select-and-channels-64)
  * [Not using notification channels](#not-using-notification-channels-65)
  * [Not using nil channels](#not-using-nil-channels-66)
  * [Being puzzled about channel size](#being-puzzled-about-channel-size-67)
  * [Forgetting about possible side effects with string formatting](#forgetting-about-possible-side-effects-with-string-formatting-etcd-data-race-example-and-deadlock-68)
  * [Creating data races with append](#creating-data-races-with-append-69)
  * [Using mutexes inaccurately with slices and maps](#using-mutexes-inaccurately-with-slices-and-maps-70)
  * [Misusing `sync.WaitGroup`](#misusing-syncwaitgroup-71)
  * [Forgetting about `sync.Cond`](#forgetting-about-synccond-72)
  * [Not using `errgroup`](#not-using-errgroup-73)
  * [Copying a `sync` type](#not-using-errgroup-74)
* [Standard Library](#standard-library)
  * [Providing a wrong time duration](#providing-a-wrong-time-duration-75)
  * [`time.After` and memory leaks](#timeafter-and-memory-leaks-76)
  * [JSON handling common mistakes](#json-handling-common-mistakes-77)
  * [Common SQL mistakes](#common-sql-mistakes-78)
  * [Not closing transient resources](#not-closing-transient-resources-http-body-sqlrows-and-osfile-79)
  * [Forgetting the return statement after replying to an HTTP request](#forgetting-the-return-statement-after-replying-to-an-http-request-80)
  * [Using the default HTTP client and server](#using-the-default-http-client-and-server-81)
* [Testing](#testing)
  * [Not categorizing tests](#not-categorizing-tests-build-tags-environment-variables-and-short-mode-82)
  * [Not enabling the race flag](#not-enabling-the-race-flag-83)
  * [Not using test execution modes](#not-using-test-execution-modes-parallel-and-shuffle-84)
  * [Not using table-driven tests](#not-using-table-driven-tests-85)
  * [Sleeping in unit tests](#sleeping-in-unit-tests-86)
  * [Not dealing with the time API efficiently](#not-dealing-with-the-time-api-efficiently-87)
  * [Not using testing utility packages](#not-using-testing-utility-packages-httptest-and-iotest-88)
  * [Writing inaccurate benchmarks](#writing-inaccurate-benchmarkshttpsteivahmediumcomhow-to-write-accurate-benchmarks-in-go-4266d7dd1a95-89) ([Read the full excerpt 👀](https://teivah.medium.com/how-to-write-accurate-benchmarks-in-go-4266d7dd1a95))
  * [Not exploring all the Go testing features](#not-exploring-all-the-go-testing-features-90)
* [Optimizations](#optimizations)
  * [Not understanding CPU caches](#not-understanding-cpu-caches-91)
  * [Writing concurrent code that leads to false sharing](#writing-concurrent-code-that-leads-to-false-sharing-92)
  * [Not taking into account instruction-level parallelism](#not-taking-into-account-instruction-level-parallelism-93)
  * [Not being aware of data alignment](#not-being-aware-of-data-alignment-94)
  * [Not understanding stack vs. heap](#not-understanding-stack-vs-heap-95)
  * [Not knowing how to reduce allocations](#not-knowing-how-to-reduce-allocations-api-change-compiler-optimizations-and-syncpool-96)
  * [Not relying on inlining](#not-relying-on-inlining-97)
  * [Not using Go diagnostics tooling](#not-using-go-diagnostics-tooling-profiling-enabling-pprof-cpu-heap-goroutines-block-and-mutex-profiling-and-execution-tracer-98)
  * [Not understanding how the GC works](#not-understanding-how-the-gc-works-99)
  * [Not understanding the impacts of running Go in Docker and Kubernetes](#not-understanding-the-impacts-of-running-go-in-docker-and-kubernetes-100)

### Code and Project Organization

#### Unintended variable shadowing (#1)

Avoiding shadowed variables can help prevent mistakes like referencing the wrong variable or confusing readers.

#### Unnecessary nested code (#2)

Avoiding nested levels and keeping the happy path aligned on the left makes building a mental code model easier.

#### Misusing init functions (#3)

When initializing variables, remember that init functions have limited error handling and make state handling and testing more complex. In most cases, initializations should be handled as specific functions.

#### Overusing getters and setters (#4)

Forcing the use of getters and setters isn’t idiomatic in Go. Being pragmatic and finding the right balance between efficiency and blindly following certain idioms should be the way to go.

#### Interface pollution (#5)

Abstractions should be discovered, not created. To prevent unnecessary complexity, create an interface when you need it and not when you foresee needing it, or if you can at least prove the abstraction to be a valid one.

#### Interface on the producer side (#6)

Keeping interfaces on the client side avoids unnecessary abstractions.

#### Returning interfaces (#7)

To prevent being restricted in terms of flexibility, a function shouldn’t return interfaces but concrete implementations in most cases. Conversely, a function should accept interfaces whenever possible.

#### `any` says nothing (#8)

Only use `any` if you need to accept or return any possible type, such as `json.Marshal`. Otherwise, `any` doesn’t provide meaningful information and can lead to compile-time issues by allowing a caller to call methods with any data type.

#### [Being confused about when to use generics](https://teivah.medium.com/when-to-use-generics-in-go-36d49c1aeda) (#9)

Relying on generics and type parameters can prevent writing boilerplate code to factor out elements or behaviors. However, do not use type parameters prematurely, but only when you see a concrete need for them. Otherwise, they introduce unnecessary abstractions and complexity.

#### Not being aware of the possible problems with type embedding (#10)

Using type embedding can also help avoid boilerplate code; however, ensure that doing so doesn’t lead to visibility issues where some fields should have remained hidden.

#### Not using the functional options pattern (#11)

To handle options conveniently and in an API-friendly manner, use the functional options pattern.

#### Project misorganization (project structure and package organization) (#12)

Following a layout such as project-layout can be a good way to start structuring Go projects, especially if you are looking for existing conventions to standardize a new project.

#### Creating utility packages (#13)

Naming is a critical piece of application design. Creating packages such as `common`, `util`, and `shared` doesn’t bring much value for the reader. Refactor such packages into meaningful and specific package names.

#### Ignoring package name collisions (#14)

To avoid naming collisions between variables and packages, leading to confusion or perhaps even bugs, use unique names for each one. If this isn’t feasible, use an import alias to change the qualifier to differentiate the package name from the variable name, or think of a better name.

#### Missing code documentation (#15)

To help clients and maintainers understand your code’s purpose, document exported elements.

#### Not using linters (#16)

To improve code quality and consistency, use linters and formatters.

### Data Types

#### Creating confusion with octal literals (#17)

When reading existing code, bear in mind that integer literals starting with 0 are octal numbers. Also, to improve readability, make octal integers explicit by prefixing them with `0o`.

#### Neglecting integer overflows (#18)

Because integer overflows and underflows are handled silently in Go, you can implement your own functions to catch them.

#### Not understanding floating-points (#19)

Making floating-point comparisons within a given delta can ensure that your code is portable.

When performing addition or subtraction, group the operations with a similar order of magnitude to favor accuracy. Also, perform multiplication and division before addition and subtraction.

#### [Not understanding slice length and capacity](https://teivah.medium.com/slice-length-vs-capacity-in-go-af71a754b7d8) (#20)

Understanding the difference between slice length and capacity should be part of a Go developer’s core knowledge. The slice length is the number of available elements in the slice, whereas the slice capacity is the number of elements in the backing array.

#### Inefficient slice initialization (#21)

When creating a slice, initialize it with a given length or capacity if its length is already known. This reduces the number of allocations and improves performance. The same logic goes for maps, and you need to initialize their size.

#### Being confused about nil vs. empty slice (#22)

To prevent common confusions such as when using the `encoding/json` or the `reflect` package, you need to understand the difference between nil and empty slices. Both are zero-length, zero-capacity slices, but only a nil slice doesn’t require allocation.

#### Not properly checking if a slice is empty (#23)

To check if a slice doesn’t contain any element, check its length. This check works regardless of whether the slice is `nil` or empty. The same goes for maps.

To design unambiguous APIs, you shouldn’t distinguish between nil and empty slices.

#### Not making slice copies correctly (#24)

To copy one slice to another using the `copy` built-in function, remember that the number of copied elements corresponds to the minimum between the two slice’s lengths.

#### Unexpected side effects using slice append (#25)

Using copy or the full slice expression is a way to prevent `append` from creating conflicts if two different functions use slices backed by the same array. However, only a slice copy prevents memory leaks if you want to shrink a large slice.

#### Slice and memory leaks (#26)

Working with a slice of pointers or structs with pointer fields, you can avoid memory leaks by marking as nil the elements excluded by a slicing operation.

#### Inefficient map initialization (#27)

See [#21](#inefficient-slice-initialization-21).

#### [Map and memory leaks](https://teivah.medium.com/maps-and-memory-leaks-in-go-a85ebe6e7e69) (#28)

A map can always grow in memory, but it never shrinks. Hence, if it leads to some memory issues, you can try different options, such as forcing Go to recreate the map or using pointers.

#### Comparing values incorrectly (#29)

To compare types in Go, you can use the == and != operators if two types are comparable: Booleans, numerals, strings, pointers, channels, and structs are composed entirely of comparable types. Otherwise, you can either use `reflect.DeepEqual` and pay the price of reflection or use custom implementations and libraries.

### Control Structures

#### Ignoring that elements are copied in `range` loops (#30)

The value element in a `range` loop is a copy. Therefore, to mutate a struct, for example, access it via its index or via a classic `for` loop (unless the element or the field you want to modify is a pointer).

#### Ignoring how arguments are evaluated in `range` loops (channels and arrays) (#31)

Understanding that the expression passed to the `range` operator is evaluated only once before the beginning of the loop can help you avoid common mistakes such as inefficient assignment in channel or slice iteration.

#### Ignoring the impacts of using pointer elements in `range` loops (#32)

Using a local variable or accessing an element using an index, you can prevent mistakes while copying pointers inside a loop.

#### Making wrong assumptions during map iterations (ordering and map insert during iteration) (#33)

To ensure predictable outputs when using maps, remember that a map data structure:
* Doesn’t order the data by keys
* Doesn’t preserve the insertion order
* Doesn’t have a deterministic iteration order
* Doesn’t guarantee that an element added during an iteration will be produced during this iteration

#### Ignoring how the `break` statement works (#34)

Using `break` or `continue` with a label enforces breaking a specific statement. This can be helpful with `switch` or `select` statements inside loops.

#### Using `defer` inside a loop (#35)

Extracting loop logic inside a function leads to executing a `defer` statement at the end of each iteration.

### Strings

#### Not understanding the concept of rune (#36)

Understanding that a rune corresponds to the concept of a Unicode code point and that it can be composed of multiple bytes should be part of the Go developer’s core knowledge to work accurately with strings.

#### Inaccurate string iteration (#37)

Iterating on a string with the `range` operator iterates on the runes with the index corresponding to the starting index of the rune’s byte sequence. To access a specific rune index (such as the third rune), convert the string into a `[]rune`.

#### Misusing trim functions (#38)

`strings.TrimRight`/`strings.TrimLeft` removes all the trailing/leading runes contained in a given set, whereas `strings.TrimSuffix`/`strings.TrimPrefix` returns a string without a provided suffix/prefix.

#### Under-optimized strings concatenation (#39)

Concatenating a list of strings should be done with `strings.Builder` to prevent allocating a new string during each iteration.

#### Useless string conversions (#40)

Remembering that the `bytes` package offers the same operations as the `strings` package can help avoid extra byte/string conversions.

#### Substring and memory leaks (#41)

Using copies instead of substrings can prevent memory leaks, as the string returned by a substring operation will be backed by the same byte array.

### Functions and Methods

#### Not knowing which type of receiver to use (#42)

The decision whether to use a value or a pointer receiver should be made based on factors such as the type, whether it has to be mutated, whether it contains a field that can’t be copied, and how large the object is. When in doubt, use a pointer receiver.

#### Never using named result parameters (#43)

Using named result parameters can be an efficient way to improve the readability of a function/method, especially if multiple result parameters have the same type. In some cases, this approach can also be convenient because named result parameters are initialized to their zero value. But be cautious about potential side effects.

#### Unintended side effects with named result parameters (#44)

See [#43](#never-using-named-result-parameters-43).

#### Returning a nil receiver (#45)

When returning an interface, be cautious about returning not a nil pointer but an explicit nil value. Otherwise, unintended consequences may result because the caller will receive a non-nil value.

#### Using a filename as a function input (#46)

Designing functions to receive `io.Reader` types instead of filenames improves the reusability of a function and makes testing easier.

#### Ignoring how `defer` arguments and receivers are evaluated (argument evaluation, pointer, and value receivers) (#47)

Passing a pointer to a `defer` function and wrapping a call inside a closure are two possible solutions to overcome the immediate evaluation of arguments and receivers.

### Error Management

#### Panicking (#48)

Using `panic` is an option to deal with errors in Go. However, it should only be used sparingly in unrecoverable conditions: for example, to signal a programmer error or when you fail to load a mandatory dependency.

#### Ignoring when to wrap an error (#49)

Wrapping an error allows you to mark an error and/or provide additional context. However, error wrapping creates potential coupling as it makes the source error available for the caller. If you want to prevent that, don’t use error wrapping.

#### Comparing an error type inaccurately (#50)

If you use Go 1.13 error wrapping with the `%w` directive and `fmt.Errorf`, comparing an error against a type or a value has to be done using `errors.As` or `errors.Is`, respectively. Otherwise, if the returned error you want to check is wrapped, it will fail the checks.

#### Comparing an error value inaccurately (#51)

See [#50](#comparing-an-error-type-inaccurately-50).

To convey an expected error, use error sentinels (error values). An unexpected error should be a specific error type.

#### Handling an error twice (#52)

In most situations, an error should be handled only once. Logging an error is handling an error. Therefore, you have to choose between logging or returning an error. In many cases, error wrapping is the solution as it allows you to provide additional context to an error and return the source error.

#### Not handling an error (#53)

Ignoring an error, whether during a function call or in a `defer` function, should be done explicitly using the blank identifier. Otherwise, future readers may be confused about whether it was intentional or a miss.

#### Not handling `defer` errors (#54)

In many cases, you shouldn’t ignore an error returned by a `defer` function. Either handle it directly or propagate it to the caller, depending on the context. If you want to ignore it, use the blank identifier.

### Concurrency: Foundations

#### Mixing up concurrency and parallelism (#55)

Understanding the fundamental differences between concurrency and parallelism is a cornerstone of the Go developer’s knowledge. Concurrency is about structure, whereas parallelism is about execution.

#### [Thinking concurrency is always faster](https://teivah.medium.com/concurrency-isnt-always-faster-in-go-de325168907c) (#56)

To be a proficient developer, you must acknowledge that concurrency isn’t always faster. Solutions involving parallelization of minimal workloads may not necessarily be faster than a sequential implementation. Benchmarking sequential versus concurrent solutions should be the way to validate assumptions.

#### Being puzzled about when to use channels or mutexes (#57)

Being aware of goroutine interactions can also be helpful when deciding between channels and mutexes. In general, parallel goroutines require synchronization and hence mutexes. Conversely, concurrent goroutines generally require coordination and orchestration and hence channels.

#### Not understanding race problems (data races vs. race conditions and the Go memory model) (#58)

Being proficient in concurrency also means understanding that data races and race conditions are different concepts. Data races occur when multiple goroutines simultaneously access the same memory location and at least one of them is writing. Meanwhile, being data-race-free doesn’t necessarily mean deterministic execution. When a behavior depends on the sequence or the timing of events that can’t be controlled, this is a race condition.

Understanding the Go memory model and the underlying guarantees in terms of ordering and synchronization is essential to prevent possible data races and/or race conditions.

#### Not understanding the concurrency impacts of a workload type (#59)

When creating a certain number of goroutines, consider the workload type. Creating CPU-bound goroutines means bounding this number close to the `GOMAXPROCS` variable (based by default on the number of CPU cores on the host). Creating I/O-bound goroutines depends on other factors, such as the external system.

#### Misunderstanding Go contexts (#60)

Go contexts are also one of the cornerstones of concurrency in Go. A context allows you to carry a deadline, a cancellation signal, and/or a list of keys-values.

### Concurrency: Practice

#### Propagating an inappropriate context (#61)

Understanding the conditions when a context can be canceled should matter when propagating it: for example, an HTTP handler canceling the context when the response has been sent.

#### Starting a goroutine without knowing when to stop it (#62)

Avoiding leaks means being mindful that whenever a goroutine is started, you should have a plan to stop it eventually.

#### Not being careful with goroutines and loop variables (#63)

To avoid bugs with goroutines and loop variables, create local variables or call functions instead of closures.

#### Expecting a deterministic behavior using select and channels (#64)

Understanding that `select` with multiple channels chooses the case randomly if multiple options are possible prevents making wrong assumptions that can lead to subtle concurrency bugs.

#### Not using notification channels (#65)

Send notifications using a `chan struct{}` type.

#### Not using nil channels (#66)

Using nil channels should be part of your concurrency toolset because it allows you to _remove_ cases from `select` statements, for example.

#### Being puzzled about channel size (#67)

Carefully decide on the right channel type to use, given a problem. Only unbuffered channels provide strong synchronization guarantees.

You should have a good reason to specify a channel size other than one for buffered channels.

#### Forgetting about possible side effects with string formatting (etcd data race example and deadlock) (#68)

Being aware that string formatting may lead to calling existing functions means watching out for possible deadlocks and other data races.

#### Creating data races with append (#69)

Calling `append` isn’t always data-race-free; hence, it shouldn’t be used concurrently on a shared slice.

#### Using mutexes inaccurately with slices and maps (#70)

Remembering that slices and maps are pointers can prevent common data races.

#### Misusing `sync.WaitGroup` (#71)

To accurately use `sync.WaitGroup`, call the `Add` method before spinning up goroutines.

#### Forgetting about `sync.Cond` (#72)

You can send repeated notifications to multiple goroutines with `sync.Cond`.

#### Not using `errgroup` (#73)

You can synchronize a group of goroutines and handle errors and contexts with the `errgroup` package.

#### Copying a `sync` type (#74)

`sync` types shouldn’t be copied.

### Standard Library

#### Providing a wrong time duration (#75)

Remain cautious with functions accepting a `time.Duration`. Even though passing an integer is allowed, strive to use the time API to prevent any possible confusion.

#### `time.After` and memory leaks (#76)

Avoiding calls to `time.After` in repeated functions (such as loops or HTTP handlers) can avoid peak memory consumption. The resources created by `time.After` are released only when the timer expires.

#### JSON handling common mistakes (#77)

* Unexpected behavior because of type embedding

  Be careful about using embedded fields in Go structs. Doing so may lead to sneaky bugs like an embedded time.Time field implementing the `json.Marshaler` interface, hence overriding the default marshaling behavior.

* JSON and the monotonic clock

  When comparing two `time.Time` structs, recall that `time.Time` contains both a wall clock and a monotonic clock, and the comparison using the == operator is done on both clocks.

* Map of `any`

  To avoid wrong assumptions when you provide a map while unmarshaling JSON data, remember that numerics are converted to `float64` by default.

#### Common SQL mistakes (#78)

* Forgetting that `sql.Open` doesn't necessarily establish connections to a database

  Call the `Ping` or `PingContext` method if you need to test your configuration and make sure a database is reachable.

* Forgetting about connections pooling

  Configure the database connection parameters for production-grade applications.

* Not using prepared statements

  Using SQL prepared statements makes queries more efficient and more secure.

* Mishandling null values

  Deal with nullable columns in tables using pointers or `sql.NullXXX` types.

* Not handling rows iteration errors

  Call the `Err` method of `sql.Rows` after row iterations to ensure that you haven’t missed an error while preparing the next row.

#### Not closing transient resources (HTTP body, `sql.Rows`, and `os.File`) (#79)

Eventually close all structs implementing `io.Closer` to avoid possible leaks.

#### Forgetting the return statement after replying to an HTTP request (#80)

To avoid unexpected behaviors in HTTP handler implementations, make sure you don’t miss the `return` statement if you want a handler to stop after `http.Error`.

#### Using the default HTTP client and server (#81)

For production-grade applications, don’t use the default HTTP client and server implementations. These implementations are missing timeouts and behaviors that should be mandatory in production.

### Testing

#### Not categorizing tests (build tags, environment variables, and short mode) (#82)

Categorizing tests using build flags, environment variables, or short mode makes the testing process more efficient. You can create test categories using build flags or environment variables (for example, unit versus integration tests) and differentiate short from long-running tests to decide which kinds of tests to execute.

#### Not enabling the race flag (#83)

Enabling the `-race` flag is highly recommended when writing concurrent applications. Doing so allows you to catch potential data races that can lead to software bugs.

#### Not using test execution modes (parallel and shuffle) (#84)

Using the `-parallel` flag is an efficient way to speed up tests, especially long-running ones.

Use the `-shuffle` flag to help ensure that a test suite doesn’t rely on wrong assumptions that could hide bugs.

#### Not using table-driven tests (#85)

Table-driven tests are an efficient way to group a set of similar tests to prevent code duplication and make future updates easier to handle.

#### Sleeping in unit tests (#86)

Avoid sleeps using synchronization to make a test less flaky and more robust. If synchronization isn’t possible, consider a retry approach.

#### Not dealing with the time API efficiently (#87)

Understanding how to deal with functions using the time API is another way to make a test less flaky. You can use standard techniques such as handling the time as part of a hidden dependency or asking clients to provide it.

#### Not using testing utility packages (`httptest` and `iotest`) (#88)

The `httptest` package is helpful for dealing with HTTP applications. It provides a set of utilities to test both clients and servers.

The `iotest` package helps write io.Reader and test that an application is tolerant to errors.

#### [Writing inaccurate benchmarks](https://teivah.medium.com/how-to-write-accurate-benchmarks-in-go-4266d7dd1a95) (#89)

* Not resetting or pausing the timer

  Use time methods to preserve the accuracy of a benchmark.

* Making wrong assumptions about micro-benchmarks

  Increasing `benchtime` or using tools such as `benchstat` can be helpful when dealing with micro-benchmarks.

  Be careful with the results of a micro-benchmark if the system that ends up running the application is different from the one running the micro-benchmark.

* Not being careful about compiler optimizations

  Make sure the function under test leads to a side effect, to prevent compiler optimizations from fooling you about the benchmark results.

* Being fooled by the observer effect

  To prevent the observer effect, force a benchmark to re-create the data used by a CPU-bound function.

#### Not exploring all the Go testing features (#90)
* Code coverage

  Use code coverage with the `-coverprofile` flag to quickly see which part of the code needs more attention.

* Testing from a different package

  Place unit tests in a different package to enforce writing tests that focus on an exposed behavior, not internals.

* Utility functions

  Handling errors using the `*testing.T` variable instead of the classic `if err != nil` makes code shorter and easier to read.

* Setup and teardown

  You can use setup and teardown functions to configure a complex environment, such as in the case of integration tests.

### Optimizations

#### Not understanding CPU caches (#91)

* CPU architecture

  Understanding how to use CPU caches is important for optimizing CPU-bound applications because the L1 cache is about 50 to 100 times faster than the main memory.

* Cache line

  Being conscious of the cache line concept is critical to understanding how to organize data in data-intensive applications. A CPU doesn’t fetch memory word by word; instead, it usually copies a memory block to a 64-byte cache line. To get the most out of each individual cache line, enforce spatial locality.

* Slice of structs vs. struct of slices

* Predictability

  Making code predictable for the CPU can also be an efficient way to optimize certain functions. For example, a unit or constant stride is predictable for the CPU, but a non-unit stride (for example, a linked list) isn’t predictable.

* Cache placement policy

  To avoid a critical stride, hence utilizing only a tiny portion of the cache, be aware that caches are partitioned.

#### Writing concurrent code that leads to false sharing (#92)

Knowing that lower levels of CPU caches aren’t shared across all the cores helps avoid performance-degrading patterns such as false sharing while writing concurrency code. Sharing memory is an illusion.

#### Not taking into account instruction-level parallelism (#93)

Use instruction-level parallelism (ILP) to optimize specific parts of your code to allow a CPU to execute as many parallel instructions as possible. Identifying data hazards is one of the main steps.

#### Not being aware of data alignment (#94)

You can avoid common mistakes by remembering that in Go, basic types are aligned with their own size. For example, keep in mind that reorganizing the fields of a struct by size in descending order can lead to more compact structs (less memory allocation and potentially a better spatial locality).

#### Not understanding stack vs. heap (#95)

Understanding the fundamental differences between heap and stack should also be part of your core knowledge when optimizing a Go application. Stack allocations are almost free, whereas heap allocations are slower and rely on the GC to clean the memory.

#### Not knowing how to reduce allocations (API change, compiler optimizations, and `sync.Pool`) (#96)

Reducing allocations is also an essential aspect of optimizing a Go application. This can be done in different ways, such as designing the API carefully to prevent sharing up, understanding the common Go compiler optimizations, and using `sync.Pool`.

#### Not relying on inlining (#97)

Use the fast-path inlining technique to efficiently reduce the amortized time to call a function.

#### Not using Go diagnostics tooling (profiling [enabling pprof, CPU, heap, goroutines, block, and mutex profiling] and execution tracer) (#98)

Rely on profiling and the execution tracer to understand how an application performs and the parts to optimize.

#### Not understanding how the GC works (#99)

Understanding how to tune the GC can lead to multiple benefits such as handling sudden load increases more efficiently.

#### Not understanding the impacts of running Go in Docker and Kubernetes (#100)

To help avoid CPU throttling when deployed in Docker and Kubernetes, keep in mind that Go isn’t CFS-aware.

## About the Author

Teiva Harsanyi is a senior software engineer at Docker 🐳. He worked in various domains, including insurance, transportation, and safety-critical industries like air traffic management. He is passionate about Go and how to design and implement reliable applications.

## Book Quotes

> This should be the required reading for all Golang developers before they touch code in Production... It's the Golang equivalent of the legendary 'Effective Java' by Joshua Bloch.

– _Neeraj Shah_

> This unique book teaches you good habits by helping you identify bad ones. Harsanyi's writing style is engaging, the examples relevant, and his insights useful. I thought it was a great read, and I think you will too.

– _Thad Meyer_

> Not having this will be the 101st mistake a Go programmer could make.

– _Anupam Sengupta_

## Resources

* How to make mistakes in Go (Go Time - episode #190)
    * [Episode](https://changelog.com/gotime/190)
    * [Spotify](https://open.spotify.com/episode/0K1DImrxHCy6E7zVY4AxMZ?si=akroInsPQ1mM5B5V2tHLUw&dl_branch=1)

* 8LU - 100% Test Coverage: [YouTube](https://youtu.be/V3FBDav6wgQ?t=1210)
