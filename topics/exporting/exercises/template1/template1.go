// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create a package named toy with a single exported struct type named Toy. Add
// the exported fields Name and Weight. Then add two unexported fields named
// onHand and sold. Declare a factory function called New to create values of
// type toy and accept parameters for the exported fields. Then declare methods
// that return and update values for the unexported fields.
//
// Create a program that imports the toy package. Use the New function to create a
// value of type toy. Then use the methods to set the counts and display the
// field values of that toy value.
package main

// Add imports.
import (
	"fmt"

	"github.com/ardanlabs/gotraining/topics/exporting/exercises/template1/toy"
)

// main is the entry point for the application.
func main() {
	// Use the New function from the toy package to create a value of
	// type toy.
	t := toy.New("Yo-yo", 1)
	// Use the methods from the toy value to set some initialize
	// values.
	t.UpdateOnHand(20)
	t.UpdateSold(0)

	// Display each field separately from the toy value.
	fmt.Println("toy name:", t.Name)
	fmt.Println("toy weight", t.Weight)
	fmt.Println("count on hand", t.OnHand())
	fmt.Println("count sold", t.Sold())
}
