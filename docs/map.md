---
layout: page
title: Map
permalink: /map/
nav_order: 10

---



## Map

We have used maps in the range section but now we dive deeper into what they are. Maps are a built-in data structure in Go which allows you to store and retrieve key-value pairs. In order programming languages, They are known as associative arrays(PHP), dictionaries(Python) or hashmaps(Java). Maps provide a way to store unordered data collection, each element being identified by a unique key. maps are a robust data structure that can store and retrieve data quickly and efficiently. The map's keys must be unique, while the value can be of any type.

### Creating a map
To create a map, we use the make function or the map literal syntax to create it.
1. Using the make function. 
`m:= make(map[keyType]valueType)`
2. Using the literal syntax.
`m:= map[keyType]valueType{}`

In the above syntax definition, the keyType is the type of keys in the map and will always have to be unique, while the valueType is the type of the value and can be of any supported datatype in go.  Eg if we want to create a map with keys to be strings and values to be integers, we can use the syntax below
```go
// using make function to create a map
var list = make(map[string]int)
list["John"] = 12
list["Mary"] = 24
list["Alice"] = 17
list["Bob"] = 25

for name, age := range list {
    fmt.Println(name, age)
}

// using the map literal syntax for creating map
var persons = map[string] int {"John": 12, "Mary": 24, "Alice": 17, "Bob": 25}

for name, age := range persons {
    fmt.Println(name, age)
}

```
*Note:* The key type Key must be comparable using ==, so that the map can test whether a given key is equal to one already within it.
Though floating-point numbers are comparable, it's a bad idea to compare floats for equality hence they are not good as keys. Value can however be of any type desired.

In the above, you will notice that we have not only declared how to create maps but also added how to add, iterate and access the map's values. To add values to a map, we assign a value to a specified key. 

```go
list["Carol"] = 30
persons["Carol"] = 30
```

Also, to access the values of a map.  we use the key inside of the square bracket as below
```go
fmt.Println(list["Carol"]) // Outputs: 30
fmt.Println(persons["Bob"]) // Outputs: 25
```
Note:

Accessing a key in a map will return the zero value for the value type if a key is unavailable. This essentially means that at any point in time if you try acessing any key on a map you will get a value. With this in mind, there is an optional second return value from accessing a map that can be used to decide whether the key exists. Map iteration is done using range, as we saw when discussing the range keyword.

```go
value, exists := persons["Carol"]
if exists {
    fmt.Println("Carol's score:", value)
} else {
    fmt.Println("Carol's score doesn't exist")
}
```
Deleting from a map

When deleting from a map, we use the delete function  eg.
```go
delete(persons, "Carol")
```
In case we call delete on a key that is not present in the map. nothing happens generally termed no-op hence it's safe to call delete even on keys that don't exist. question is why will you want to do that?

#### Exercise

1. Create a map of countries to their capitals.
2. Create a map of colors to their hexadecimal values.
3. Create a map of words to their definitions.
4.  Write a function in Go that takes a map of student names and their scores as input and returns the average score.  Refer to functions section above for help.

Check  Miscellanous for more

