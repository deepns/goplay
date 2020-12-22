package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {
	BasicTypes()
	Constants()
	Loops()
	Conditionals()
	Structs()
	Slices()
	Maps()
}

// BasicTypes lists usage of some basic data types in Go
func BasicTypes() {
	// Basic types
	// bool, int, uint, float32, float64, complex64, complex128, string, rune
	// int and uint can be specified bit precision (8, 16, 32, 64) e.g. int8, uint16

	// var name type' creates a variable of type with the given name and
	// initializes with the type's zero value
	var zip int
	var city string
	var temperature float32

	fmt.Printf("Zero values of int=%v, string=%v, float=%v\n", zip, city, temperature)

	// type inference using :=
	// name := value creates a variable with the value and type inferred from the value
	cityHasAirport := false
	fmt.Printf("cityHasAirport := %v creates a variable cityHasAirport of type %T\n", cityHasAirport, cityHasAirport)

	// Go also supports pointers, similar to C/C++
	// pointers can be declared with var keyword
	// var X *T // declares a pointer variable named X of type T
	var pZip *int
	pZip = &zip

	// pointers can also be obtained using & operator during type inference
	// p := &i
	pCity := &city

	fmt.Printf("addr(zip)=%p, addr(city)=%p\n", pZip, pCity)

	// pointers are deferenced using * operator
	*pZip = 90901

	// Pointer arithmetic is not supported in Go.
	// This will report compilation error as *int and int are
	// mismatched operands for the operation +
	// pZip += 1
}

// Constants ...
func Constants() {
	// Constants declared using const keyword
	// No type inference for const, so can't assign using :=
	const Pi = 3.14
	const Foo = "Bar"

	// Constant declarations can also be grouped like variable declarations
	const (
		OnePi = Pi
		TwoPi = 2 * Pi
	)

	fmt.Printf("Declared constants OnePi(%v), TwoPi(%v)\n", OnePi, TwoPi)
}

// TypeConversion example
func TypeConversion() {
	f := 43.2
	// Print type of a value using %T specifier
	fmt.Printf("type(f)=%T\n", f)

	// Type conversion is done by type(value)
	fmt.Printf("int(%v)=%v\n", f, int(f))
}

// Sqrt returns square root of x
// Example for a function taking a float64 argument and returning a float
func Sqrt(x float64) float64 {
	// limiting the difference to 0.0000005
	const floor = 0.0000005
	z := 1.0
	for delta := math.Inf(1); math.Abs(delta) > floor; {
		// Based on Newton's method of approximation
		delta = (z*z - x) / (2 * z)
		z -= delta
	}
	return z
}

// Loops shows a demo of Go's looping construct
func Loops() {
	// Go has only for looping construct. No `while`
	// for init-stmt; cond-expr; post-stmt {
	//
	// }
	// No parantheses needed after `for`
	for i := 0; i < 3; i++ {
		fmt.Printf("Simple for loop at iteration %v\n", i)
	}

	// init can take multiple values too, just like any other initialization
	for i, j := 0, 3; i < j; i++ {
		fmt.Printf("For loop with multiple inits: at iteration %v\n", i)
	}

	// init-stmt and post-stmt can be ignored to have a `while` like behavior
	i := 0
	for i < 3 {
		fmt.Printf("while(%v < 3)\n", i)
		i++
	}

	// infinite loop
	for {
		// I don't want to run forever.
		break
	}

	// Seeding the random generator with current time
	// Otherwise we get the same number every time
	rand.Seed(time.Now().UnixNano())

	// Sqrt shows another example of using the for loop constructions
	for i := 0; i < 5; i++ {
		x := rand.Intn(100)
		fmt.Printf("Sqrt(%v)=%v\n", x, Sqrt(float64(x)))
	}
}

// Conditionals ...
func Conditionals() {
	ifElse()
	// switch..case
	// defer
}

