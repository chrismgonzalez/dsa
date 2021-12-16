# Arrays and Slices

## Overview

Arrays and slices in Go are two fundamental data structures of the language. How are they the same? They both consist of
an ordered sequence of elements, you can keep data together that belongs together. Yet, there are significant distances
of which to note about the two structures. If you are approaching this reading from a traditional computer science
background. An array in Go is analogous to the static array data structure, whereas the slice shares its likeness with
the dynamic array. Let's discuss the major areas of importance to understand about both of these fundamental data
structures.

### Arrays

As mentioned above, Go arrays are static arrays. Their allocation of space in memory only occurs once. This idea
provides a bit of rigidity in how you use them, but that tradeoff can see speed and performance benefits. Especially if
you are working with data of a fixed size and want to perform operations in place on the structure. All elements of the
Go array must be of the same type (i.e. `int`, `String`)
and must follow the pattern, `[capacity]data_type{element_values}`.

Example instantiation:

```go
var newArray = [5]int{1, 2, 3, 4, 5} // [1 2 3 4 5]

//if you declare now values, the array is instantiated with 0 values like below
var array = [3]int{}
// [0 0 0]
```

Arrays, like every other index based data structure in other common programming languages are 0 indexed. Ergo, in 5
element array about we would see something like this when trying to access the array elements by their index

```go
newArray[0] // 1
newArray[1] // 2
newArray[2] // 3
newArray[3] // 4
newArray[4] // 5
```

To alter the value at an array index, we can do so like in other languages. However, your value must be of the same type
as what was defined at instantiation

```go
newArray[0] = 7

fmt.printLn(newArray[0]) // 7
```

Unlike Slices, we cannot use the append() method. Why? Well, the append() under the hood performs a copy of the current
data and assigns it to a new array (allocating a new spot in memory). This would violate the static property that the
array has. If you know you will be adding many elements to your array like structure, then you should most definitely
use a slice.

### Slices

#### References

- [Understanding Arrays and Slices in Go](https://www.digitalocean.com/community/tutorials/understanding-arrays-and-slices-in-go)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)