![go-logo](/images/go_logo.png)

# learning-go

## Resources

- [Go Tour](https://go.dev/tour)
- [Golang Tutorials Playlist](https://www.youtube.com/playlist?list=PLzMcBGfZo4-mtY_SE3HuzQJzuj4VlUG0q)

## Introduction to Go Programming

Main features:
- An imperative language
- Developed by a team at Google

Downloads
- Visit https://go.dev/dl/ to download the latest Go build
- Add GOTOPATH if not done automatically
- Add Go extension in VSCode

### Printing Hello World

```go
// need to declare a package on top of every file
package main

// import fmt module to display messages onto console
import "fmt"

// entry point to our functions
func main() {
	// everything inside this happens when we run this code
	fmt.Println("Hello world!")
}
```

### Build Go code
This will build the code present in main.go file and create a new executable file - main.exe
```
>> go build main.go
>>

>> main.exe
Hello world!
```

### Build and Run Go code
This automatically compiles the code and run the same
```
>> go run main.go
Hello world!
```

## Variables

### Variable Naming
- Name can have alphabets, numbers and underscores
- Name cannot start with a number
- Name cannot have spaces
- Name cannot be a reserved keyword like func, fmt, main, etc.

### Variable Declaration and Initialization

#### Declaring and Initializing the variable at the same time
```go
var <variable_name> <data_type> = <variable_value>
```

#### Declaring the variable first and Initializing later
```
var <variable_name> <data_type>
<variable_name> = <variable_value>
```

#### Examples
```go
var first_name string = "Rajat"
var last_name string
last_name = "Goyal"
```

### Data Types

Data Types are listed here - [data_types.md](./data_types.md)

```go
var number uint8 = 260 // overflows uint8

var number uint16 = 260 // works since 260 is in the range of uint16
number = number + 5
fmt.Println("Hello World!", number)
```

### Implicit vs Explicit Assignment

#### Explicit Declaration
Explicity mentioning the data type of the variable
```go
var number uint16 = 260
```

#### Implicit Declaration
Asking GoLang to guess the data type of the variable from its value
```go
var number = 260
```

```go
func main() {
	var var1 = 2
	var var2 = 200.98
	var var3 = "hello"

	// prints the type of the variable
	fmt.Printf("%T", var1) // int
	fmt.Printf("%T", var2) // float64
	fmt.Printf("%T", var3) // string
}
```

#### Walrus operator
```go
var1 := 6
var1 = "str" // cannot use "str" (type string) as type int in assignment

var2 := "hello"
var2 = 5 // cannot use "5" (type int) as type string in assignment
```

```go
var number uint64 // default value of number i 
var boolean bool

fmt.Println(number, boolean) // 0 false
```

## More about fmt and Printing to Console
```go
func main() {
	fmt.Printf("Hello %T %v", 10, 10) // int 10
	var x string = fmt.Sprintf(...) // this will store the formatted string as the value of the variable
}
```

```go
func main() {
	var out string = fmt.Sprintf("Number: \t %07d is cool", 45) // Number:		0000045 is cool
	fmt.Println(out)
}
```

## Console Input
```go
package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // new scanner object which will read from os standard input
	fmt.Printf("Type Something: ")
	scanner.Scan()
	stringInput := scanner.Text() // read a string value from the console

	fmt.Printf("Input: %q", input)
}
```

## Type Conversion

### Number to String
```go
numberValue := 34
stringValue := str(numberValue)

floatValue = 34.56
stringValue = str(floatValue)
```

### Integer to Float
```go
numberValue := 34
float32Value := float32(numberValue)
float64Value := float64(numberValue)
```

### Float to Integer
```go
float32Value := 123.456
numberValue := uint8(float32Value)
```

### String To Integer
```go
package main

import (
	"fmt"
	"strconv" // module required for string conversion
)

func main() {
	// string to integer conversion syntax - numberInput, _ := strconv.ParseInt(scanner.Text(), <base>, <bitSize>)
	integerValue, _ := strconv.ParseInt("234", 10, 64)
	fmt.Println(integerValue)
}
```

## Arithmetic Operators & Math

### Division Operator
```go
func main() {
	var num1 int = 8
	var num2 int = 4
	div1 := num1 / num2
	fmt.Printf("%d", div1) // 2

	var num3 float64 = 9
	var num4 int = 4
	div2 := num4 / float64(num3) // Convert num3 to float64 for proper division
	fmt.Printf("%d", div2) // 1.25

	var num5 int = 9
	var num6 int = 4
	div3 := num5 / num6 // This will store the integral part of the division
	fmt.Printf("%d", div3) // 2
}
```

### Modulus Operator
```go
func main() {
	var num1 int = 9
	var num2 int = 4
	mod1 := num1 % num2 // This will store the remainder in the division
	fmt.Printf("%d", mod1) // 1

	var num3 int = 0
	var num4 int = 4
	mod2 := num3 % num4
	fmt.Printf("%d", mod2) // 0

	var num5 int = 4
	var num6 int = 0
	mod3 := num5 % num6
	fmt.Printf("%d", mod) // panic: runtime error: divide by zero
}
```

## Conditions & Boolean Expressions
Following operators are the standard operators supported by Go:
```
>, <, >=, <=, ==, !=
```

Examples
```go
x = 5
y = 6.5
val := float64(x) < y // comparison can be done within variables of same data types
fmt.Printf("%t", val)

str1 = "hello"
str2 = "Hello"
val = str1 == str2 // false
val = str1 < str2 // false, since the comparison is being done based on the ASCII values of the characters
```

### Chained Conditionals
AND, OR, NOT

```go
true AND true = true
true AND false = false
false AND true = false
false AND false = false

true OR true = true
true OR false = true
false OR true = true
false OR false = false

NOT false = true
NOT true = false
```

## If, Else If, Else
```go
func main() {
	if condition {
		// statements inside this will run only if the condition is true
	} else if condition {
		// statements inside this will run only when the conditions before this are false, and the current condition is true
	} else {
		// statements inside this will run only if all the above conditions are false
	}
}
```

```go
package main

import "fmt"

func main() {
	number := 15
	if number % 3 == 0 AND number % 5 == 0 {
		fmt.Println("FizzBuzz")
	} else if number % 3 == 0 {
		fmt.Println("Fizz")
	} else if number % 5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println("Number if neither divisible by 3 nor by 5")
	}
}
```

## For Loops
There are no "while" loops in Go. A "for" loop can work as "while" loop also.

```go
func main() {
	x := 0
	for x < 5 {
		fmt.Println(x)
		x += 1 // x++ is equivalent to x+1, and it won't update the variable, either do x = x + 1 or x += 1
	}
} 
```

```go
func main() {
	// Initialization, Termination, Increment
	for x := 0; x <= 5; x++ {
		fmt.Println(x)
		if x % 3 == 0 {
			continue // using this keyword will help us not execute the statements after this block and directly continue to the next iterator
		} else if x % 5 == 0 {
			break // using this keyword will help us break the entire loop and start executing the statements after the loop
		}
	}
} 
```

## Swtich Statement

```go
func main() {
	ans := 5
	switch ans {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3,4,5:
			fmt.Println("greater than two and less than six")
		default: // This is an optional statement
			// These statements will be executed if none of the cases met the condition.
			fmt.Println("not a case")
	}

	ans = -1
	// switch case can be run without any variable/condition. In that scenario, cases will act like if, else if conditions.
	switch {
		case ans > 0:
			fmt.Println("ans is greater than zero")
		case ans < 0:
			fmt.Println("ans is less than zero")
		case ans == 0:
			fmt.Println("ans is equal to zero")
	}
}
```

## Arrays

### Syntax
```go
// Declaring an array
var <array_variable_name> [<array_size>]<array_data_type>

// Initializing the array along with declaration
<array_variable_name> := [<array_size>]<array_data_type>{...<elements>...}
```

### Example
```go
func main() {
	var stringArray [5]string
	var integerArray [5]int

	// Explicitly defining an array
	arr := [3]int{4,5,6}
	fmt.Println(arr) // prints the elements in the array, space-separated
	fmt.Println(arr) // prints the length of the array
	fmt.Println(arr[1]) // this will print the second element of the array, since the arrays are 0-based indexed
	arr[2] = 7 // this will modify the element at index = 2 to the new value

	arr2D := [2][2]int{{1,2},{3,4}}
	fmt.Println(arr2D)
	fmt.Println(arr2D[0][1]) // 2

	for i := 0; i < len(arr2D); i++ {
		for j := 0; j < len(arr2D[i]); j++ {
			fmt.Print(arr2D[i][j])
		}
		fmt.Println() 
	}
} 
```

## Slices

### Declaration & Initialization
```go
func main() {
	var x [5]int = [5]int{1,2,3,4,5} // initializing an array
	var s1 []int = x[:] // start picking the elements from the very start and pick till the very end
	var s2 []int = x[1:3] // start picking the elements including start index and until (excluding) end index
	fmt.Println(len(s2)) // length of this slice is 2
	fmt.Println(cap(s2)) // capacity is 4

	fmt.Println(s2[:cap(s2)]) // [2,3,4,5]

	// Slice of slice
	fmt.Println(s2[1:]) // [3]

	var a []int = []int{5,6,7,8,9} // initializing a slice directly, without specifying the size initially
	fmt.Println(cap(a)) // 5
	fmt.Println(cap(a[:3])) // 5 - capacity is still 5 because it points to the underlying array which was created at the start
} 
```

### Append Method
```go
func main() {
	var a []int = []int{5,6,7,8,9}
	b := append(a, 10) // creating a new slice
	fmt.Println(b) // [5,6,7,8,9,10]

	a = append(a, 10) // appending an element in the existing slice
	fmt.Println(a) // [5,6,7,8,9,10]
}
```

### Make keyword
```go
func main() {
	a := make([]int, 5)
	fmt.Println(a)
	fmt.Println("%T", a) // []int
} 
```

## Range

```go
func main() {
	var a []int = []int{1,3,4,56,7,12,4,6}

	// standard way to loop over the slice/array
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// we can use range keyword to iterate over the slice/array
	for index, element := range a {
		fmt.Printf("%d: %d\n", index, element)
	}
	
	// we can use underscore ('_') if we don't plan to use index value
	for _, element := range a {
		fmt.Printf("%d\n", element)
	}
}
```

## Maps

Syntax
```go
func main() {
	var <variable_name> map[<key_data_type>]<value_data_type>
}
```

```go
func main() {

	var mp map[string]int // map from string to integer
	var mp map[string]int = map[string]int{
		"apple":5,
		"pear":6,
		"orange":9,
	}
	fmt.Println(mp) // map does not keep track of the order in which they were added
	fmt.Println(mp["apple"])

	mp = make(map[string]int)
	mp["apple"] = 900
	fmt.Println(mp["apple"])

	mp["banana"] = 900 // if key already exists, override the existing value, otherwise add this key-value pair
	delete(mp, "apple") // deletes the entry with the given key
	fmt.Println(mp) // map[banana:900]

	// How to check if the key exists in the map?
	val, ok := mp["apple"] // If key exists, val will store the value of the given key in the map, ok will be set as true. If the key doesn't exist, then val will store the default value of value data type, which is 0 here in case of int, and ok will be set to false.

	fmt.Println(len(mp)) // Prints the length of the map
}
```

## Functions

```go
package main

import "fmt"

// Function without any parameters
func test1() {
	fmt.Println("test1")
}

// Function with one parameter
func printInteger(x int) {
	fmt.Println(x)
}

// Function with two parameters
func printAddition(x int, y int) {
	fmt.Println(x + y)
}

// Function with return type
func addNumbers(x int, y int) int {
 	return x + y
}

// Function with multiple return values
func addAndSubtractNumbers(x int, y int) (int, int) {
 	return x + y, x - y
}

// Function with labelled return values
func addAndSubtractNumbers(x int, y int) (r1 int, r2 int) {
	r1 = x + y
	r2 = x - y
 	return
}

// Function with labelled return values
func addAndSubtractNumbers(x int, y int) (r1, r2 int) {
	r1 = x + y
	r2 = x - y
 	return
}

// Function with deferred statement
func addAndSubtractNumbers(x int, y int) (r1 int, r2 int) {
	defer fmt.Println("exiting the function") // this statement will get executed after the return statement
	r1 = x + y
	r2 = x - y
 	return
}

func main() {
	test1() // this is how we call the function, can be re-used multiple times

	printInteger(5)
	printAddition(5,6)

	sum := addNumbers(6,7)
	fmt.Println(sum)

	sum, subtract := addAndSubtractNumbers(10, 3)
	fmt.Println(sum, subtract)
}
```

## Advanced Function Concepts

```go
func main() {
	// create a function like this, which can be called anytime using that function variable
	test := func() {
		fmt.Println("hello!")
	}

	// calling the function using that function variable
	test()

	// we can call the function inline at the same time when we are defining the function
	ans := func(x int) int {
		return x * -1
	}(8)
	fmt.Println(ans) // -8

	test1 = func(x int) int {
		return x * -1
	}
	test2 = func(x int) int {
		return x * 7
	}
	printFunctionOutput(test1, 7) // -7
	printFunctionOutput(test2, 7) // 49
}

// we can pass a function as function parameter also
func printFunctionOutput(myfunc func(int) int, x int) {
	fmt.Println(myFunc(x))
}
```

### Function Closure
```go
package main

import "fmt"

// Function can be used as return value also
func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	returnFunc("hello")() // prints "hello"
	x := returnFunc("goodbye")
	x() // prints "goodbye"
}
```

## Mutable & Immutable Data Types

Example of Immutable data type
```go
func main() {
	var x int = 5
	y := x // y is equal to the value of x
	y = 7 // value of y gets updated here
	fmt.Println(x, y) // 5, 7
}
```

Example of Mutable data type
```go
func main() {
	var x []int = []int{3, 4, 5} 
	y := x // y is equal to the slice that x is pointing to
	y[0] = 100 // value of y gets updated here, which results into the updation in x also
	fmt.Println(x, y) // [100, 4, 5] [100, 4, 5]

	var x map[string]int = map[string]int{"hello": 2}
	y := x
	y["world"] = 100
	fmt.Println(x, y) // map[hello:2 world:100] map[hello:2 world:100]

	var x [2]int = [2]int{3,4} // x := []int{3,4}
	y := x
	y[0] = 100
	x[1] = 200 // this will change value of y also, since both are referring to the same address
}
```

```go
package main

import "fmt"
func changeFirst(slice []int) {
	slice[0] = 100
}
func  main() {
	var x []int = []int{3,4,5}
	fmt.Println(x)
	changeFirst(x)
	fmt.Println(x)
}
```

## Pointers & Dereference Operator

```go
func main() {
	x := 7
	y := &x // stores the reference of x, or pointer to x
	fmt.Println(x, y) // prints something like 7 0xc0000100a0

	*y = 8 // de-reference the pointer
	fmt.Println(x, y) // 7 8
}
```

```go
func changeValue(str *string) { // asterisk to left of a data type means that it is a pointer to that data type
	*str = "changed!" // asterisk to left of a variable name means the value of this variable (value at this address)
}

func changeValue2(str string) {
	str = "changed again!"
}

func main() {
	toChange := "hello"
	fmt.Println(toChange)
	// &x means the address or pointer to that variable, which points to the actual value
	changeValue(&toChange) // this will change the value since str is passed as pointer / reference
	fmt.Println(toChange)
	changeValue2(toChange) // this will not change the value since str is passed as value directly
	fmt.Println(toChange)

	// * - dereference pointer
	// & - get pointer 
}
```

```go
func main() {
	toChange := "hello"
	var pointer *string = &toChange          // pointer stores the address of the string
	var pointerToPointer **string = &pointer // pointer to pointer stores the address of the pointer
	fmt.Println(pointer, pointerToPointer)
}
```

## Structs and Custom Types

```go
package main

type Point struct {
	x int32
	y int32
}

func main() {
	var p1 Point = Point{1,2}
	var p2 Point = Point{-5,7}
	fmt.Println(p1.x)
	fmt.Println(p2.y)

	p3 := Point{3,4}
	p4 := Point{x: 7} // variables which are not provided some value will take default values

	p5 := &Point{y:3}
	fmt.Println(p5)
	changeX(p5)
	fmt.Println(p5)
	(*p5).x = 50 // this is the usual thinking since we have to de-reference the pointer first, and then change the values
	p5.x = 50 // this is also supported with structs where we don't need to dereference the pointer before accessing the values	
}

func changeX(pt *Point) {
	pt.x = 100
}
```

### Embedded Struct
```go
func main() {
	type Point struct {
		x int32
		y int32
	}

	type Circle struct {
		radius float64
		center *Point
		*Point
	}

	type Circle2 struct {
		radius float64
		*Point // Embedded struct - this is possible only when the fields inside Point struct are not already present in Circle2 struct.
	}

	func main() {
		c1 := Circle{4.56, &Point{4,5}}
		fmt.Println(c1) // 4.56 0xc00000a0
		fmt.Println(c1.center.x) // 4

		c2 := Circle2{4.56, &Point{4,5}}
		fmt.Println(c2)
		fmt.Println(c2.x) // 4 - we can directly access fields of embedded struct as if they are first class fields
		fmt.Println(c2.y) // 5
	}
}
```

## Methods
```go

type Student struct {
	name string
	grades []int
	age int
}

func (s Student) getAge() int {
	return s.age
}

func (s Student) setAge(age int) {
	s.age = age
}

func (s *Student) setAge2(age int) { // passing student as pointer so that fields can be updated
	s.age = age
}

func (s Student) getAverageGrade() float32 { // we don't need to pass the pointer to student since we are not changing anything
	sum := 0
	for _, value := range s.grades {
		sum += value
	}
	return float32(sum) / float32(len(s.grades))
}

func main() {
	s1 := Student{"Rajat", []int{70, 90, 80, 85}, 24}
	s1.getAge()

	fmt.Println(s1)
	s1.setAge(20) // this will not update the age since student has been passed as value
	fmt.Println(s1)
	s1.setAge2(20) // this will update the age since student has been passed as pointer
	fmt.Println(s1)

	avg := s1.getAverageGrade()
	fmt.Println(avg)
}
```

## Interfaces
```go
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width float64
	height float64
}

type Shape2 interface { // structs can implement multiple interfaces
	perimeter() float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r *Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	c1 := Circle{4.5}
	r1 := Rectangle{5, 7}
	shapes := []Shape{c1, &r1}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
```