func ifElse() {
	// Syntax goes like this
	// if condition {
	//		doSomething
	// }

	// No ( ) needed around the conditions

	// if assignment; condition {
	//		doSomething
	// }

	// another commonly used idiom is when a function returns
	// value and boolean status
	// if val, ok := someFunc(); !ok {
	// 		somethingIsNotOk()
	// }

	// another one, when a function returns an error value
	// in the return values
	// if val, err := someFunc(); err != nil {
	//		handleError()
	// }

	var err error
	if err == nil {
		fmt.Println("Checking err == nil in a sample if block")
	}

	// variables declared in the if assignment is visible only
	// in the scope of the if..else block
	if num := rand.Intn(100); num%2 == 0 {
		fmt.Printf("random number(%v) is even\n", num)
	} else {
		fmt.Printf("random number(%v) is odd\n", num)
	}

	// can't access num here, since it is outside the scope
	// fmt.Println(num)

	// this stype of variable assignment in one line, followed
	// by a if block to check error is quite common in Go world
	fh, err := os.Open("foo.bar")
	if err != nil {
		fmt.Printf("Failed to open foo.bar due to %v\n", err)
	}
	fh.Close()

	// an example to write multi line conditions
	num := rand.Intn(100)
	if num%2 == 0 &&
		num%3 == 0 &&
		num%4 == 0 {
		fmt.Printf("random number(%v) is a multiple of 12\n", num)
	} else {
		fmt.Printf("random number(%v) is not a multiple of 12\n", num)
	}
}

// Structs ...
func Structs() {
	// struct in Go is little similar to C struct

	// Creates a struct named X with a field
	// named Y of type T
	// type X struct {
	//     Y T
	// }

	type Point struct {
		x int
		y int
	}

	// structs declared like any other type (var X T)
	var p1 Point

	// struct fields are accessed using dot
	p1.x, p1.y = 10, 20

	// struct fields are initialized with values inside {...}
	p2 := Point{10, 30}
	fmt.Printf("%v and %v are structs of type %T\n", p1, p2, p1)

	// when multiple fields of a struct are of the same type,
	// the declaration can be grouped together.
	type Line struct {
		p1, p2 Point
	}

	// structs can be created with struct literals too
	// fields can be specified partially too
	verticalLine := Line{
		p2: Point{10, 15},
		p1: Point{10, 25},
	}

	lineOnYAxis := Line{
		p2: Point{0, 20},
		// Line.p1 is implictly derived using the zero values
	}

	// in this case, both Line.p1 and Line.p2 are implicity derived.
	// so origin here is {{0, 0}, {0, 0}}
	origin := Line{}

	// I hate the lint errors. so using the above variables
	// in a print call
	fmt.Printf("%v, %v and %v are of type %T\n",
		verticalLine,
		lineOnYAxis,
		origin,
		origin)

	// structs can also be accessed using pointers.
	pLine := &verticalLine
	fmt.Printf("addr(%#v)=%p\n", pLine, pLine)

	// for simplicity, struct fields can be accessed using "."
	// for both pointer and non-pointer types
	pLine.p1.y = 40
	(*pLine).p1.y = 50
}

