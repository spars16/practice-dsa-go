package sorting

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

func TestSortPersonByAge(t *testing.T) {
	people := ByAge{
		Person{
			Name: "Sparsh",
			Age:  35,
		},
		Person{
			Name: "Cindy",
			Age:  23,
		},
		Person{
			Name: "Ananda",
			Age:  28,
		},
	}
	fmt.Println(people)
	sort.Sort(people)
	fmt.Println(people)
}

func TestSortPersonByName(t *testing.T) {
	people := []Person{
		{Name: "Sparsh", Age: 35},
		{Name: "Cindy", Age: 23},
		{Name: "Ananda", Age: 28},
	}
	fmt.Println(people)
	fmt.Println(ByName(people))
}

func TestInbuiltSort(t *testing.T) {
	// sort ints
	inputInts := []int{413, 5663, 215, 125, 634, 243}
	fmt.Println(inputInts)
	sort.Ints(inputInts)
	fmt.Println(inputInts)

	// sort floats
	inputFloat64s := []float64{62.214, 25.25, 6.3432, .5622, 251.63, 743.2}
	fmt.Println(inputFloat64s)
	sort.Float64s(inputFloat64s)
	fmt.Println(inputFloat64s)

	// sort strings
	inputStrings := []string{"lagsdl", "hdfs", "aswy", "ASKJA", "sgsdg", "SAGKS"}
	fmt.Println(inputStrings)
	sort.Strings(inputStrings)
	fmt.Println(inputStrings)

	// sort string
	word := "asgSLlOAWNvkiebc"
	inputString := strings.Split(word, "") // convert word into []string, with empty seperator
	fmt.Println(inputString)
	sort.Strings(inputString)
	fmt.Println(strings.Join(inputString, "")) // revert/"join" []string into string

	// sort slice
	s := []rune(word) // convert word into rune slice
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	fmt.Println(string(s))
}
