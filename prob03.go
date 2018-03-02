package main

import (
	"fmt"
	"math"
	"flag"
)

func main ( ) {

	var num, factor, maxfactor  int

	flag.IntVar(&num, "n", 600851475143, "number to factor")
	flag.Parse()

	maxfactor = int(math.Sqrt(float64(num)))

	factor = 2
	for num%factor == 0 {
		fmt.Print(factor," ")
		num = num/factor
	}

	for factor = 3; factor <= maxfactor ; factor = factor + 2 {
			if num%factor == 0 {
				fmt.Print(factor," ")
				num = num/factor
			}
	}
	fmt.Println()
}
