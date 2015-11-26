package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
)

func main() {
	var something = rand.Seed(42)
	fmt.Println(something)

	var x int = 4
	var y float64 = 10.222
	fmt.Println(y + float64(x))
	fmt.Println(int(y))

	fmt.Println(string([]byte{':', 'D'}))
	fmt.Println([]byte(";)"))

	var max = "This is suppose to be used for the max length"
	var s = "-42"
	fmt.Print(strconv.Atoi(s))
	for i := 0; i < len(max); i++ {
		fmt.Println("Wait, give me", strconv.Itoa(i), "mins.")
		fmt.Print("\tlook im escaping\n")
		fmt.Println("and im a rune from", i, ": ", max[i])
		fmt.Println("watch me slice this: ", max[:i])
		fmt.Println("and now to stitch it back up: ", max[:i]+max[i:])
	}

	var num float64
	fmt.Printf("%T", num)
	fmt.Println(reflect.TypeOf(num))
	fmt.Println("give me a number to round up: ")
	fmt.Scanln(&num)
	fmt.Println("rounded number: ", math.Ceil(num))

	fmt.Println("Find the type of a var at golang.org/pkg/reflect")
}