// Slices ...
func Slices() {
	// Arrays are fixed size collection of a type
	// Similar to arrays in C/C++
	// var X [L]T // creates an array of type T with length L
	// values are filled with zero-value of the given type
	var scores [11]int

	// Go has length and capacity for container types
	// Since arrays are fixed size, length and capacity are same
	// Also, type of array will include the array size as well
	fmt.Printf("Created an array scores of type %T of length=%v capacity=%v\n",
		scores, len(scores), len(scores))

	// Slices are variable length types
	// A slice is more of a reference to an underlying array
	// A slice can be created in many ways

	// By slicing an existing array/slice
	openerScores := scores[0:2]
	fmt.Printf("%v is a slice of %T at pointing to %p\n",
		openerScores, openerScores, openerScores)

	// A slice can also be created using
	//	make([]T, length)
	//	make([]T, length, capacity)
	wickets := make([]int, 11)

	// this creates an empty slice of 0 length
	// an empty slice compares to nil
	var players []string
	if players == nil {
		fmt.Printf("players is an empty %T\n", players)
	}

	// A slice can also be created using literals
	jerseyNumbers := []int{10, 7, 23, 99, 24}

	fmt.Printf("wickets is a slice created using 'make(%T, %v)'\n", wickets, len(wickets))
	fmt.Printf("players is a slice created using 'var players %T'\n", players)
	fmt.Printf("jerseyNumbers is a slice created using %T{%v}, len=%v, cap=%v\n",
		jerseyNumbers, jerseyNumbers, len(jerseyNumbers), cap(jerseyNumbers))

	// multi-dimensional slices
	// (aka slice of slices)
	var extras [][]int
	fmt.Printf("extras is a slice of slice, of type %T\n", extras)

	type Stadium struct {
		city     string
		capacity uint
		rating   float32
	}

	venues := [][]Stadium{
		{
			Stadium{"Melbourne", 50000, 8.3},
			Stadium{"Sydney", 100000, 6.3},
		},
		{
			Stadium{"Chennai", 75000, 10.0},
			Stadium{"Bengaluru", 50000, 6.0},
		},
		{
			Stadium{"Lords", 120000, 7},
			Stadium{"Old Trafford", 50000, 7},
		},
		[]Stadium{}, // empty slice
	}

	fmt.Printf("venues is a slice of slice of type %T, of len=%v, cap=%v\n",
		venues,
		len(venues),
		cap(venues))

	//
	// Accessing an array or a slice
	//

	// Changes in the underlying array will be reflected
	// in the slice as well since the slice is just pointing
	// to the underlying array
	scores[0] = 100 // now openerScores[0] will also be 100
	fmt.Println(openerScores)

	// slicing is typically done by specifying start:end
	// special cases are:
	//		start: // slices the array from the given start till the length of the slice/array
	//		:end // slices the array from 0 till the given end
	batsmenScores := scores[:6]
	bowlerScores := scores[6:]
	fmt.Printf("batsmenScores and bowlerScores are of type %T and %T\n", batsmenScores, bowlerScores)
	fmt.Printf("len(batsmenScores)=%v, cap(batsmenScores)=%v\n", len(batsmenScores), cap(batsmenScores))
	fmt.Printf("len(bowlerScores)=%v, cap(bowlerScores)=%v\n", len(bowlerScores), cap(bowlerScores))

	// bowlerScores is only of length 5, so accessing beyond that
	// size will raise index out of range panic.
	// fmt.Println(bowlerScores[6])

	// No Python style reverse lookup using negative indexes
	// fmt.Println(openerScores[-1])

	// New items can be added to a slice using append
	// append is a built in function that takes a slice reference and adds
	// another slice or a value of the slice type to the given slice.
	// If the slice doesn't have enough capacity, then a new underlying
	// array will be allocated and its address will be returned
	players = append(players, "Aaron")
	fmt.Printf("len(players)=%v, cap(players)=%v, addr(players)=%p\n",
		len(players), cap(players), players)

	// when a new array has to be allocated underneath,
	// the capacity is typically increased in powers of 2
	for _, player := range []string{"Chris", "Mahi", "Rahul", "Ashwin"} {
		players = append(players, player)
		// this will show how the length and capacity of the slice
		// changes while we extend it. Notice the address changes
		// when new array is allocated.
		fmt.Printf("len(players)=%v, cap(players)=%v, addr(players)=%p\n",
			len(players), cap(players), players)
	}

	// appending to a slice within slice.
	venues[3] = append(venues[3], Stadium{"Jamaica", 35000, 9.3})
}

// Maps ...
func Maps() {
	// maps are key-value pairs
	// similar to slice, they can be created with "make" or literals

	// declare a map like this
	// var mapname map[key-type]value-type
	var counter map[string]int

	// Declaration only declares a nil map
	// New values can't be inserted into the map
	fmt.Println(counter)
	// this is not allowed since map is not allocated yet
	// counter["odds"] = 5

	// allocate a map with make keyword
	counter = make(map[string]int)

	// Insert a key-value pair into the map
	counter["odds"] = 5
	fmt.Println(counter)

	// map literals
	// keys are inserted and retrieved in sorted order
	stepCount := map[string]int{
		"Wednesday": 3949,
		"Monday":    2500,
		"Tuesday":   4000,
		"Saturday":  4929,
	}

	fmt.Printf("stepCount is a %T created with map literals with values %v\n",
		stepCount,
		stepCount)

	// range keyword can be used to iterate maps as well
	// range returns a key, value from the map
	for day, steps := range stepCount {
		fmt.Printf("I walked %v steps on %v\n", steps, day)
	}

	// Deleting an element from map
	// delete(map, key)
	dayToDelete := "Wednesday"
	delete(stepCount, dayToDelete)
	fmt.Printf("Deleted %v from %v\n", dayToDelete, stepCount)

	// When key is missing, lookup will return zero value of
	// the value-type. In this example, stepCount maps string
	// to int. Since zero value of int is 0, accessing a
	// missing would return 0.
	steps := stepCount[dayToDelete]
	fmt.Printf("I walked %v steps on %v\n", steps, dayToDelete)

	// Existence of a key in map can be tested with two value assignment
	_, ok := stepCount[dayToDelete]
	if !ok {
		fmt.Printf("Step count not available for Wednesday\n")
	}

	// Maps support len() too, returning the number of elements
	// in the map
	fmt.Printf("len(stepCount)=%v\n", len(stepCount))
}
