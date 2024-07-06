package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	// randomize to better validate the printing functionality
	for i := range nums {
		randI := r.Intn(len(nums))
		nums[i], nums[randI] = nums[randI], nums[i]
	}

	for _, n := range nums {
		numType := "even"

		if n%2 == 1 {
			numType = "odd"
		}

		fmt.Printf("Num %v is %v\n", n, numType)
	}
}
