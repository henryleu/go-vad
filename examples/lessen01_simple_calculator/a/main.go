package main

import "fmt"

func main() {
	fmt.Println("calculator exercises")
	var a, b int
	var product int
	fmt.Scan(&a)
	fmt.Scan(&b)
	product = a + b
	fmt.Printf("%d + %d = %d", a, b, product)
	fmt.Println()
}
