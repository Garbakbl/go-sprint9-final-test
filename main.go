package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE = 100_000_000

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
	var res int
	for i := range data {
		if data[i] > res {
			res = data[i]
		}
	}
	return res
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var (
		matrix [CHUNKS][]int
		maxArr []int
		wg     sync.WaitGroup
		mu     sync.Mutex
	)

	tail := SIZE % CHUNKS
	if tail != 0 {
		maxArr = append(maxArr, maximum(data[len(data)-tail:]))
	}

	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ {
		go func() {
			matrix[i] = data[i*(SIZE/CHUNKS) : (i+1)*(SIZE/CHUNKS)]
			mu.Lock()
			maxArr = append(maxArr, maximum(matrix[i]))
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
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
	max = maxChunks(arr)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
