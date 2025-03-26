package day6

import (
	"advent-of-code-go/util"
	"fmt"
)

const day = "day6"

var inputFile string

func Solve(easy bool) (name string, res int, err error) {
	name = day
	setInput(easy)

	// Read grid directly
	grid, err := util.ReadRuneGrid(inputFile)
	if err != nil {
		return
	}
	if len(grid) == 0 {
		return name, 0, nil
	}

	// Solve part one
	res, err = partOne(grid)

	// Uncomment and implement part two
	// res, err = partTwo(grid)

	return
}

func setInput(easy bool) {
	if easy {
		inputFile = day + "/" + util.InputFileEasy
	} else {
		inputFile = day + "/" + util.InputFileFull
	}
}

// Position represents a unique grid position
type Position struct {
	X, Y int
}

func partOne(grid [][]rune) (int, error) {
	// Track unique visited positions
	visited := make(map[Position]bool)

	// Find starting position ('^' character)
outer:
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '^' {
				// Start exploring from this position
				explorePositions(grid, x, y, visited)
				break outer
			}
		}
	}

	return len(visited), nil
}

func explorePositions(grid [][]rune, startX, startY int, visited map[Position]bool) {
	// Define right turns for each direction
	rightTurns := map[[2]int][2]int{
		{0, -1}: {1, 0},  // up -> right
		{1, 0}:  {0, 1},  // right -> down
		{0, 1}:  {-1, 0}, // down -> left
		{-1, 0}: {0, -1}, // left -> up
	}

	// Queue for BFS
	type queueItem struct {
		pos Position
		dir [2]int
	}
	queue := []queueItem{{Position{startX, startY}, [2]int{0, -1}}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Print the grid with current position marked
		// printGridWithPosition(grid, current.pos.X, current.pos.Y)

		// Mark current position as visited
		visited[current.pos] = true

		// Try moving in current direction
		nextX := current.pos.X + current.dir[0]
		nextY := current.pos.Y + current.dir[1]

		// If next position is off the grid, continue to next queue item
		if nextX < 0 || nextX >= len(grid[0]) ||
			nextY < 0 || nextY >= len(grid) {
			continue
		}

		// If not an obstacle, add to queue
		if grid[nextY][nextX] != '#' {
			newItem := queueItem{
				pos: Position{nextX, nextY},
				dir: current.dir,
			}
			queue = append(queue, newItem)
			continue
		}

		// If blocked, turn right
		rightTurn := rightTurns[current.dir]
		nextX = current.pos.X + rightTurn[0]
		nextY = current.pos.Y + rightTurn[1]

		// If next position is off the grid, stop exploring
		if nextX < 0 || nextX >= len(grid[0]) ||
			nextY < 0 || nextY >= len(grid) {
			continue
		}

		// If not an obstacle, add to queue
		if grid[nextY][nextX] != '#' {
			queue = append(queue, queueItem{
				pos: Position{nextX, nextY},
				dir: rightTurn,
			})
		}
	}
}

// printGridWithPosition prints the grid with the current position marked with '^'
func printGridWithPosition(grid [][]rune, x, y int) {
	// Create a copy of the grid to avoid modifying the original
	printGrid := make([][]rune, len(grid))
	for i := range grid {
		printGrid[i] = make([]rune, len(grid[i]))
		copy(printGrid[i], grid[i])
	}

	// Mark the current position with '^'
	printGrid[y][x] = '^'

	// Print the grid
	for _, row := range printGrid {
		fmt.Println(string(row))
	}
	fmt.Println() // Add a blank line for readability
}

func partTwo(grid [][]rune) (int, error) {
	// TODO: Implement part two solution
	return 0, nil
}
