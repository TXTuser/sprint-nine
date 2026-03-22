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

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	result := make([]int, size)
	var wg sync.WaitGroup

	chunkSize := size / CHUNKS
	if chunkSize == 0 {
		chunkSize = size
	}

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == CHUNKS-1 {
			end = size
		}

		if start >= size {
			break
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()

			source := rand.NewSource(time.Now().UnixNano() + int64(start))
			r := rand.New(source)

			for j := start; j < end; j++ {
				result[j] = r.Int()
			}
		}(start, end)
	}
	wg.Wait()
	return result
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]

	for i := 0; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}

	return max
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	var wg sync.WaitGroup

	maxVal := make([]int, CHUNKS)

	chunkSize := len(data) / CHUNKS

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize

		if i == CHUNKS-1 {
			end = len(data)
		}

		go func(idx int, slice []int) {
			defer wg.Done()

			max := maximum(slice)

			maxVal[idx] = max
		}(i, data[start:end])
	}

	wg.Wait()
	result := maximum(maxVal)
	return result
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")

	startSingle := time.Now()

	maxSingle := maximum(data)

	elapsedSingle := time.Since(startSingle).Nanoseconds() / 1000

	fmt.Printf("Максимальное значение элемента: %d\n", maxSingle)
	fmt.Printf("Время поиска: %d мкс\n", elapsedSingle)

	fmt.Println() // gap

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)

	startMulti := time.Now()

	maxMulti := maxChunks(data)

	elapsedMulti := time.Since(startMulti).Nanoseconds() / 1000

	fmt.Printf("Максимальное значение элемента: %d\n", maxMulti)
	fmt.Printf("Время поиска: %d мкс\n", elapsedMulti)

	if maxSingle != maxMulti {
		fmt.Printf("\nВНИМАНИЕ: Результаты отличаются! (%d vs %d)\n", maxSingle, maxMulti)
	}
}