// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// https://play.golang.org/p/1xUWjHMB3I

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare variables that are set to their zero value.
	var anInteger int
	var isGoodPerson bool
	var aFactor float64

	fmt.Printf("anInteger: %v of type: %T\n", anInteger, anInteger)
	fmt.Printf("anJudgement: %v of type: %T\n", isGoodPerson, isGoodPerson)
	fmt.Printf("anFactor: %v of type: %T\n", aFactor, aFactor)

	// Display the value of those variables.
	fmt.Println()
	// Declare variables and initialize.
	anotherInteger := 10
	anotherJudgement := true
	anotherFactor :=3.14

	// Using the short variable declaration operator.

	// Display the value of those variables.
	fmt.Printf("anInteger: %v of type: %T\n",anotherInteger, anotherInteger)
	fmt.Printf("anJudgement: %v of type: %T\n", anotherJudgement, anotherJudgement)
	fmt.Printf("anFactor: %v of type: %T\n", anotherFactor, anotherFactor)

	// Perform a type conversion.
	anotherFloat := float64(anotherInteger)
	fmt.Println()
	// Display the value of that variable.
	fmt.Printf("anotherInteger:%v has been converted to anotherFloat:%v of type: %T", anotherInteger, anotherFloat, anotherFloat)

}
