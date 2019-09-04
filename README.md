# Algorithms in Go

[![Build Status](https://travis-ci.org/aswinkarthik/algorithms-in-go.svg?branch=master)](https://travis-ci.org/aswinkarthik/algorithms-in-go)

This repository is to implement various data structures and alogrithms in Go.

## Datastructures

- [Stack](https://godoc.org/github.com/aswinkarthik/algorithms-in-go/stack)
- [Queue](https://godoc.org/github.com/aswinkarthik/algorithms-in-go/queue)
- [Heap](https://godoc.org/github.com/aswinkarthik/algorithms-in-go/heap)

## Algorithms

- [Rabin-Karp Substring Search](https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm)

## Usage

### Stack

```go
func main() {
    s := stack.New()

    s.Push(6)
    s.Push(8)

    top, _ := s.Pop()
    fmt.Println(top) // Prints 8

    top, _ = s.Peek()
    fmt.Println(top) // Prints 6
}
```

For more docs on [stack](https://godoc.org/github.com/aswinkarthik/algorithms-in-go/stack)

### Queue

```go
func main() {
    q := queue.New()

    q.Enqueue(6)
    q.Enqueue(8)

    first, _ := q.Dequeue()
    fmt.Println(first) // Prints 6

    first, _ = q.First()
    fmt.Println(first) // Prints 8
}
```

For more docs on [queue](https://godoc.org/github.com/aswinkarthik/algorithms-in-go/queue)

### Heap

By default, you could create a MinHeap or MaxHeap with int64 items

```go
func main() {
    h := heap.NewMinHeapInt64()

    h.Insert(int64(5))
    h.Insert(int64(2))
    h.Insert(int64(3))

    first, _ := h.Remove()
    fmt.Println(first) // Prints 2

    second, _ := h.Remove()
    fmt.Println(second) // Prints 3

    third, _ := h.Top()
    fmt.Println(third) // Prints 5
}
```

You could also create a heap of any type. You would have to define a method that asserts how your heap property should be maintained.

```go
func main() {
    type myStruct struct {
        value int
        label string
    }

    // define how your heap property is maintained
    comparator := func(a, b interface{}) bool {
        return a.(myStruct).value < b.(myStruct).value
    }

    first := myStruct{5, "a"}
    second := myStruct{8, "b"}
    third := myStruct{3, "c"}

    h := heap.New(comparator)

    h.Insert(first)
    h.Insert(second)
    h.Insert(third)

    val, _ := h.Top()

    fmt.Println(val) // Prints {3 c}
}
```

For more docs on [heap](https://godoc.org/github.com/aswinkarthik/algorithms-in-go/heap)

## Strings

Methods offered in strings package

### Contains

This uses Rabin-Karp Substring match algorithm

```go
func main() {
	strings.Contains("source string", "ce s") // true
}
```

For more docs on [strings](https://godoc.org/github.com/aswinkarthik/algorithms-in-go/strings)

## Test locally

```bash
go test -v ./...
```
