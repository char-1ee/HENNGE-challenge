/*
 Author: Li Xingjian
 Email: xingjianli59@gmail.com

 Compute the squares sums of given integers.
 Noticed that no for loop is allowed, we use recursive calls to replace use of loops.
 To improve performance of recursion:
	1. use tail recursions
	2. pass by pointers rather than values
	3. avoid global variables
*/
package main

import "fmt"

func main() {
	var numOfCases int // N: number of test cases to follow
	fmt.Scanln(&numOfCases)
	results := make([]int, numOfCases) // array to record sum of squares from each test case
	readCase(&results, 0, numOfCases)
}

// readCase() takes pointer to `result`, current index i, recursion count n as parameters.
// Before index i reach n, it reads from stdin the number of space-seperated integers and call readArray()
// in every recursion. Once index i reach n, readCase() call print() to print the sums in `results`.
func readCase(results *[]int, i int, n int) {
	if i == n {
		print(results, 0, n)
		return
	}

	var size int // X: number of space-seperated integers
	if m, err := fmt.Scan(&size); m != 1 {
		panic(err)
	}

	arr := make([]int, size)
	(*results)[i] = readArray(&arr, 0, size)
	readCase(results, i+1, n)
}

// readArray() takes pointer to test case array, current index i, recursion count n as parameters,
// and return the sum of squares of positive integers.
// Before index reach n, it reads an integer from Scan() and sum up recursively. When i equals n, return 0.
func readArray(a *[]int, i int, n int) int {
	if i == n {
		return 0
	}
	var num int // Yn: input integer of every test case
	if m, err := fmt.Scan(&num); m != 1 {
		panic(err)
	}

	if num >= 0 {
		return num*num + readArray(a, i+1, n)
	}
	return readArray(a, i+1, n)
}

// print() takes pointer to result array, current index i, recursion count n as parameters.
// Before recursing up to n times, print() prints the sums from result array in recursion to stdout.
func print(a *[]int, i int, n int) {
	if i == n {
		return
	}
	fmt.Println((*a)[i])
	print(a, i+1, n)
}
