---
---

Golang
======

Oh Go! I finally get to use you.

## Notes from "A Tour of Go"

If you already installed Go then you can run an interactive tour of Go locally: `go tool tour`

(This code probably won't compile)

```go
package main

// Packages to import,
// Better style to use a single import statement.
import (
	"fmt"       // package fmt
	"math/rand" // package rand
)

// NOTE: names are exported from package if they begin with a capital letter

// Types come after the variable name.
// Return type comes end of function signature.
func add(x int, y int) int {
  return x + y
}

// you can omit type from all but last function parameters if they are the same
func add(x, y int) int {
  return x + y
}

// functions can return zero or more values
func swap(x, y string) (string, string) {
	return y, x
}

// return values may be named
// just return once you set them
func divmod(a int, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

// variables can be declared in packages or functions
var num1 int

func nums() {
	
	// variables can be initialized at declaration
	var num2 int = 2
	
	// multiple variables can be declared in one line
	var num3, num4 int = 3, 4
	
	// within functions you can use short variable declarations
	// all statements outside of a function must begin with a keyword
	// type is inferred by value
	num5 := 5
	num6 := num2 // num6 = 2
}

// variables can be declared in blocks too
// let's look at other types too
var (
	bool       inBlock = true
	string     message = "Hi!"
	int        num = 1 		     // alias for int32 or int64 -- arch dependent
	// also int(8,16,32,64)
	uint       unum = 1         // alias for int32 or int64 -- arch dependent
	// also uint(8,16,32,64)
	unintptr   ptr = nil	       // alias for int32 or int64 -- arch dependent
	byte       char = 'A'       // alias for unit8
	rune       emoji = '😀'     // alias for int32
	float32    decimal = 1.2345 // also float32
	complex128 crazy = -5 + 12i
)

// variables without explicit initial value are given their zero value
// QUESTION: is still setting to zero explicitly considered bad style?
var bool   myBool // = false
var int    myNum  // = 0
var string myName // = ""

// call type as function for type conversion
// type conversion must be explicit
// implicit conversion will result in compiler error
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// constants are declared like variables
// but you must use const and are not explicitly typed
const Initials = "HGPA"

// for loop
sum := 0
for i := 0; i < 10; i++ {
		sum += i
}

// init and post statements are optional
// and if you drop the semi-colons then you have Go's while statement
for sum < 1000 {
	sum += sum
}

// infinite loop
for {
	
}

// NOTE: remember that Go has no while statement

// If-statement
if sum > 1000 {
	fmt.Println("That's not that big.")
} else {
	fmt.Println("That's def not that big.")
}

// If with short statement.
// Variables declared with short statement are only in scope for if-else block
if sum := a + b; sum > 1000 {
	fmt.Println("%d is big!", sum)
} else {
	fmt.Println("%d is small", sum)
}

// Switch statement
// They're evaluating from top to bottom.
// They're is no "break" since fallthrough is NOT default behavior
// You must use `fallthrough` statement to get this behavior
switch os := runtime.GOOS; os {
case "darwin":
	fmt.Println("OS X.")
case "linux":
	fmt.Println("Linux.")
default:
	fmt.Printf("Who cares! (BSD is still cool)")
}

// Switch cases may be conditions.
// A conditionless switch is clean way to write long if-then-else chains.
t := time.Now()
switch {
case t.Hour() < 12:
	fmt.Println("Good morning!")
case t.Hour() < 17:
	fmt.Println("Good afternoon.")
default:
	fmt.Println("Good evening.")
}

// Defer statements defer the execution of a function until surrounding function returns.
// Deferred call's arguments are evaluated immediately, but the call is not executed.
// Deferred call are pushed onto a stack.
// The following will print "hello\nworld\n"
func helloWorld()  {
	defer fmt.Println("world")
	fmt.Println("hello")
}

// Go has pointers (but no pointer arithmetic).
i := 42
p := &i // point to i (address of i)
j := *p // get i through pointer (dereference)
*p = 21 // set i through pointer (dereference)

// Go has structs.
type Vertex struct {
	X int
	Y int
}

// You can initialize a variable of type struct a few ways.
v1 := Vertex{1,2}
v2 := Vertex{X: 1} // Y: 0
v3 := Vertex{}     // X: 0, Y:0

// You can also point to a struct.
vp := &Vertex{1,2} // *Vertex

// You can get/set struct fields with dot operator.
v1.X = 3
v1.Y = 4

// dot operator will also dereference if necessary.
vp.X = 5 // equivalent to (*vp).X

// arrays are of the form [n]T
// where n is the fixed size and T is the type
var arr [10]int

// you can also declare contents of the array
primes := [6]int{2,3,4,5,7,11,13}

// a "slice" is a dynamically-sized flexible view/window into the elements of an array
// it does not store data, it just describes section of an array
var firstThreePrimes []int = primes[0:3] // {2,3,5}

// a slice literal create an array then builds slice that references it
someBools := []bool{true, false, true, true, false, true}

// slice expressions return other slices of the same array
// these slice expressions are equivalent
// similar to python
var a [10]int
a[0:10]
a[:10]
a[0:]
a[:]

// slices have both length and capacity
len(firstThreePrimes) // 3
cap(firstThreePrimes) // 6

// you can create a nil slice
var s []int

// create a dynamically-sized array via slice using make
a := make([]int, 5) //len(a) =5

// you can also specify a capacity
b := make([]int, 5, 10) //len(b) = 5, cap(b) = 10

// slices can include other slices
board := [][]string{
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
	[]string{"_", "_", "_"},
}

// append to a slice
s = append(s, 0)
s = append(s, 1)

// append more than one element at a time
s = append(s, 2,3,4,5)

// https://blog.golang.org/go-slices-usage-and-internals

// there is a range version of the for loop
for i, v := range primes {
	fmt.Printf("Not prime: %d\n", v * 2)
}

// You can skip index by using _ instead

// dynamically create a two-dimensional slice
m := make([][]uint8, 10, 10)
for x := range p {
	m[x] = make([]uint8, 10, 10)
	for y := range m[x] {
		f := x * y
		m[x][y] = uint8(f)
	}
}

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}

```

## Atom Snippets

- `fp` -> `fmt.Println("")`
- `ff` -> `fmt.Println("", var)`
- `for`
- `func`
- `if`
- `iferr`
- `switch`

> Goroutines let you run multiple computations simultaneously.
> Channels let you coordinate the computation, by explicit communication.
