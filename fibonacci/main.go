package main

import (
	"os"
	"runtime/trace"
)

// return N Fibonacci digit
func fib(n int) int64 {
	if n == 1 || n == 2 {
		return 1
	}
	cache := make(map[int]int64, n)
	cache[1], cache[2] = 1, 1
	return fibRec(n-1, &cache) + fibRec(n-2, &cache)
}

func fibRec(n int, cache *map[int]int64) int64 {
	if res, ok := (*cache)[n]; ok {
		return res
	}
	(*cache)[n] = fibRec(n-1, cache) + fibRec(n-2, cache)
	return (*cache)[n]
}

func main() {
	f, _ := os.Create("m.trace")
	trace.Start(f)
	defer trace.Stop()
	fib(1000000)
}
