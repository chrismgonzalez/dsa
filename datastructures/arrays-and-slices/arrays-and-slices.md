# Arrays and Slices

## Overview

Arrays and slices in Go are two fundamental data structures of the language. How are they the same? They both consist of
an ordered sequence of elements, you can keep data together that belongs together. Yet, there are significant distances
of which to note about the two structures. If you are approaching this reading from a traditional computer science
background. An array in Go is analogous to the static array data structure, whereas the slice shares its likeness with
the dynamic array. Let's discuss the major areas of importance to understand about both of these fundamental data
structures.

### Arrays

As mentioned above, Go arrays are static arrays. Their allocation of space in memory only occurs once. It is an
immutable data structure in the sense that once it allocates a specific block of spaces in memory, it cannot change. You
can change the values within the array, just not the size. This idea provides a bit of rigidity in how you use them, but
that tradeoff can see speed and performance benefits. Especially if you are working with data of a fixed size and want
to perform operations in place on the structure. All elements of the Go array must be of the same type (i.e. `int`
, `String`) and must follow the pattern, `[capacity]data_type{element_values}`.

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

As mentioned in the overview above, slices are the dynamic array structure in Go. It is a mutable data structure meaning
that you can change the size at anytime by appending elements to it. Under the hood the append() method will copy the
values in the current slice and assign them to a new array that is double the size in order to have enough space. You'll
want to work with a slice data type when you want to quickly iterate over elements and operate on them.

Slices are defined similar to arrays, except they do not have to declare an allocation amount within their square
brackets. Like so, `[]data_type{value}`

As far as instantiation goes, this can be done several ways:

- assign the slice to a variable with no values
- assigning the slice to a variable and defining its values
- using the built-in make function when you want to create a slice of length n
- use the built-in make function to pre-allocate the capacity

```go
// assign the slice to a variable with no values (ideal if you want to define the slice as an auxiliar data structure in your algorithm)
emptySlice := []int{}

//assigning the slice to a variable and defining its values
sliceWithValues := []string{"Hello", "World", "Nice", "to", "see", "you"}

//using the built-in make function when you want to create a slice of length n, the first argument is the type, and second argument is the length n
sliceWithMake := make([]string, 3)

//use the built-in make function to pre-allocate the capacity.  Below we are creating a zeroed slice of 3 elements, with the option to expand it to 5 elements
preAllocateSlice := make([]int, 3, 5)
```

How can we create a slice from an array? It's rather simple. You can do this operation by what we call "slicing" the
array. Simply put, declare the indices which you would like to slice given the format `[first_index:second_index]`.
Remember that arrays and slices are zero indexed. When you do slice an array, the result is a slice, not another array.
It is helpful to think of this operation in the following terms:

- when you declare your indices, keep in mind that the first value is the starting index, and the second value is the
  starting index(inclusive) + ending index(exclusive)
- if you want to start at the first index or end at the last index, you can leave the value off before or after the
  colon, respectively

Examples:

```go
// using the slices above we will perform some operations
sliceWithValues := []string{"Hello", "World", "Nice", "to", "see", "you"}

fmt.printLn(sliceWithValues[1:3]) // [World Nice] this prints index 1 and up to, but not including index 3.  This can also be understood by sliceWithValues[starting_index: (starting_index + how many elements you want returned)
fmt.printLn(sliceWithValues[:4]) // [Hello World Nice To]
fmt.printLn(sliceWithValues[1:]) // [World Nice to see you] // starts at index 1 and prints the rest of the slice including the last index


```

What if you want to convert an array in to a slice? Well, this is simple as well. All you need to do is invoke the same
command as above with only a colon between the square brackets

```go
slicedArray := newArray[:] // [1 2 3 4 5]
```

Now a new value can be appended to this slice, which we could not do when it was an array. We can also append more than
one element at a time, or even another slice altogether.

```go
//appending a single value
slicedArray = append(slicedArray, 6) // [1 2 3 4 5 6]

//appending multiple values
slicedArray = append(slicedArray, 6, 7, 8) //[1 2 3 4 5 6 7 8]

//append another slice
sliceToAdd := []int{9, 10, 11}
slicedArray = append(slicedArray, sliceToAdd) // [1 2 3 4 5 6 7 8 9 10 11]
```

We can also measure the capacity of a slice, at least, the capacity that has been allocated for it by using the cap()
method. This helps us fill the slice programmatically, and also reduces the memory allocation as we do it.

```go
// using append will allocate more memory as we add values to the slice
numbers := []int{}
for i := 0; i < 4; i++ {
numbers = append(numbers, i)
}
fmt.Println(numbers)

// if you have an idea of how large you want your slice to be, go ahead and pre-allocate the memory, then add values iteratively using cap()
numbers := make([]int, 4)
for i := 0; i < cap(numbers); i++ {
numbers[i] = i
}

fmt.Println(numbers)

The second example above stores values within the already pre-allocated space of memory that the make() function used right above the loop declaration.
```

#### References

- [Understanding Arrays and Slices in Go](https://www.digitalocean.com/community/tutorials/understanding-arrays-and-slices-in-go)
- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)