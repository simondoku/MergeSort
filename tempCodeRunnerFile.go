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
        startTime := time.Now()
        memoryUsed := mergesort2(0, size-1, verbose && run == 0)
        elapsedTime := time.Since(startTime).Milliseconds()
        totalExecutionTimeRandom += elapsedTime
        totalMemoryUsageRandom += memoryUsed

        // Sorted Data
        generateSortedData(size)
        startTime = time.Now()
        memoryUsed = mergesort2(0, size-1, verbose && run == 0)
        elapsedTime = time.Since(startTime).Milliseconds()
        totalExecutionTimeSorted += elapsedTime
        totalMemoryUsageSorted += memoryUsed

        // Reversed Data
        generateReversedData(size)
        startTime = time.Now()
        memoryUsed = mergesort2(0, size-1, verbose && run == 0)
        elapsedTime = time.Since(startTime).Milliseconds()
        totalExecutionTimeReversed += elapsedTime
        totalMemoryUsageReversed += memoryUsed

        // Nearly Sorted Data
        generateNearlySortedData(size)
        startTime = time.Now()
        memoryUsed = mergesort2(0, size-1, verbose && run == 0)
        elapsedTime = time.Since(startTime).Milliseconds()
        totalExecutionTimeNearlySorted += elapsedTime
        totalMemoryUsageNearlySorted += memoryUsed
    }

    fmt.Printf("mergesort2 - Size: %d\n", size)
    fmt.Printf("  Random Data - Avg Time (ms): %d, Avg Memory (bytes): %d\n", totalExecutionTimeRandom/Runs, totalMemoryUsageRandom/Runs)
    fmt.Printf("  Sorted Data - Avg Time (ms): %d, Avg Memory (bytes): %d\n", totalExecutionTimeSorted/Runs, totalMemoryUsageSorted/Runs)
    fmt.Printf("  Reversed Data - Avg Time (ms): %d, Avg Memory (bytes): %d\n", totalExecutionTimeReversed/Runs, totalMemoryUsageReversed/Runs)
    fmt.Printf("  Nearly Sorted Data - Avg Time (ms): %d, Avg Memory (bytes): %d\n", totalExecutionTimeNearlySorted/Runs, totalMemoryUsageNearlySorted/Runs)
}

func main() {
    sizes := []int{10, 100, 1000, 10000, 100000, 1000000}
    for _, size := range sizes {
        runSortTests(size)
    }
}
