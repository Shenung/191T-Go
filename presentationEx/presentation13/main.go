package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func hash(c string) string {
	h := sha256.New()
	fmt.Fprint(h, c)
	return hex.EncodeToString(h.Sum(nil))
}

func KeysFromFile(name string) (string, string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", "", err
	}
	defer file.Close()

	var secret, access string
	_, err = fmt.Fscan(file, &secret, &access) // <-- here
	return secret, access, err
}

func greet(a string, people ...string) {
	fmt.Println(a)
	for _, val := range people {
		fmt.Println(val)
	}
}

func sum(numbers ...int) {
	var result int
	for _, val := range numbers {
		result = val + result
	}
	fmt.Println(result)
}

func main() {
	// fmt.Println(hash("LOOK AT ME IM A HASH AND FPRINT IS BEING USED"))
	//
	// i := 23
	// s := fmt.Sprint("Hello World. I'm ", i, " years old.")
	// fmt.Print(s)

	// answers := "1 0 8"
	// var a, b, c int
	// fmt.Sscan(answers, &a, &b, &c)
	// fmt.Println(a, b, c)

	// fmt.Print(KeysFromFile("./example.txt"))

	s := []string{"me", "you", "who"}
	greet("hello: ", s...)

	n := []int{1, 2, 3}
	sum(n...)

	var inp int
	fmt.Println("give me an input:")
	fmt.Scanln(&inp)
	fmt.Println(inp)
}
