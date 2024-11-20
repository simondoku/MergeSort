package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

const (
	MaxSize = 1000000
	Runs    = 100
)

type Record struct {
	Key  int
	Link int
}

var S [MaxSize]Record

// Print the current state of a linked list
func printList(start int) {
	index := start
	for index != -1 {
		fmt.Printf("[%d | %d] -> ", S[index].Key, S[index].Link)
		index = S[index].Link
	}
	fmt.Println("NULL")
}

// Recursive mergesort4 function with space tracking
func mergesort4(low, high int, mergedList *int, verbose bool, spaceUsed *int) {
	if low == high {
		*mergedList = low
		S[low].Link = -1 // -1 indicates the end of the list
		if verbose {
			fmt.Printf("Single element at position %d: [%d | -1]\n", low, S[low].Key)
		}
	} else {
		mid := (low + high) / 2

		if verbose {
			fmt.Printf("\nDividing range [%d to %d]:\n", low, high)
			fmt.Printf("Left part: [%d to %d]\n", low, mid)
			fmt.Printf("Right part: [%d to %d]\n", mid+1, high)
		}

		var list1, list2 int
		mergesort4(low, mid, &list1, verbose, spaceUsed)
		mergesort4(mid+1, high, &list2, verbose, spaceUsed)

		if verbose {
			fmt.Printf("\nMerging sorted ranges [%d to %d] and [%d to %d]:\n", low, mid, mid+1, high)
			fmt.Println("Left list before merge:")
			printList(list1)
			fmt.Println("Right list before merge:")
			printList(list2)
		}

		merge4(list1, list2, mergedList, verbose, spaceUsed)
	}

	if verbose {
		fmt.Printf("After merging range [%d to %d]: ", low, high)
		printList(*mergedList)
	}
}

// Merge function for mergesort4 with space tracking
func merge4(list1, list2 int, mergedList *int, verbose bool, spaceUsed *int) {
	var lastSorted int

	if S[list1].Key < S[list2].Key {
		*mergedList = list1
		list1 = S[list1].Link
	} else {
		*mergedList = list2
		list2 = S[list2].Link
	}

	lastSorted = *mergedList
	*spaceUsed += int(unsafe.Sizeof(S[0].Link))

	for list1 != -1 && list2 != -1 {
		if S[list1].Key < S[list2].Key {
			S[lastSorted].Link = list1
			lastSorted = list1
			list1 = S[list1].Link
		} else {
			S[lastSorted].Link = list2
			lastSorted = list2
			list2 = S[list2].Link
		}
		*spaceUsed += int(unsafe.Sizeof(S[0].Link))
	}

	if list1 == -1 {
		S[lastSorted].Link = list2
	} else {
		S[lastSorted].Link = list1
	}

	if verbose {
		fmt.Println("Merged current lists into:")
		printList(*mergedList)
	}
}

// Generate random data
func generateRandomData(size int) {
	for i := 0; i < size; i++ {
		S[i] = Record{Key: rand.Intn(100) + 1, Link: -1}
	}
}

// Generate sorted data
func generateSortedData(size int) {
	for i := 0; i < size; i++ {
		S[i] = Record{Key: i + 1, Link: -1}
	}
}

// Generate reversed data
func generateReversedData(size int) {
	for i := 0; i < size; i++ {
		S[i] = Record{Key: size - i, Link: -1}
	}
}

// Generate nearly sorted data
func generateNearlySortedData(size int) {
	generateSortedData(size)
	for i := 0; i < size/10; i++ {
		idx1, idx2 := rand.Intn(size), rand.Intn(size)
		S[idx1].Key, S[idx2].Key = S[idx2].Key, S[idx1].Key
	}
}

// Get the current time in milliseconds
func getTimeInMs() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// Function to test sorting for all dataset types with space tracking
func runTest(size int) {
	dataGenerators := []func(int){
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
		fmt.Printf("\nTesting %s with size %d:\n", dataTypeNames[i], size)
		totalTime := int64(0)
		totalSpace := 0

		for run := 0; run < Runs; run++ {
			gen(size)
			startTime := time.Now().UnixNano() // Start time in nanoseconds
			var listFront int
			spaceUsed := 0
			mergesort4(0, size-1, &listFront, size == 10 && run == 0, &spaceUsed)
			elapsedTime := time.Now().UnixNano() - startTime // Elapsed time in nanoseconds

			totalTime += elapsedTime
			totalSpace += spaceUsed
		}

		fmt.Printf("Average time for %s with size %d: %d ns\n", dataTypeNames[i], size, totalTime/Runs)
		fmt.Printf("Average space allocated for %s with size %d: %d bytes\n", dataTypeNames[i], size, totalSpace/Runs)
	}
}

func main() {
	sizes := []int{10, 100, 1000, 10000, 100000, 1000000}
	rand.Seed(time.Now().UnixNano())

	for _, size := range sizes {
		runTest(size)
	}
}
