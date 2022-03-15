package main

import (
	"fmt"
	"time"
)

func main() {
	var numOfCases int
	fmt.Scanln(&numOfCases)
	results := make([]int, numOfCases)

	start := time.Now()

	readCase(&results, 0, numOfCases)

	elapsed := time.Since(start)
	fmt.Printf("timing: %s", elapsed*1000)
}

func readCase(results *[]int, i int, n int) {
	if i == n {
		print(results, 0, n)
		return
	}
	var size int
	if m, err := fmt.Scan(&size); m != 1 {
		panic(err)
	}

	arr := make([]int, size)
	// var ptr *int = (&results)[i]
	// readArray(&arr, ptr, 0, size)
	(*results)[i] = readArray(&arr, 0, size)
	readCase(results, i+1, n)
}

func readArray(a *[]int, i int, n int) int {
	if i == n {
		return 0
	}
	var num int
	if m, err := fmt.Scan(&num); m != 1 {
		panic(err)
	}

	if num >= 0 {
		return num*num + readArray(a, i+1, n)
	}
	return readArray(a, i+1, n)
}

func print(a *[]int, i int, n int) {
	if i == n {
		return
	}
	fmt.Println((*a)[i])
	print(a, i+1, n)
}
