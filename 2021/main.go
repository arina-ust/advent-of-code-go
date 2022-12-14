package main

import (
	"advent-of-code-go/2021/day16"
	"fmt"
)

func main() {

	fmt.Println("Calculating solution for...")

	day, res, err := day16.Solve(false)

	fmt.Println("---> ", day, " <---")
	fmt.Println()

	if err != nil {
		fmt.Printf("Something went wrong: %v", err)
	}
	fmt.Println(res)
}
