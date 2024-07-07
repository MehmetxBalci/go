package main

import (
	"fmt" // Package for formatted I/O operations
)

func main() {

	// Using fmt package to print to the screen
	fmt.Println("Hello World")

	// Single-line comment (same as in Java)

	/*
		Multi-line comment (same as in Java)
	*/

	// Variables
	var x = 5         // In Go, variables can be declared without specifying the type
	var y, z = 10, 15 // Multiple variables can be assigned in a single line
	fmt.Println(x, y, z)

	// Constants
	const PI = 3.14 // We use const instead of the final keyword in Java
	fmt.Println(PI)

	// Output
	var i = "Hello"
	var j = 15
	fmt.Print(i, y, z)                   // Print variables without spaces between them
	fmt.Print(i, "\n", y, "\n", z, "\n") // Print each variable on a new line using the "\n" character
	fmt.Println(i, y, z)                 // Print variables with spaces between them and each on a new line

	//General usage
	fmt.Printf("i has value: %v and type %T\n", i, i)     // "%v" prints the value in a default format, "%T" prints the type
	fmt.Printf("i has value: %#v and print %% sign\n", i) // "%#v" prints the value in Go-syntax format, "%%" prints the "%" sign

	//For Integers
	fmt.Printf("%b\n", j)   // "%b" prints in binary
	fmt.Printf("%d\n", j)   // "%d" prints in decimal
	fmt.Printf("%+d\n", j)  // "%+d" prints in decimal and shows the sign
	fmt.Printf("%o\n", j)   // "%o" prints in octal
	fmt.Printf("%O\n", j)   // "%O" prints in octal and adds "0o" prefix
	fmt.Printf("%x\n", j)   // "%x" prints in hexadecimal (lowercase)
	fmt.Printf("%X\n", j)   // "%X" prints in hexadecimal (uppercase)
	fmt.Printf("%#x\n", j)  // "%#x" prints in hexadecimal and adds "0x" prefix
	fmt.Printf("%4d\n", j)  // "%4d" prints with a width of 4 characters, right-aligned with spaces
	fmt.Printf("%-4d\n", j) // "%-4d" prints with a width of 4 characters, left-aligned with spaces
	fmt.Printf("%04d\n", j) // "%04d" prints with a width of 4 characters, padded with zeros

	//For Strings
	fmt.Printf("%s\n", i)   // "%s" prints the value as plain text
	fmt.Printf("%q\n", i)   // "%q" prints the value in double quotes
	fmt.Printf("%8s\n", i)  // "%8s" prints the value right-aligned with a width of 8 characters
	fmt.Printf("%-8s\n", i) // "%-8s" prints the value left-aligned with a width of 8 characters
	fmt.Printf("%x\n", i)   // "%x" prints the hex dump of the byte values
	fmt.Printf("% x\n", i)  // "% x" prints the hex dump of the byte values with spaces

	//For Boolean
	var boolt bool = true
	fmt.Printf("%t\n", boolt) // "%t" prints the boolean value as true or false (same as "%v")

	// For Float
	var k = 3.141
	fmt.Printf("%e\n", k)    // "%e" prints in scientific notation (with 'e' for the exponent)
	fmt.Printf("%f\n", k)    // "%f" prints in decimal notation without exponent
	fmt.Printf("%.2f\n", k)  // "%.2f" prints with default width and 2 decimal places
	fmt.Printf("%6.2f\n", k) // "%6.2f" prints with a width of 6 characters and 2 decimal places
	fmt.Printf("%g\n", k)    // "%g" prints in the most compact form, using exponent only if necessary
}
