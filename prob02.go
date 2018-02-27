package main

import "fmt"

func main ( ) {

	var (
		last int = 1
		curr int = 2
		next int 
		sum int = 2
	)

	for curr < 4000000 {
		next = curr + last
		last = curr
		curr = next
		if curr%2 == 0 {
			sum = sum + curr
		}
	}
	fmt.Println(sum)
}
