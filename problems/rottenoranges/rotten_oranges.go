package rottenoranges

/*
Argo AI - 1st round
Duo - 1st round
*/

import "fmt"

const (
	freshOrangeInt  = 1
	rottenOrangeInt = 2
)

type orange struct {
	X int
	Y int
}

// Given a grid of oranges (fresh, rotten, or none)
// A fresh orange becomes rotten if it is neighboring a rotten orange, in the four grid directions
// RETURN the number of "minutes" before all the fresh oranges are rotten
// otherwise return -1 if there will always be some number of fresh oranges remaining in the grid
func RottenOranges(grid [][]int, numIterations int, prevNumFreshOranges int) int {
	rottenOranges, numFreshOranges := getOranges(grid)
	if numFreshOranges == 0 {
		fmt.Printf("No more fresh oranges remaining after %d minutes\n", numIterations)
		return numIterations
	}

	if prevNumFreshOranges == numFreshOranges {
		fmt.Printf("No solution found, %d fresh orange(s) remaining.\n", numFreshOranges)
		return -1
	}

	updatedRottenOranges := make([]orange, 0)
	for _, rottenOrange := range rottenOranges {
		freshNeighbors := getFreshNeighbors(grid, rottenOrange)
		updatedRottenOranges = append(updatedRottenOranges, freshNeighbors...)
	}
	grid = updateGrid(grid, updatedRottenOranges)
	return RottenOranges(grid, numIterations+1, numFreshOranges)
}

// RETURNS a slice of the locations of all of the rotten oranges,
// and the number of fresh oranges remaining in the grid.
func getOranges(grid [][]int) (rottenOrangesArray []orange, numFreshOranges int) {
	for i := 0; i < len(grid); i++ {
		for j := range grid[i] {
			if grid[i][j] == rottenOrangeInt {
				rottenOrangesArray = append(rottenOrangesArray, orange{X: i, Y: j})
			} else if grid[i][j] == freshOrangeInt {
				numFreshOranges++
			}
		}
	}
	return rottenOrangesArray, numFreshOranges
}

// Checks the four directions for a given rotten orange location,
// and RETURNS an updated slice of all of the "new" rotten oranges -
//  (the fresh oranges that became rotten from neighboring rotten oranges)
func getFreshNeighbors(grid [][]int, rottenOrange orange) []orange {
	freshNeighbors := make([]orange, 0)
	x := rottenOrange.X
	y := rottenOrange.Y
	if x < len(grid)-1 && grid[x+1][y] == freshOrangeInt {
		// TODO: check for duplicates before appending to slice
		freshNeighbors = append(freshNeighbors, orange{X: x + 1, Y: y})
	} else if y < len(grid)-1 && grid[x][y+1] == freshOrangeInt {
		freshNeighbors = append(freshNeighbors, orange{X: x, Y: y + 1})
	} else if x > 0 && grid[x-1][y] == freshOrangeInt {
		freshNeighbors = append(freshNeighbors, orange{X: x - 1, Y: y})
	} else if y > 0 && grid[x][y-1] == freshOrangeInt {
		freshNeighbors = append(freshNeighbors, orange{X: x, Y: y - 1})
	}
	return freshNeighbors
}

// Updates the grid to include the new rotten oranges,
// and RETURNS the updated grid
func updateGrid(grid [][]int, rottenOranges []orange) [][]int {
	for _, rottenOrange := range rottenOranges {
		grid[rottenOrange.X][rottenOrange.Y] = rottenOrangeInt
	}
	return grid
}
