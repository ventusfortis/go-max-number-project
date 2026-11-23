package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 10_000_000
	CHUNKS = 8
)

// generateRandomElements creates size positive pseudo-random integers.
// A non-positive size produces an empty slice.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	elements := make([]int, size)
	for i := range elements {
		elements[i] = rand.Intn(1_000_000) + 1
	}
	return elements
}

// maximum returns the largest element. For an empty slice it returns 0.
// This is safe because the task works with positive integers.
func maximum(elements []int) int {
	if len(elements) == 0 {
		return 0
	}

	maxValue := elements[0]
	for _, value := range elements[1:] {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

// maxChunks finds the maximum using up to CHUNKS goroutines.
// The range boundaries distribute all elements, including any remainder.
func maxChunks(elements []int) int {
	if len(elements) == 0 {
		return 0
	}

	chunkCount := CHUNKS
	if len(elements) < chunkCount {
		chunkCount = len(elements)
	}

	maxValues := make([]int, chunkCount)
	var wg sync.WaitGroup
	wg.Add(chunkCount)

	for i := 0; i < chunkCount; i++ {
		start := i * len(elements) / chunkCount
		end := (i + 1) * len(elements) / chunkCount

		go func(index int, chunk []int) {
			defer wg.Done()
			maxValues[index] = maximum(chunk)
		}(i, elements[start:end])
	}

	wg.Wait()
	return maximum(maxValues)
}

func main() {
	elements := generateRandomElements(SIZE)

	start := time.Now()
	sequentialMax := maximum(elements)
	sequentialDuration := time.Since(start)

	start = time.Now()
	parallelMax := maxChunks(elements)
	parallelDuration := time.Since(start)

	fmt.Printf("Количество элементов: %d\n", len(elements))
	fmt.Printf("Максимум (один поток): %d\n", sequentialMax)
	fmt.Printf("Время (один поток): %d мкс\n", sequentialDuration.Microseconds())
	fmt.Printf("Максимум (8 горутин): %d\n", parallelMax)
	fmt.Printf("Время (8 горутин): %d мкс\n", parallelDuration.Microseconds())
}
