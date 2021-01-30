package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	fmt.Println("Using function with single return...")

	sum := sum(7, 6)

	fmt.Println(" sum 7 + 6 = ", sum)

	fmt.Println("Using function with multiple return...")

	result, err := sqrt(16)

	if err != nil {
		fmt.Println("sqrt of 16 = ", result)
	} else {
		fmt.Println("Error: ", err)
	}

	result1, err1 := sqrt(0)

	if err1 != nil {
		fmt.Println("sqrt of 0 = ", result1)
	} else {
		fmt.Println("Error: ", err1)
	}

}

func sum(x int, y int) int {
	return x + y
}

//functions in go can return multiple values
//like nuples in scala?
//so in go can i return functions???

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	} else {
		return math.Sqrt(x), nil
	}
}
