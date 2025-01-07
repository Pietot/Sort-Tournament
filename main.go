package main

import (
	"fmt"
	sort "gosort/algorithms"
	"math/rand"
	"time"
)

func main() {
	array := generateUniqueRandomList(10000)
	start := time.Now()
	sort.BinaryInsertionSort(append([]int(nil), array...))
	fmt.Println("Time taken for BinaryInsertionSort:", time.Since(start))
	start = time.Now()
	sort.BubbleSortOptimized(append([]int(nil), array...))
	fmt.Println("Time taken for BubbleSortOptimized:", time.Since(start))

}

func generateUniqueRandomList(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}

	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		nums[i], nums[j] = nums[j], nums[i]
	}

	return nums
}
