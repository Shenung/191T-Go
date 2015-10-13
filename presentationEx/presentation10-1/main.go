package main

import (
	"fmt"

	"./names"
)

func main() {
	fmt.Println(names.MyName)
	fmt.Println(names.yourName) //<---doesn't work because the y in yourName is in lower case so it cannot be exported to be used outside of its own file
}
