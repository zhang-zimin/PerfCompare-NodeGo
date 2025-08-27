package main

import (
	"fmt"
	"time"
)

func fibonacciRecursive(n int) int64 {
	if n <= 1 {
		return int64(n)
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func fibonacciIterative(n int) int64 {
	if n <= 1 {
		return int64(n)
	}
	a, b := int64(0), int64(1)
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func runTest(name string, fn func(int) int64, n int, iterations int) {
	fmt.Printf("\n=== %s (n=%d) ===\n", name, n)
	var times []float64

	for i := 0; i < iterations; i++ {
		start := time.Now()
		result := fn(n)
		elapsed := time.Since(start)
		timeMs := float64(elapsed.Nanoseconds()) / 1000000.0
		times = append(times, timeMs)
		fmt.Printf("Run %d: %.3fms - Result: %d\n", i+1, timeMs, result)
	}

	var total float64
	for _, t := range times {
		total += t
	}
	avgTime := total / float64(len(times))

	minTime := times[0]
	maxTime := times[0]
	for _, t := range times {
		if t < minTime {
			minTime = t
		}
		if t > maxTime {
			maxTime = t
		}
	}

	fmt.Printf("Average: %.3fms\n", avgTime)
	fmt.Printf("Min: %.3fms\n", minTime)
	fmt.Printf("Max: %.3fms\n", maxTime)
}

func main() {
	fmt.Println("Go CPU Intensive Tests")
	fmt.Println("======================")

	// 测试不同规模的斐波那契数列
	testCases := []int{35, 40, 42}

	for _, n := range testCases {
		runTest("Fibonacci Recursive", fibonacciRecursive, n, 5)
		runTest("Fibonacci Iterative", fibonacciIterative, n, 5)
	}

	// 大数计算测试
	fmt.Println("\n=== Large Number Fibonacci (Iterative) ===")
	largeNumbers := []int{1000, 5000, 10000}
	for _, n := range largeNumbers {
		runTest("Large Fibonacci", fibonacciIterative, n, 3)
	}
}
