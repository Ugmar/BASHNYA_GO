package main

import (
	"sync"
)

func square(a int, sum *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	*sum += a * a
	mu.Unlock()
}

func SearchSumSquare(numbers []int) (sum int) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(len(numbers))

	for _, number := range numbers {
		go square(number, &sum, &wg, &mu)
	}

	wg.Wait()

	return sum
}
