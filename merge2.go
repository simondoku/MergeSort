package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	MaxSize = 1000000
	Runs    = 100
)

var S [MaxSize]int // Global array S for use in mergesort2

// Merge function for mergesort2
func merge2(low, mid, high int, verbose bool) int64 {
	memoryUsage := int64(high-low+1) * int64(4) // Memory for U array (int = 4 bytes)

	U := make([]int, high-low+1)

	i, j, k := low, mid+1, 0

	for i <= mid && j <= high {
		if S[i] <= S[j] {
			U[k] = S[i]
			i++
		} else {
			U[k] = S[j]
			j++
		}
		k++
	}

	for i <= mid {
		U[k] = S[i]
		i++
		k++
	}

	for j <= high {
		U[k] = S[j]
		j++
		k++
	}

	for k, i := 0, low; i <= high; i, k = i+1, k+1 {
		S[i] = U[k]
	}

	if verbose {
		fmt.Printf("After merging [%d..%d]: ", low, high)
		for i := low; i <= high; i++ {
			fmt.Printf("%d ", S[i])
		}
		fmt.Println()
	}

	return memoryUsage
}

// Recursive mergesort2 function with space usage tracking
func mergesort2(low, high int, verbose bool) int64 {
	if low < high {
		mid := (low + high) / 2

		if verbose {
			fmt.Printf("Dividing [%d..%d] into [%d..%d] and [%d..%d]\n", low, high, low, mid, mid+1, high)
		}

		memoryUsage := mergesort2(low, mid, verbose)
		memoryUsage += mergesort2(mid+1, high, verbose)

		memoryUsage += merge2(low, mid, high, verbose)

		return memoryUsage
	}
	return 0
}

// Data generation functions
func generateRandomData(size int) {
	for i := 0; i < size; i++ {
		S[i] = rand.Intn(100) + 1
	}
}

func generateSortedData(size int) {
	for i := 0; i < size; i++ {
		S[i] = i + 1
	}
}

func generateReversedData(size int) {
	for i := 0; i < size; i++ {
		S[i] = size - i
	}
}

func generateNearlySortedData(size int) {
	generateSortedData(size)
	for i := 0; i < size/10; i++ {
		idx1 := rand.Intn(size)
		idx2 := rand.Intn(size)
		S[idx1], S[idx2] = S[idx2], S[idx1]
	}
}

// Run tests for mergesort2
func runSortTests(size int) {
	totalExecutionTimeRandom := int64(0)
	totalMemoryUsageRandom := int64(0)

	totalExecutionTimeSorted := int64(0)
	totalMemoryUsageSorted := int64(0)

	totalExecutionTimeReversed := int64(0)
	totalMemoryUsageReversed := int64(0)

	totalExecutionTimeNearlySorted := int64(0)
	totalMemoryUsageNearlySorted := int64(0)

	verbose := size == 10

	for run := 0; run < Runs; run++ {
		// Random Data
		generateRandomData(size)
		startTime := time.Now().UnixNano()
		memoryUsed := mergesort2(0, size-1, verbose && run == 0)
		elapsedTime := time.Now().UnixNano() - startTime
		totalExecutionTimeRandom += elapsedTime
		totalMemoryUsageRandom += memoryUsed

		// Sorted Data
		generateSortedData(size)
		startTime = time.Now().UnixNano()
		memoryUsed = mergesort2(0, size-1, verbose && run == 0)
		elapsedTime = time.Now().UnixNano() - startTime
		totalExecutionTimeSorted += elapsedTime
		totalMemoryUsageSorted += memoryUsed

		// Reversed Data
		generateReversedData(size)
		startTime = time.Now().UnixNano()
		memoryUsed = mergesort2(0, size-1, verbose && run == 0)
		elapsedTime = time.Now().UnixNano() - startTime
		totalExecutionTimeReversed += elapsedTime
		totalMemoryUsageReversed += memoryUsed

		// Nearly Sorted Data
		generateNearlySortedData(size)
		startTime = time.Now().UnixNano()
		memoryUsed = mergesort2(0, size-1, verbose && run == 0)
		elapsedTime = time.Now().UnixNano() - startTime
		totalExecutionTimeNearlySorted += elapsedTime
		totalMemoryUsageNearlySorted += memoryUsed
	}

	fmt.Printf("mergesort2 - Size: %d\n", size)
	fmt.Printf("  Random Data - Avg Time (ns): %d, Avg Memory (bytes): %d\n", totalExecutionTimeRandom/Runs, totalMemoryUsageRandom/Runs)
	fmt.Printf("  Sorted Data - Avg Time (ns): %d, Avg Memory (bytes): %d\n", totalExecutionTimeSorted/Runs, totalMemoryUsageSorted/Runs)
	fmt.Printf("  Reversed Data - Avg Time (ns): %d, Avg Memory (bytes): %d\n", totalExecutionTimeReversed/Runs, totalMemoryUsageReversed/Runs)
	fmt.Printf("  Nearly Sorted Data - Avg Time (ns): %d, Avg Memory (bytes): %d\n", totalExecutionTimeNearlySorted/Runs, totalMemoryUsageNearlySorted/Runs)
}

func main() {
	sizes := []int{10, 100, 1000, 10000, 100000, 1000000}
	for _, size := range sizes {
		runSortTests(size)
	}
}
