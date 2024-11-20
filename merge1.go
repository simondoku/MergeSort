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

type KeyType int

// Merge function with space tracking
func merge(h, m int, U, V, S []KeyType) int64 {
	memoryUsage := int64(h*4 + m*4) // Estimate memory usage in bytes (KeyType assumed to be 4 bytes)

	i, j, k := 0, 0, 0
	for i < h && j < m {
		if U[i] < V[j] {
			S[k] = U[i]
			i++
		} else {
			S[k] = V[j]
			j++
		}
		k++
	}

	for i < h {
		S[k] = U[i]
		i++
		k++
	}

	for j < m {
		S[k] = V[j]
		j++
		k++
	}

	return memoryUsage
}

// Recursive Merge Sort function with space usage tracking
func myMergeSort(n int, S []KeyType, verbose bool) int64 {
	if n > 1 {
		h := n / 2
		m := n - h

		U := make([]KeyType, h)
		V := make([]KeyType, m)

		copy(U, S[:h])
		copy(V, S[h:])

		if verbose {
			fmt.Printf("Dividing: Left half: %v | Right half: %v\n", U, V)
		}

		memoryUsage := myMergeSort(h, U, verbose)
		memoryUsage += myMergeSort(m, V, verbose)
		memoryUsage += merge(h, m, U, V, S)

		if verbose {
			fmt.Printf("Merging: %v\n", S)
		}

		return memoryUsage
	}
	return 0
}

// Data generation functions
func generateRandomData(size int) []KeyType {
	data := make([]KeyType, size)
	for i := range data {
		data[i] = KeyType(rand.Intn(100) + 1)
	}
	return data
}

func generateSortedData(size int) []KeyType {
	data := make([]KeyType, size)
	for i := range data {
		data[i] = KeyType(i + 1)
	}
	return data
}

func generateReversedData(size int) []KeyType {
	data := make([]KeyType, size)
	for i := range data {
		data[i] = KeyType(size - i)
	}
	return data
}

func generateNearlySortedData(size int) []KeyType {
	data := generateSortedData(size)
	for i := 0; i < size/10; i++ {
		idx1, idx2 := rand.Intn(size), rand.Intn(size)
		data[idx1], data[idx2] = data[idx2], data[idx1]
	}
	return data
}

// Run tests for different dataset types
func runTest(size int) {
	dataGenerators := []func(int) []KeyType{
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

	verbose := size == 10

	for i, gen := range dataGenerators {
		totalExecutionTime := int64(0)
		totalMemoryUsage := int64(0)

		for run := 0; run < Runs; run++ {
			data := gen(size)

			if verbose && run == 0 {
				fmt.Printf("\nInitial %s:\n%v\n", dataTypeNames[i], data)
			}

			startTime := time.Now().UnixNano()
			memoryUsed := myMergeSort(size, data, verbose && run == 0)
			elapsedTime := time.Now().UnixNano() - startTime

			totalExecutionTime += elapsedTime
			totalMemoryUsage += memoryUsed

			if verbose && run == 0 {
				fmt.Printf("Sorted %s:\n%v\n", dataTypeNames[i], data)
			}
		}

		avgExecutionTime := totalExecutionTime / Runs
		avgMemoryUsage := totalMemoryUsage / Runs

		fmt.Printf("Size: %d, %s - Average Execution Time (ns): %d, Average Space Allocated (bytes): %d\n",
			size, dataTypeNames[i], avgExecutionTime, avgMemoryUsage)
	}
}

func main() {
	sizes := []int{10, 100, 1000, 10000, 100000, 1000000}
	rand.Seed(time.Now().UnixNano())

	for _, size := range sizes {
		fmt.Printf("\nRunning tests for size: %d\n", size)
		runTest(size)
	}
}
