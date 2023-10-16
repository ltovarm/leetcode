package main

import (
	"fmt"
	"math/big"
)

func getFactorial(n int) *big.Int {

	out := new(big.Int)
	out.SetInt64(1)

	if n <= 1 {
		return out
	}

	// Calculate factorial
	for i := 2; i < n+1; i++ {
		out.Mul(out, big.NewInt(int64(i)))
	}

	return out
}

func getRow(rowIndex int) []int {
	var out []int

	for i := 0; i < rowIndex+1; i++ {
		// C = n!/(r!*(n-r)!)
		c := new(big.Int)
		c.SetInt64(1)
		if i != 0 {
			c.Mul(getFactorial(i), getFactorial(rowIndex-i))
			c.Div(getFactorial(rowIndex), c)
		}
		c_int := c.Int64()
		out = append(out, int(c_int))
	}
	return out
}

func main() {

	var sol []int
	sol = getRow(21)
	for _, value := range sol {
		fmt.Print(value)
		fmt.Print("\t")
	}
	fmt.Print("\n")
}
