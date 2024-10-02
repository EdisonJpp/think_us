package main

import (
	"fmt"
)

func fn(val []int) []int {
	i := 0

	for i < len(val) {
		acc := 0

		for j := i; j < len(val); j++ {
			acc += val[j]

			if acc == 0 {
				val = append(val[:i], val[j+1:]...)
				break
			}
		}

		if acc != 0 {
			i++
		}
	}

	return val
}

func main() {
	val := []int{3, 4, -7, 5, -6, 2, 5, -1, 8}
	result := fn(val)
	fmt.Println(result)
}
