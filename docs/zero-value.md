---
layout: page
title: Zero Value
permalink: /zero-value/
nav_order: 14

---

## Zero value
Note: You have been hearing us made mention of zero value through out the document particularlyw when we talked about types and data types.  Now we explain what a zero value is about. 
The zero value

When storage is allocated for a variable, either through a declaration or a call of new, or when a new value is created, either through a composite literal or a call of make, and no explicit initialization is provided, the variable or value is given a default value. Each element of such a variable or value is set to the zero value for its type: false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps. This initialization is done recursively, so for instance each element of an array of structs will have its fields zeroed if no value is specified.

These two simple declarations are equivalent:
```go
var i int
var i int = 0
```

```go
type T struct { i int; f float64; next *T }
t := new(T)

```
the following holds:
```go
t.i == 0
t.f == 0.0
t.next == nil
```
The same would also be true after

```go
var t T
```

[Reference Zero value](https://go.dev/ref/spec#The_zero_value)


