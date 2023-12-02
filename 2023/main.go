package main

import (
	"advent-of-code-go/2023/day1"
	"fmt"
)

func main() {

	fmt.Println("Calculating solution for...")

	day, res, err := day1.Solve(false)

	fmt.Println("---> ", day, " <---")
	fmt.Println()

	if err != nil {
		fmt.Printf("Something went wrong: %v", err)
	}
	fmt.Println(res)
}
