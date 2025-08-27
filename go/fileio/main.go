package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func createTestData(size int) string {
	var builder strings.Builder
	for i := 0; i < size; i++ {
		builder.WriteString(fmt.Sprintf("Line %d: %s\n", i, strings.Repeat("x", 50)))
	}
	return builder.String()
}

func testFileWrite(filePath string, data string, iterations int) float64 {
	fmt.Printf("Testing file write (%d bytes)...\n", len(data))
	var times []float64

	for i := 0; i < iterations; i++ {
		testFile := fmt.Sprintf("%s_%d.txt", filePath, i)
		start := time.Now()
		err := os.WriteFile(testFile, []byte(data), 0644)
		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("Error writing file: %v\n", err)
			continue
		}

		timeMs := float64(elapsed.Nanoseconds()) / 1000000.0
		times = append(times, timeMs)

		// 清理文件
		os.Remove(testFile)
		fmt.Printf("Write %d: %.3fms\n", i+1, timeMs)
	}

	var total float64
	for _, t := range times {
		total += t
	}
	avgTime := total / float64(len(times))
	fmt.Printf("Average write time: %.3fms\n\n", avgTime)
	return avgTime
}

func testFileRead(filePath string, iterations int) float64 {
	fmt.Printf("Testing file read...\n")
	var times []float64

	for i := 0; i < iterations; i++ {
		start := time.Now()
		data, err := os.ReadFile(filePath)
		elapsed := time.Since(start)

		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			continue
		}

		timeMs := float64(elapsed.Nanoseconds()) / 1000000.0
		times = append(times, timeMs)
		fmt.Printf("Read %d: %.3fms (%d bytes)\n", i+1, timeMs, len(data))
	}

	var total float64
	for _, t := range times {
		total += t
	}
	avgTime := total / float64(len(times))
	fmt.Printf("Average read time: %.3fms\n\n", avgTime)
	return avgTime
}

func testConcurrentFileOps(basePath string, fileCount int) float64 {
	fmt.Printf("Testing concurrent file operations (%d files)...\n", fileCount)
	data := createTestData(1000)

	start := time.Now()

	// 并发写入
	var wg sync.WaitGroup
	for i := 0; i < fileCount; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fileName := fmt.Sprintf("%s_concurrent_%d.txt", basePath, idx)
			os.WriteFile(fileName, []byte(data), 0644)
		}(i)
	}
	wg.Wait()

	// 并发读取
	for i := 0; i < fileCount; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			fileName := fmt.Sprintf("%s_concurrent_%d.txt", basePath, idx)
			os.ReadFile(fileName)
		}(i)
	}
	wg.Wait()

	elapsed := time.Since(start)
	timeMs := float64(elapsed.Nanoseconds()) / 1000000.0

	// 清理文件
	for i := 0; i < fileCount; i++ {
		fileName := fmt.Sprintf("%s_concurrent_%d.txt", basePath, i)
		os.Remove(fileName)
	}

	fmt.Printf("Concurrent operations completed: %.3fms\n", timeMs)
	fmt.Printf("Average per file: %.3fms\n\n", timeMs/float64(fileCount))

	return timeMs
}

func main() {
	fmt.Println("Go File I/O Tests")
	fmt.Println("=================")

	dataDir := filepath.Join("..", "..", "data")
	os.MkdirAll(dataDir, 0755)

	basePath := filepath.Join(dataDir, "test_go")

	// 测试不同大小的文件
	sizes := []int{1000, 10000, 100000} // 行数

	for _, size := range sizes {
		fmt.Printf("\n=== Testing with %d lines ===\n", size)
		data := createTestData(size)
		filePath := fmt.Sprintf("%s_%d.txt", basePath, size)

		// 写入测试
		testFileWrite(filePath, data, 5)

		// 先创建文件用于读取测试
		os.WriteFile(filePath, []byte(data), 0644)

		// 读取测试
		testFileRead(filePath, 5)

		// 清理测试文件
		os.Remove(filePath)
	}

	// 并发测试
	fmt.Printf("\n=== Concurrent File Operations ===\n")
	testConcurrentFileOps(basePath, 10)
	testConcurrentFileOps(basePath, 50)
}
