package slidematrixgame

/*# Write a function that determines whether or not it is possible to make it from the starting coordinates to the ending coordinates given the following:
# - You are given a "map" of 0's and 1's. 0's are empty space and can be moved through. 1's are walls and you cannot move through them
# - You can move up, down, left, or right
# - If you move in any directions you must "slide" in that direction until you either reach a coordinate right before a wall (you hit the wall) or the edge of the map
# - Youre given start and end location are 0s
#
# EX start (0,0), end (3, 3) (marked with s and e just to visualize)
#    s 0 0 1
#    0 0 0 1
#    0 0 1 0
#    1 0 0 e
*/

const (
	UNVISITED = iota // 0
	VISITED          // 1
)

const (
	LEFT  = iota // 0
	RIGHT        // 1
	UP
	DOWN
	NUM_DIRECTIONS
)

type Coordinate struct {
	value   int
	x       int
	y       int
	visited int // 0
}

// Coordinate: x int, y int, visited int
// logic:
// - at start, dfs check directions until you reach
// a wall, a visited coordinate, or end
// - if wall, recursively check directions again
// - if visited, return false (backtrack)
// - if end, return true

// Time complexity: O(N)
// Space complexity: O(1)

func FindPath(matrix [][]*Coordinate, start, end *Coordinate) bool {
	return checkAllDirections(matrix, start, end) != nil
}

func checkAllDirections(matrix [][]*Coordinate, start, end *Coordinate) *Coordinate {
	for i := 0; i < NUM_DIRECTIONS; i++ {
		c := checkAllDirections(matrix, start, end)
		if c.x == end.x && c.y == end.y {
			return c
		}
	}
	return nil
}

func checkDirection(matrix [][]*Coordinate, point *Coordinate, direction int) *Coordinate {
	// go in a direction until you hit a wall
	switch direction {
	case UP:
		for i := point.x; i > 0; i-- {
			currPoint := matrix[i][point.y]
			nextPoint := matrix[i-1][point.y]
			if nextPoint.value == 1 {
				if currPoint.visited == VISITED {
					return nil
				} else {
					return currPoint
				}
			}
		}
	case DOWN:
		for i := point.x; i < len(matrix)-1; i++ {
			currPoint := matrix[i][point.y]
			nextPoint := matrix[i+1][point.y]
			if nextPoint.value == 1 {
				if currPoint.visited == VISITED {
					return nil
				} else {
					return currPoint
				}
			}
		}
	case LEFT:
		for i := point.y; i > 0; i-- {
			currPoint := matrix[point.x][i]
			nextPoint := matrix[point.x][i-1]
			if nextPoint.value == 1 {
				if currPoint.visited == VISITED {
					return nil
				} else {
					return currPoint
				}
			}
		}
	case RIGHT:
		for i := point.y; i < len(matrix[point.x])-1; i++ {
			currPoint := matrix[point.x][i]
			nextPoint := matrix[point.x][i+1]
			if nextPoint.value == 1 {
				if currPoint.visited == VISITED {
					return nil
				} else {
					return currPoint
				}
			}
		}
	default:
		// do something
	}
	return nil
}
