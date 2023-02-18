package main

import "fmt"

func multiAdd(nums ...int) (int, error) {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum, nil
}

func main() {
	sum, _ := multiAdd(1, 2, 3, 4, 5)
	fmt.Printf("sum=%d", sum)
}
