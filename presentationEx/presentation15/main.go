package main

import "fmt"

func main() {
	//printing using different formats
	var num = 394
	fmt.Printf("%b\n", num)
	fmt.Printf("%X\n", num)

	//write a program that creates a slice of ints using shorthand notation
	//then prints the slice
	mySlice := []int{1, 2, 3, 4, 5, 6}
	for i, val := range mySlice {
		fmt.Println(i, " - ", val)
	}

	//write a program that creates a slice of strings using shorthand notation
	//then prints the slice
	myString := []string{"I", "am", "a", "string"}
	for _, val := range myString {
		fmt.Print(val, " ")
	}
	fmt.Println()

	//create a slice of ints using make; set the length to 5 and capacity to 10
	var newSlice = make([]int, 5, 10)
	newSlice[0] = 12345
	fmt.Println(newSlice[0])

	//append a slice to a slice
	var slice1, slice2 = []int{1, 2, 3, 4}, []int{5, 6, 7, 8}

	slice1 = append(slice1, slice2...)
	for _, val := range slice1 {
		fmt.Print(val)
	}
	fmt.Println()

	//delete an element from a slice
	slice1 = append(slice1[:2], slice1[3:]...)
	for _, val := range slice1 {
		fmt.Print(val)
	}
	fmt.Println()
	//create a slice then make your program throw an “index out of range” error
	var errorSlice = []int{1, 2, 3, 4}
	fmt.Println(errorSlice[5])

	//create a program that:
	// creates a map using shorthand notation
	var someMap = map[int]string{
		0: "Hello",
		1: "World!",
	}

	// adds an entry to the map
	someMap[2] = "I'm New"

	// changes an entry in the map
	someMap[2] = "So am I"

	// deletes an entry in the map
	delete(someMap, 2)

	// prints all of the entries in the map using range
	for key, val := range someMap {
		fmt.Println(key, " - ", val)
	}

	// prints the len of the map
	fmt.Println(len(someMap))

	// uses the comma ok idiom
	if val, exists := someMap[3]; exists {
		fmt.Println("val->", val)
		fmt.Println(exists)
	} else {
		fmt.Println("val->", val)
		fmt.Println(exists)
	}
	// create a program that:
	// declares a variable of type int using new
	// (note: this is not idiomatic go code to create an int this way)
	var someNewVar = new(int)
	// print out the memory address of the variable
	fmt.Println(someNewVar)
	// print out the value of the variable
	fmt.Println(*someNewVar)
	// true or false:
	// 	new returned a pointer
	// Ans-->TRUE

	// 	create a program that:
	// declares a variable of type string using new
	// (note: this is not idiomatic go code to create a string this way)

	// print out the memory address of the variable

	// print out the value of the variable

	// true or false:
	// new returned a pointer

}
