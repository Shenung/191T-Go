package main

import "fmt"

func main() {
	var num1, num2, result int

	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter the second number: ")
	fmt.Scanln(&num2)
	result = num1 % num2
	fmt.Println(result)

	for i := 1; i <= 1000; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	for i := 0; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println(i, "FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "FIZZ")
		} else if i%5 == 0 {
			fmt.Println(i, "BUZZ")
		} else {
			fmt.Println(i)
		}
	}

	counter := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println(counter)
}
