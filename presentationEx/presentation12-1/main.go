package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	con := true
	i := 0
	for con {
		i++
		if i == 10 {
			con = false
		}
	}
	j := 0
	for {
		j++
		if j == 100000 {
			break
		}
	}
}
