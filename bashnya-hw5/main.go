package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	sum := SearchSumSquare(numbers)
	fmt.Println(sum)

	// ChanelWork(10);
	Conveyor(numbers)
}
