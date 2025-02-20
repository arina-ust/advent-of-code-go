package day5

import (
	"advent-of-code-go/util"
	"fmt"
	"strconv"
	"strings"
)

const day = "day5"

var inputFile string

func Solve(easy bool) (name string, res int, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadStringList(inputFile)
	if err != nil {
		return
	}
	if len(lines) == 0 {
		return name, 0, nil
	}

	// Uncomment and implement part one
	// res, err = partOne(lines)

	// Uncomment and implement part two
	res, err = partTwo(lines)

	return
}

func setInput(easy bool) {
	if easy {
		inputFile = day + "/" + util.InputFileEasy
	} else {
		inputFile = day + "/" + util.InputFileFull
	}
}

func partOne(lines []string) (int, error) {
	// Parse the input
	forwardMap, reverseMap, startIndex := parseInput(lines)

	// Debug print the maps
	// fmt.Println("Forward Map:")
	// for k, v := range forwardMap {
	// 	fmt.Printf("%d: %v\n", k, v)
	// }
	// fmt.Println("\nReverse Map:")
	// for k, v := range reverseMap {
	// 	fmt.Printf("%d: %v\n", k, v)
	// }

	// Track valid line indices
	validLineIndices := make([]int, 0)

	// Process comma-separated lines
	for i, line := range lines[startIndex:] {
		if line == "" {
			continue
		}

		numbers := strings.Split(line, ",")
		if len(numbers) < 2 {
			continue
		}

		// Convert strings to integers
		nums := make([]int, len(numbers))
		for j, numStr := range numbers {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return 0, err
			}
			nums[j] = num
		}

		// Check if the line is valid
		if isValidLine(nums, forwardMap, reverseMap) {
			validLineIndices = append(validLineIndices, i+startIndex)
		}
	}

	// Print valid line indices
	// fmt.Println("Valid Line Indices:")
	// for _, idx := range validLineIndices {
	// 	fmt.Printf("%d: %s\n", idx, lines[idx])
	// }

	// Sum middle numbers of valid lines
	middleSum := 0
	for _, idx := range validLineIndices {
		numbers := strings.Split(lines[idx], ",")
		if len(numbers) > 1 {
			middleNum, _ := strconv.Atoi(numbers[len(numbers)/2])
			middleSum += middleNum
		}
	}

	return middleSum, nil
}

func parseInput(lines []string) (map[int][]int, map[int][]int, int) {
	forwardMap := make(map[int][]int)
	reverseMap := make(map[int][]int)

	// Find the boundary between the first and second sections
	var boundaryIndex int
	for i, line := range lines {
		if strings.Contains(line, ",") {
			boundaryIndex = i - 1
			break
		}
	}

	// Parse the first section (pairs)
	for _, line := range lines[:boundaryIndex] {
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			continue
		}

		first, err1 := strconv.Atoi(parts[0])
		second, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			continue
		}

		// Add to forward map (first -> second)
		forwardMap[first] = append(forwardMap[first], second)

		// Add to reverse map (second -> first)
		reverseMap[second] = append(reverseMap[second], first)
	}

	return forwardMap, reverseMap, boundaryIndex + 1
}

func partTwo(lines []string) (int, error) {
	// Parse the input
	forwardMap, reverseMap, startIndex := parseInput(lines)

	// Track reordered lines
	reorderedLines := make(map[int][]int)

	// Process comma-separated lines
	for i, line := range lines[startIndex:] {
		if line == "" {
			continue
		}

		numbers := strings.Split(line, ",")
		if len(numbers) < 2 {
			continue
		}

		// Convert strings to integers
		nums := make([]int, len(numbers))
		for j, numStr := range numbers {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return 0, err
			}
			nums[j] = num
		}

		// Check if the line is invalid
		if !isValidLine(nums, forwardMap, reverseMap) {
			// fmt.Printf("Found invalid line %d: %v\n", i+startIndex, nums)
			// Try to reorder the line
			reorderedNums := reorderLine(nums, forwardMap, reverseMap)

			// if len(reorderedNums) > 0 {
			// 	fmt.Printf("Reordered line %d: %v\n", i+startIndex, reorderedNums)
			// } else {
			// 	fmt.Printf("Failed to reorder line %d\n", i+startIndex)
			// }

			// If reordering is successful and the new line is valid
			if len(reorderedNums) > 0 && isValidLine(reorderedNums, forwardMap, reverseMap) {
				reorderedLines[i+startIndex] = reorderedNums
				// fmt.Printf("Successfully reordered and validated line %d\n", i+startIndex)
			}
		}
	}

	// Sum middle numbers of reordered lines
	middleSum := 0
	for _, nums := range reorderedLines {
		middleNum := nums[len(nums)/2]
		middleSum += middleNum
	}

	fmt.Printf("Found %d reordered lines\n", len(reorderedLines))
	return middleSum, nil
}

// modifies nums in place
func reorderLine(nums []int, forwardMap, reverseMap map[int][]int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	for !isValidLine(nums, forwardMap, reverseMap) {
		// Try each number as a potential starting point
	outer:
		for start := 0; start < n-1; start++ {
			num1 := nums[start]
			num2 := nums[start+1]

			reverseMappings, hasReverseMapping := reverseMap[num2]
			if hasReverseMapping {
				for _, mapping := range reverseMappings {
					if mapping == num1 {
						continue outer
					}
				}

				// Swap the numbers
				nums[start], nums[start+1] = num2, num1

				// Check if the new line is valid
				if isValidLine(nums, forwardMap, reverseMap) {
					return nums
				}
			}

			forwardMappings, hasForwardMapping := forwardMap[num1]
			if hasForwardMapping {
				for _, mapping := range forwardMappings {
					if mapping == num2 {
						continue outer
					}
				}

				// Swap the numbers
				nums[start], nums[start+1] = num2, num1

				// Check if the new line is valid
				if isValidLine(nums, forwardMap, reverseMap) {
					return nums
				}
			}
		}
	}

	return nums
}

func isValidLine(nums []int, forwardMap, reverseMap map[int][]int) bool {
	// Forward map check: iterate from left to right
	for i := 0; i < len(nums); i++ {
		currentNum := nums[i]

		// Check forward mappings for the current number
		forwardMappings, hasForwardMapping := forwardMap[currentNum]
		if hasForwardMapping {
			// Look at numbers to the right of the current number
			for j := i + 1; j < len(nums); j++ {
				nextNum := nums[j]

				// Check if nextNum is in the forward mappings
				found := false
				for _, mapping := range forwardMappings {
					if mapping == nextNum {
						found = true
					}
				}

				// If nextNum is not in forward mappings, line is invalid
				if !found {
					return false
				}
			}
		}
	}

	// Reverse map check: iterate from right to left
	for i := len(nums) - 1; i >= 0; i-- {
		currentNum := nums[i]

		// Check reverse mappings for the current number
		reverseMappings, hasReverseMapping := reverseMap[currentNum]
		if hasReverseMapping {
			// Look at numbers to the left of the current number
			for j := i - 1; j >= 0; j-- {
				prevNum := nums[j]

				// Check if prevNum is in the reverse mappings
				found := false
				for _, mapping := range reverseMappings {
					if mapping == prevNum {
						found = true
					}
				}

				// If prevNum is not in reverse mappings, line is invalid
				if !found {
					return false
				}
			}
		}
	}

	return true
}
