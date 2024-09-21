---
layout: page
title: Operator
nav_order: 4
permalink: /syntax/operator/
parent: Syntax

---


### Operators


Operators are used to perform operations on values. Golang has operators for the following operations:
- Arithmetic operations: Addition, subtraction, multiplication, division, and modulus.

```
+    sum                    integers, floats, complex values, strings
-    difference             integers, floats, complex values
*    product                integers, floats, complex values
/    quotient               integers, floats, complex values
%    remainder              integers
```

- Comparison operations: Equal to, not equal to, less than, less than or equal to, greater than, and greater than or equal to.

```go
==    equal
!=    not equal
<     less
<=    less or equal
>     greater
>=    greater or equal
```

- Logical operations: And, or, and not.

```go
&&     AND    p && q  is  "if p then q else false"
||     OR     p || q  is  "if p then true else q"
!      NOT    !p      is  "not p"
```
- Assignment operations: Assignment, addition assignment, subtraction assignment, multiplication assignment, division assignment, and modulus assignment.

```go
const a = 2 + 3.0          // a == 5.0   (untyped floating-point constant)
const b = 15 / 4           // b == 3     (untyped integer constant)
const c = 15 / 4.0         // c == 3.75  (untyped floating-point constant)
const Θ float64 = 3/2      // Θ == 1.0   (type float64, 3/2 is integer division)
const n float64 = 3/2.     // Π == 1.5   (type float64, 3/2. is float division)

person := GithubUser{Login: "John", Name: "Doe", PublicRepos: 7}
person.Name = "Bob"         // struct field

```

- Binary Operator: 

```go
* / % << >> & &^
```
