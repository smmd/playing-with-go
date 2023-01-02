package main

import "fmt"

func main() {
	println(change([]int{2, 2, 1}))
}

func change(input []int) string {
	total := 0

	for _, coin := range input {
		total = total + coin
	}

	coins := map[int]int{}
	denominations := []int{10, 5, 2, 1}

	for _, coinValue := range denominations {
		coins[coinValue] = total / coinValue
		total = total % coinValue

		if total == 0 {
			break
		}
	}

	return fmt.Sprintf("You have %v", coins)
}
