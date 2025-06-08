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
	if size <= 0 {
		fmt.Println("size must be greater than zero")
		return []int{}
	}
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Int()
		if arr[i] == 0 {
			arr[i]++
		}
	}
	return arr
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 || data == nil {
		return 0
	}
	var res = data[0]
	for i := range data {
		if data[i] > res {
			res = data[i]
		}
	}
	return res
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int, chunks int) int {
	var wg sync.WaitGroup
	maxArr := make([]int, chunks)

	wg.Add(chunks)
	for i := 0; i < chunks; i++ {
		go func(chunk []int) {
			maxArr[i] = maximum(chunk)
			wg.Done()
		}(data[i*(len(data)/chunks) : (i+1)*(len(data)/chunks)])
	}
	wg.Wait()

	tail := len(data) % chunks
	if tail != 0 {
		maxArr = append(maxArr, maximum(data[len(data)-tail:]))
	}
	return maximum(maxArr)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	arr := generateRandomElements(SIZE)
	if len(arr) == 0 {
		return
	}
	fmt.Println("------------------------------------------")

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(arr)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
	fmt.Println("------------------------------------------")

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(arr, CHUNKS)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
