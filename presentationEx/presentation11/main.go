package main

import "fmt"

func main() {
	const age = 23
	fmt.Println("age: ", age)

	for i := 0; i < 3; i++ {
		const (
			num = iota
			num2
		)
		fmt.Println(num)
		fmt.Println(num2)
	}

	var name string
	fmt.Println("What is your name? ")
	fmt.Scanln(&name)
	fmt.Println(&name)

	var Pname *string = &name
	fmt.Println(*Pname)
}
