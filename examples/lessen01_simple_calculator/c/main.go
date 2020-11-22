package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("calculator exercises")
	var (
		n int
		a []int
	)
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &n)

	a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &a[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Print(a[i])
		} else {
			fmt.Printf(" + %d", a[i])
		}
		sum += a[i]
	}
	fmt.Printf(" = %d", sum)
	fmt.Println()
}
