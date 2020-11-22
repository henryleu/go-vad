package main

import "fmt"

func main() {
	fmt.Println("calculator exercises")
	var a, b, c int
	var product int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	product = a + b - c
	fmt.Printf("%d + %d - %d = %d", a, b, c, product)
	fmt.Println()
}
