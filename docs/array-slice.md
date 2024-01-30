---
layout: page
title: Arrays & Slice
permalink: /arrays/
nav_order: 8

---


## Arrays/ Slices
As introduced in the data type section Array, Slice are reference types. 
An array in Go is a data structure that consists of an ordered sequence of elements of the same type.

Arrays in go are fixed size, meaning that the number of elements in the array cannot be changed once it is declared.
Arrays elements are referenced using indexes, The First element of an array is usually referenced with an index starting at 0.
Arrays are declared using the `[]` syntax, followed by the type of the elements and the number of elements in the array. For example,
if we have a set of integers of the following
`1, 2, 3, 4, 5, 6, 7,8,9,10`
 the following code declares an array of 10 integers as above:
```go
var numbers = [10]int{1,2,3,4,5,6,7,8,9,10}
```

while working with arrays, you are expected to know the index before you can get the elements. say we want to get the value 5 from the array above, we will need to know the index that 5 is in the array; in this case, the value 5 is in index 4. Below is how to retrieve 5
```go
var five int  = numbers[4]
```

It is a common practice to most at time go through the array. Suppose we want to scan an array moving from the first item to the last. There are various ways in which we can do that.  One simpler way if to use the range function to loop through the array. You will see more about this in later sections. 

```go
// Print the indices and elements.
for index, value := range numbers {
    fmt.Printf("%d %d\n", index, value)
}
```

Remember that we just discussed about functions before here. One thing to not is that just like you saw under the functions sections an array by default when passed to a function will be passed by value( a copy of the array is sent to the function), however there is a way in which we can ensure that an array is not passed by value but by reference as a reference type.  In order to archieve passing an array to a function by reference we use what is called pointer(*) you will learn more about pointers in subsequent section. 
```go
// function that takes an array of 10 integers and squares each one of the numbers. notice that the original array refered to by ptr will be modified inplace
func squareNumbers(ptr *[10]int) {
	for i := range ptr {
		ptr[i] = ptr[i] * ptr[i]
	}
}

// the caller can simply do the below to pass the array by reference
squareNumbers(&numbers)

```



## Slice
A slice is a data structure similar to an array but with a variable size. Slices are declared using the `[]` syntax, followed by the type of the elements and the slice elements. For example, the following code declares a slice of integers:
```go
var ages []int
```

The ages slice is initially empty, but it can be used to store any number of integers. To add elements to a slice, you can use the append() function. For example, the following code adds the number 40 to the ages slice:

```go
ages = append(ages, 40)
```
Slices also have a length and capacity. The length of a slice is the number of elements that it currently contains. The capacity of a slice is the maximum number of elements that the slice can contain without having to reallocate memory. to get the length of a slice we use the built in `len` function. 


### Differences between Arrays and Slices

The main difference between arrays and slices is that arrays are fixed-size, while slices are variable-size. This means you cannot change the number of elements in an array after it is declared, but you can change the number of elements in a slice at any time.

Another difference between arrays and slices is that slices contain a pointer to the underlying array. This means that if you change the contents of a slice, the underlying array will also be changed.