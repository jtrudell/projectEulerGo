// projecteuler.net Problem 2: Even Fibonacci numbers
// By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.

package main

import "fmt"

func evenFib(previous, last, sum int) int {

	if last >= 4000000 {
		return sum
	}

	nextFib := previous + last

	if nextFib%2 == 0 {
		sum += nextFib
	}

	return evenFib(last, nextFib, sum)
}

func main() {
	fmt.Println(evenFib(1, 2, 2))
}
