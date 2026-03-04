# Escape to Heap

In the world of Go, "escape analysis" is the compiler’s way of deciding
whether a variable should live on the stack or the heap. Think of it as
the compiler’s memory management strategy to keep your program fast
and efficient.

## Stack vs. Heap: The Basics

To understand escaping, you first need to know the two "buckets" where
Go stores data:

The Stack: Very fast. It’s like a scratchpad for a specific
function. When the function finishes, everything on its stack is instantly
wiped away.

The Heap: A larger, shared pool of memory. It’s slower to manage
because the Garbage Collector (GC) has to periodically clean it up to
prevent memory leaks.

The Rule of Thumb: If the compiler can prove a variable isn't used
after the function returns, it stays on the stack. If it "escapes"
the function's scope, it must go to the heap.

## Why Does "Escape" Happen?

A variable "escapes to the heap" when its lifetime needs to outlive the
function that created it. Here are the most common scenarios:

### Sharing Upwards (Pointers)

If you return a pointer to a local variable, that variable can no longer
live on the stack because the stack frame is destroyed when the function
returns.

```Go
func createIncrementer() *int {
    x := 1          // x is created here
    return &x       // x "escapes" to the heap because the caller needs it
}
```

### Large Variables

The stack has a limited size. If you create a massive slice or struct,
the compiler might move it to the heap simply because it’s too big
for the stack’s "scratchpad."

### Dynamic Types (Interfaces)

Passing a variable to a function that takes an interface{} (like
fmt.Println) often triggers an escape. Since the compiler doesn't know
the concrete type at compile time, it plays it safe and puts the data
on the heap.

## Why Should You Care?

While Go manages this automatically, "heap escapes" have a performance cost:

1. CPU Overhead: Allocating on the heap is more complex than a simple
   stack pointer move.
2. GC Pressure: The more objects you have on the heap, the harder the
   Garbage Collector has to work. This can lead to "Stop the World" pauses
   that slow down your app.

## How to See It in Action

Here is the output when running `make`:

```
go build -o /tmp/bin -gcflags="-m" heap_escape.go
# command-line-arguments
./heap_escape.go:11:6: can inline NewPerson
./heap_escape.go:15:6: can inline NewPersonP
./heap_escape.go:20:16: inlining call to NewPerson
./heap_escape.go:21:13: inlining call to fmt.Println
./heap_escape.go:23:18: inlining call to NewPersonP
./heap_escape.go:24:13: inlining call to fmt.Println
./heap_escape.go:11:16: leaking param: first to result ~r0 level=0
./heap_escape.go:11:30: leaking param: last to result ~r0 level=0
./heap_escape.go:15:17: leaking param: first
./heap_escape.go:15:31: leaking param: last
./heap_escape.go:16:9: &Person{...} escapes to heap
./heap_escape.go:21:13: ... argument does not escape
./heap_escape.go:21:14: "Person:" escapes to heap
./heap_escape.go:21:25: p escapes to heap
./heap_escape.go:23:18: &Person{...} does not escape
./heap_escape.go:24:13: ... argument does not escape
./heap_escape.go:24:14: "Person2:" escapes to heap
./heap_escape.go:24:26: *pp escapes to heap
/tmp/bin/heap_escape
Person: {Anna Karenina 27}
Person2: {Ken Thompson 81}
```

## Summary Table

| Feature    | Stack Allocation               | Heap Allocation            |
|------------|--------------------------------|----------------------------|
| Speed      | Extremely Fast                 | Slower                     |
| Cleanup    | Automatic (when function ends) | Requires Garbage Collector |
| Visibility | Local to the function          | Shared across the program  |
| Management | Managed by CPU instructions    | Managed by Go Runtime      |

