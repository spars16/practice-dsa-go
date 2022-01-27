package sorting

import "sort"

type (
	Person struct {
		Name   string
		Age    int
		Gender string
	}

	ByAge []Person
)

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func ByName(people []Person) []Person {
	byName := make([]Person, len(people))
	copy(byName, people)

	sort.Slice(byName, func(i, j int) bool {
		return byName[i].Name < byName[j].Name
	})

	return byName
}
