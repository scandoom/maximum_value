package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 1 {
		fmt.Println("error: size < 1")
		return nil
	}

	genData := make([]int, size)

	for i := 0; i < size; i++ {
		genData[i] = rand.Int()
	}

	return genData
	// ваш код здесь
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) <= 1 {
		fmt.Println("error: len slice data < 1")
		return 0
	}

	var maxCount = 0
	for _, i := range data {
		if i > maxCount {
			maxCount = i
		}
	}

	return maxCount
	// ваш код здесь
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) <= 1 {
		fmt.Println("error: data < 1")
		return 0
	}

	var wg sync.WaitGroup

	wg.Add(CHUNKS)

	maxData := make([]int, CHUNKS)

	if len(data) < CHUNKS {
		return maximum(data)
	}

	for i := 0; i < CHUNKS; i++ {
		v := len(data) / CHUNKS
		beginningChunk := i * v
		endChunk := beginningChunk + v

		if i == CHUNKS-1 {
			endChunk = len(data)
		}

		partData := data[beginningChunk:endChunk]

		go func(data []int) {
			defer wg.Done()

			maxData[i] = maximum(partData)

		}(partData)
	}

	wg.Wait()

	return maximum(maxData)

}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)

	general := generateRandomElements(SIZE) // ваш код здесь

	fmt.Println("Ищем максимальное значение в один поток")

	start := time.Now() // ваш код здесь
	max := maximum(general)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	start = time.Now()
	max = maxChunks(general) // ваш код здесь
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

}
