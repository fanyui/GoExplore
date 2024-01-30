---
layout: page
title: Literal
nav_order: 3
permalink: /syntax/literal/
parent: Syntax

---

### Literals
Literals are constants that have a specific value. Golang has literals for the following types:
- Integers: Integer literals can be written in decimal, octal, or hexadecimal format.
- Floating-point numbers: Floating-point literals can be written in decimal or scientific format.
- Strings: String literals are enclosed in double quotes.
- Booleans: Boolean literals can be true or false.
- Runes: Rune literals are enclosed in single quotes.

```go
var integerLiteral = 42       // Integer literal
var floatLiteral = 3.14       // Floating-point literal
var hexLiteral = 0x1F         // Hexadecimal literal (31 in decimal)
var octalLiteral = 0755       // Octal literal (493 in decimal)
var stringLiteral = "Hello, World!"  // String literal
var boolLiteralTrue = true    // Boolean literal
var boolLiteralFalse = false   // Boolean literal

ascii := 'a'     // Runes Literal
unicode := 'B'   // Runes Literal
newline := '\n'  // Runes Literal
// when printed using the fmt package  we have the below
fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
fmt.Printf("%d %[1]c %[1]q\n", unicode) // "68 D 'D'"
fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"


```