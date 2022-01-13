package sort

import (
	"fmt"
	"sort"
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
