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

	wg.Add(8)

	maxData := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		if len(data) < CHUNKS {
			fmt.Println("error: CHUNKS not correct")
			return 0
		}

		v := len(data) / CHUNKS

		partData := data[i*v : i*v+v]
		if len(partData) <= 1 {
			fmt.Println("error: data < 1")
			return 0
		}

		go func(data []int) {
			defer wg.Done()

			count := maximum(data)

			if count == 0 {
				fmt.Println("error: count = 0")
				return
			}

			maxData[i] = count

		}(partData)
	}

	wg.Wait()

	maxCount := maximum(maxData)
	if maxCount == 0 {
		fmt.Println("errror: maxCount = 0")
		return 0
	}

	return maxCount

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
