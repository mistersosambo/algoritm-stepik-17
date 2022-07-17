package main

import "fmt"

func main() {
	var a int
	res := 1
	fmt.Scan(&a)
	fmt.Print(res, " ")
	for i := 0; i < a; i++ {
		res *= 2
		if res <= a {
			fmt.Print(res, " ")
		} else {
			return
		}
	}
}
