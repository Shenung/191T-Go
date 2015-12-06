package main

import "fmt"

func even(num int) (int, bool) {
	if num%2 == 0 {
		return num, true
	}
	return num, false
}

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func info(name string, age int) {
	fmt.Println(name, " is ", age, " years old.")
}

func dogyear(name string, age int) {
	if age > 25 {
		fmt.Println(name, " is ", age*7, " in dog years and is old.")
	} else {
		fmt.Println(name, " is ", age*7, " in dog years and is not old.")
	}
}

func dogyears2(age int) int {
	return age * 7
}

func year(age int) int {
	return dogyears2(age)
}

func foo(numbers ...int) {
	fmt.Println(numbers)
}

func inside(something int) int {
	return something
}

func another(thing int) int {
	return thing
}

func fib(something int) int {
	if something == 0 {
		return 0
	} else if something == 1 {
		return 1
	} else {
		return fib(something-1) + fib(something-2)
	}
}

func hello() {
	fmt.Println("Hello")
}
func world() {
	fmt.Println("World")
}

func main() {
	// Write a function which takes an integer and returns two values:
	// the integer divided by 2
	// whether or not the integer is even (true, false)
	// For example
	// 	half(1) should return (0, false)
	// 	half(2) should return (1, true).
	var num, truth = even(2)
	fmt.Println(num, ", ", truth)

	num, truth = even(1)
	fmt.Println(num, ", ", truth)

	// Write a function with one variadic parameter that finds the greatest number in a list of numbers.
	greatest := max(4, 7, 9, 123, 543, 23, 435, 53, 125)
	fmt.Println(greatest)

	// Write a program that prints the value of this expression:
	//  (true && false) || (false && true) || !(false && false)
	fmt.Println((true && false) || (false && true) || !(false && false))

	// Write a program that calls a function which takes first name and age
	// then returns a string like this, “John is 27 years old.”
	var name string
	var age int
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	fmt.Println()
	fmt.Print("age: ")
	fmt.Scanln(&age)
	fmt.Println()
	info(name, age)

	// Write a program that calls a function which takes first name and age
	// then returns an int and a bool
	// the int: person’s age * 7 (dog years)
	// the bool: whether or not the person is old (age > 25)
	// use those two returns in a sentence like this,
	// (“John is 140 in dog years and is not old”)
	// or like this, (“Jane is 280 in dog years and is old”)
	dogyear(name, age)

	// Write a program that calls a function which takes age
	// then returns dogYears int which is age * 7
	year(age)

	// Write a program that has variadic parameters
	// use that function in a program
	// Write a program that has variadic parameters
	// use that function in a program, passing in variadic arguments
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()

	// Write a program that uses a func expression
	hi := func() {
		fmt.Println("Hi")
	}
	hi()
	// You wrote a program that uses a func expression
	// now add a print statement that shows
	// the type of the variable to which the function is assigned
	fmt.Printf("%T", hi)

	// create a program that uses closure
	var x = 0
	inc := func() int {
		x++
		return x
	}
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

	// create a func that returns a func
	// use that func in a program
	newfunc := func() func() {
		return func() {
			fmt.Println("Hello")
		}
	}
	fmt.Println(newfunc)

	// create a program that uses a callback
	// (a func is being passed in as an argument)
	fmt.Println(inside(another(5)))

	// The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).
	fib(4)

	// create a program that uses defer
	defer world()
	hello()
}
