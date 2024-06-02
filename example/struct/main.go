package main

import (
	"fmt"
	"reflect"
	"slices"

	"github.com/google/go-cmp/cmp"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	names := []string{"Alice", "Bob", "Alice", "Charlie", "Bob"}
	names2 := []string{"Alice", "Bob", "Alice", "Charlie", "Bob"}
	// fmt.Println(names)

	s1 := Person{
		Name: "Alice",
		Age:  25,
	}
	s2 := Person{
		Name: "Alice",
		Age:  25,
	}

	slc1 := []int{0, 1, 2, 3, 4}
	slc2 := []int{0, 1, 2, 3, 4}
	if slices.Equal(slc1, slc2) {
		fmt.Println("slc1 == slc2: true")
	} else {
		fmt.Println("slc1 == slc2: false")
	}

	isSliceDeepEqual(names, names2)
	isStructDeepEqual(s1, s2)

	cmpStruct()
}

func isSliceDeepEqual(slice []string, slice2 []string) {
	result := reflect.DeepEqual(slice, slice2)
	fmt.Println(result)
}
func isStructDeepEqual(person Person, person2 Person) {
	result := reflect.DeepEqual(person, person2)
	fmt.Println(result)
}

func cmpStruct() {
	type PersonWithFriends struct {
		Person
		Friends []string
	}
	person1 := PersonWithFriends{
		Person:  Person{"Alice", 30},
		Friends: []string{"Bob", "Charlie"},
	}
	person2 := PersonWithFriends{
		Person:  Person{"Alice", 30},
		Friends: []string{"Bob", "Charlie"},
	}

	if cmp.Equal(person1, person2) {
		fmt.Println("Persons are equal")
	} else {
		fmt.Println("Persons are not equal")
		fmt.Println(cmp.Diff(person1, person2))
	}
}
