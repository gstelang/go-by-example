// 3. Variables
package main

import "fmt"

// Variables are explicitly declared and used by the compiler to check type-correctness
func main() {

	var a = "initial"
	fmt.Println(a)

	// declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued.
	var e int
	fmt.Println(e)

	// shorthand for declaring and initializing a variable
	f := "apple"
	fmt.Println(f)
}
