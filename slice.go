package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	fmt.Printf("len(a) = %d\n", len(a))
	fmt.Printf("cap(a) = %d\n", cap(a))

	fmt.Println()

	b := make([]int, 0, 8)

	fmt.Printf("len(a) = %d\n", len(b))
	fmt.Printf("cap(a) = %d\n", cap(b))
}
