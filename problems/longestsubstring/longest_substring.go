package longestsubstring

import (
	"fmt"
	"strings"
)

/*
// Bytedance - round 2
# Given a string, print the longest substring without repeating characters.
# For example, the longest substrings without repeating characters for
#  a. “ABCDEFGABEF” are “BCDEFGA” and “CDEFGAB”, with length 7.
#  b.  “BBBB” the longest substring is “B”, with length 1.
#  c.  ABCABCD --> 4

map[string]int
A -> 1, B -> 1,

time complexity: O(N)
space complexity: O(N)

*/

/*
ABCDEFGABEF

ABCDEFG A. len = 7

BCDEFGA B, len = 7

CDEFGAB E,
*/

// time complexity: O(N)
// space complexity: O(N)
func FindLongestSubstring(str string) int {
	m := make(map[string]int)
	strArray := strings.Split(str, "")

	// create a map to count the number of occurences of each character: O(N)
	window := [2]int{0, 0}
	maxLength := -1
	for i, currChar := range strArray {
		window[1] = i
		// check the occurence
		if index, contains := m[currChar]; contains {
			window[0] = index
		}
		m[currChar] = i

		currLength := window[1] - window[0]
		if currLength > maxLength {
			fmt.Println(strArray[window[0]:window[1]])
			maxLength = currLength
		}
	}
	return maxLength
}
