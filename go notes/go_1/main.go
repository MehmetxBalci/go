package main

import (
	"fmt"
)

func main() {
	// Data Types
	//For Integers
	var q int = 20      // Platform-dependent: ranges from -2147483648 to 2147483647 on 32-bit systems, and from -9223372036854775808 to 9223372036854775807 on 64-bit systems.
	var w int8 = 120    // Ranges from -128 to 127
	var e int16 = 4200  // Ranges from -32768 to 32767 (int32 and int64 can also be used)
	var r uint = 500    // Platform-dependent: ranges from 0 to 4294967295 on 32-bit systems, and from 0 to 18446744073709551615 on 64-bit systems. (Only used for positive integers)
	var t uint8 = 242   // Ranges from 0 to 255
	var y uint16 = 6558 // Ranges from 0 to 65535
	fmt.Println(q, w, e, r, t, y)
	// For String
	var u string = "Hello" // Only one type of declaration
	fmt.Println(u)
	// For Boolean
	var o bool = true // Only one type of declaration
	fmt.Println(o)
	// For Float
	var p float32 = 3.14     // Ranges from -3.4e+38 to 3.4e+38
	var a float64 = 1.7e+302 // Ranges from -1.7e+308 to +1.7e+308
	fmt.Println(p, a)
	// Arrays
	var cars = [8]string{"Volkswagen", "Audi", "Seat", "Cupra", "Skoda", "Bentley", "Lamborghini", "Scania"} // To create an array -> var array_name = [length]type{values}
	var test = [...]string{"a", "b", "c"}                                                                    // Creating an array without specifying the length
	var arr1 = [5]int{}                                                                                      // If you create an array without assigning values, all values will be 0 (same as in Java)
	var arr2 = [5]int{1, 2}                                                                                  // You can assign only the first values, the rest will be 0
	var arr3 = [5]int{0: 9, 4: 16}                                                                           // You can assign values to specific indexes, the rest will be 0
	fmt.Println(cars)                                                                                        // Prints all array values
	fmt.Println(cars[1])                                                                                     // Prints only the 2nd element (starts counting from 0 like in Java)
	fmt.Println(test)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	cars[7] = "Porsche" // This way, you can change the 8th element
	fmt.Println(cars)
	fmt.Println(len(test)) // len() is used to find the length
	// Slices
	myslice := []int{} // Creating an empty slice (similar to ArrayList in Java)
	slicewithvariable := []int{
		9, 16, 42} // Creating a slice with initial values
	fmt.Println("length of slice:", len(myslice))
	fmt.Println("capapcity of slice", cap(slicewithvariable)) // cap() is used to find the capacity
	fmt.Println(myslice)
	slicefromarray := cars[3:6] // This way, you can create a slice from an array. Starts from index 3 to 6 (not included) (similar to substring in Java)
	fmt.Println("length of slice:", len(slicefromarray))
	fmt.Println("capacity of slice", cap(slicefromarray)) // The capacity of the slice: from the start index to the capacity of the original array
	fmt.Println(slicefromarray)
	arraywithfunc := make([]int, 5) // slice_name := make([]type, length, capacity). If capacity is not specified, it is equal to the length
	fmt.Println(arraywithfunc)
	fmt.Println(len(arraywithfunc))
	fmt.Println(cap(arraywithfunc))
	arraywithfunc = append(arraywithfunc, 5)                    // You can add elements to a slice using append().
	arraywithfunc = append(arraywithfunc, slicewithvariable...) // You can merge two slices using append() by adding ... at the end.
	fmt.Println(arraywithfunc)
	fmt.Println(len(arraywithfunc))
	fmt.Println(cap(arraywithfunc))
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	neededNumbers := numbers[:len(numbers)-10]     // Removing the last 10 elements and assigning the remaining elements to another slice
	numbersCopy := make([]int, len(neededNumbers)) // Creating another slice of the same length to copy the needed elements
	copy(numbersCopy, neededNumbers)               // Using copy(destination, source) to copy the needed elements to another slice.
}
