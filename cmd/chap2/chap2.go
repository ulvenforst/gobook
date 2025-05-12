// each package starts with a package declaration
package chap2

import (
	"bufio"
	"fmt"
	"github.com/ulvenforst/gobook/cmd/chap2/genconv"
	"github.com/ulvenforst/gobook/cmd/chap2/tempconv"
	"os"
	"strconv"
	"strings"
)

// Package level declaration
const boilingF = 212.0 // Constants are declared with the const keyword

// All the source code lives in the universe block, which is a lexical block.

func RunChap() { // This braces define a syntactic block
	fmt.Println("Hello, from chap2!")

	// EXERCISE 2.1
	fmt.Println("=== EXERCISE 2.1 ===")
	x := 2
	// Celsius to Kelvin
	fmt.Println(tempconv.CToK(tempconv.Celsius(x)))
	// Kelvin to Celsius
	fmt.Println(tempconv.KToC(tempconv.Kelvin(x)))
	// Fahrenheit to Kelvin
	fmt.Println(tempconv.FToK(tempconv.Fahrenheit(x)))
	// Kelvin to Fahrenheit
	fmt.Println(tempconv.KToF(tempconv.Kelvin(x)))

	// EXERCISE 2.2
	fmt.Println("=== EXERCISE 2.2 ===")
	units := os.Args[1:]
	if len(units) == 0 {
		input := bufio.NewScanner(os.Stdin)
		if input.Scan() {
			line := input.Text()
			units = strings.Fields(line)
		}
		ShowConversions(units)
	} else {
		ShowConversions(units)
	}

}

// Exercise 2.2
func ShowConversions(units []string) {
	for _, unit := range units {
		unit, err := strconv.ParseFloat(unit, 64) // Here I'm shadowing the unit variable
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("=== Converting %g ===\n", unit)
		// Temperature
		fmt.Printf("%s == %s, %s == %s\n",
			genconv.Fahrenheit(unit),
			genconv.FToC(genconv.Fahrenheit(unit)),
			genconv.Celsius(unit),
			genconv.CToF(genconv.Celsius(unit)))
		// Length
		fmt.Printf("%s == %s, %s == %s\n",
			genconv.Meter(unit),
			genconv.MeterToFeet(genconv.Meter(unit)),
			genconv.Feet(unit),
			genconv.FeetToMeter(genconv.Feet(unit)))
		// Weight
		fmt.Printf("%s == %s, %s == %s\n",
			genconv.Kilogram(unit),
			genconv.KilogramToPound(genconv.Kilogram(unit)),
			genconv.Pound(unit),
			genconv.PoundToKilogram(genconv.Pound(unit)))
	}
}

// --- NAMES

func names() {
	// Names should always start with a letter and can contain letters, digits, and underscores.
	// The convention in Go is to use camelCase for variable names.
	HTMLScape := "<h1>Names</h1>"

	println(HTMLScape)
}

// --- DECLARATIONS

func declarations() {
	// There are four major kinds of declarations:
	// var, const, func, and type.
	var f = boilingF

	var c = fToC(f)

	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}

// The function fToC converts Fahrenheit to Celsius.
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

// --- VARIABLES

func variables() {
	// A var creates a variable of a particular type, attaches
	// a name to it, and optionally initializes it.
	// * var name type = expression
	// If expression is omitted, the variable is initialized to
	// the zero value for its type, which is 0 for numeric types,
	// false for booleans, and "" for strings, and nil for interfaces
	// and reference types.
	var i, j, k int
	var b, f, s = true, 2.3, "four"
	println(i, j, k, b, f, s) // Notice println() doesn't format the output

	// If we omit the type, the compiler infers it from the
	// initializer. This is called short variable declaration.
	// * name := expression
	// It should be used everytime the type is obvious from
	// the initializer, otherwise we should use the var keyword.

	// KEEP IN MIND THAT := IS A DECLARATION, NOT AN ASSIGNMENT AS =

	p := new(int) // p is a pointer to an int

	fmt.Println(*p) // dereference the pointer to get the value
	fmt.Println(p)  // print the pointer itself

}

// --- TYPES

type Celsius float64
type Fahrenheit float64

// The Stringer interface is used to define a string representation of a type.
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func types() {
	var c Celsius = 100

	fmt.Printf("%g°C == %g°F\n", c, CToF(c))
}

func CToF(c Celsius) Fahrenheit {
	var x float32 = 2.2
	var y int = int(x)
	fmt.Println(x, y)             // 2.2 2
	return Fahrenheit(c*9/5 + 32) // convert Celsius to Fahrenheit (Makes it explicit)
	// We've couldn't done int(2.2) but we can do int(2) since the funtionallity
	// is to make it explicit.
}

// --- PACKAGE INITIALIZATION

// Package level variables are initialized in dependecy order.

// The init() function is called after all package variables are initialized.
// and serves as an initializer for complex initialization.
// It is called before the main() function and is executed only once.
// this may be helpful to avoid costly initialization of variables in
// execution time.

// The main() function is the last function to be initialized.
