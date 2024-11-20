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

// Global variable to track memory usage
var totalMemory int64

// Helper function to print an array
func printArray(S []int) {
	for _, val := range S {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}

// Merge function for Merge Sort 3
func merge3(low, mid, high int, S, temp []int) {
	// Memory allocation for temporary variables during merging
	memoryAllocated := (high - low + 1) * 4 // Assuming each int is 4 bytes
	totalMemory += int64(memoryAllocated)

	i, j, k := low, mid+1, low

	for i <= mid && j <= high {
		if S[i] <= S[j] {
			temp[k] = S[i]
			i++
		} else {
			temp[k] = S[j]
			j++
		}
		k++
	}

	for i <= mid {
		temp[k] = S[i]
		i++
		k++
	}

	for j <= high {
		temp[k] = S[j]
		j++
		k++
	}

	for i = low; i <= high; i++ {
		S[i] = temp[i]
	}
}

// Iterative Merge Sort 3
func mergesort3(n int, S, temp []int, verbose bool) {
	size := 1

	if verbose {
		fmt.Println("Initial array:")
		printArray(S)
	}

	for size < n {
		if verbose {
			fmt.Printf("\nMerging subarrays of size %d:\n", size)
		}

		for low := 0; low < n; low += 2 * size {
			mid := low + size - 1
			high := min(low+2*size-1, n-1)

			if mid < n-1 {
				merge3(low, mid, high, S, temp)

				if verbose {
					fmt.Printf("After merging [%d..%d]: ", low, high)
					printArray(S)
				}
			}
		}

		size *= 2
	}

	if verbose {
		fmt.Println("\nFinal sorted array:")
		printArray(S)
	}
}

// Utility function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Data generation functions
func generateRandomData(size int) []int {
	data := make([]int, size)
	totalMemory += int64(len(data) * 4) // Memory allocated for array
	for i := range data {
		data[i] = rand.Intn(100) + 1
	}
	return data
}

func generateSortedData(size int) []int {
	data := make([]int, size)
	totalMemory += int64(len(data) * 4) // Memory allocated for array
	for i := range data {
		data[i] = i + 1
	}
	return data
}

func generateReversedData(size int) []int {
	data := make([]int, size)
	totalMemory += int64(len(data) * 4) // Memory allocated for array
	for i := range data {
		data[i] = size - i
	}
	return data
}

func generateNearlySortedData(size int) []int {
	data := generateSortedData(size)
	for i := 0; i < size/10; i++ {
		idx1, idx2 := rand.Intn(size), rand.Intn(size)
		data[idx1], data[idx2] = data[idx2], data[idx1]
	}
	return data
}

// Run tests for Merge Sort 3
func runTest(size int, verbose bool) {
	dataGenerators := []func(int) []int{
		generateRandomData,
		generateSortedData,
		generateReversedData,
		generateNearlySortedData,
	}
	dataTypeNames := []string{
		"Random Data",
		"Sorted Data",
		"Reversed Data",
		"Nearly Sorted Data",
	}

	for i, gen := range dataGenerators {
		totalTime := int64(0)
		totalMemory = 0 // Reset memory tracker

		for run := 0; run < Runs; run++ {
			S := gen(size)
			temp := make([]int, size)
			totalMemory += int64(len(temp) * 4) // Memory allocated for temp array

			startTime := time.Now().UnixNano()
			mergesort3(size, S, temp, verbose && run == 0)
			elapsedTime := time.Now().UnixNano() - startTime

			totalTime += elapsedTime
		}

		avgTime := totalTime / int64(Runs)
		avgMemory := totalMemory / int64(Runs)
		fmt.Printf("Size: %d, %s - Average Execution Time (ns): %d, Average Memory Usage (bytes): %d\n",
			size, dataTypeNames[i], avgTime, avgMemory)
	}
}

func main() {
	sizes := []int{10, 100, 1000, 10000, 100000, 1000000}

	rand.Seed(time.Now().UnixNano())

	for _, size := range sizes {
		verbose := size == 10
		runTest(size, verbose)
	}
}
