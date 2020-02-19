package main

import "fmt"

func MaxMultiple(d, b int) int {
	for a := b; a >= d; a-- {
		if a%d == 0 {
			return a
		}
	}

	return 0
}

func main() {
	fmt.Println(MaxMultiple(2, 7))
	fmt.Println(MaxMultiple(10, 50))
	fmt.Println(MaxMultiple(37, 200))
}
