package main

import (
	"fmt"
	"slices"
)

func main() {
	slice := []string{"Golang", "Java"}
	slice2 := []string{"Golang", "Java"}
	slice3 := []string{"Golang", "Javascript"}
	slice4 := []string{"Javascript", "Golang"}
	slice5 := []string{"Javascript", "Golang", "Golang"}
	slice6 := slice5[:1]

	isEq := slices.Equal(slice, slice2)
	isEq2 := slices.Equal[[]string](slice, slice3)
	isEq3 := slices.Equal[[]string](slice3, slice4)
	fmt.Println(isEq)
	fmt.Println(isEq2)
	fmt.Println(isEq3)
	// Output:
	// [1 2 3 4]

	i := slices.Index[[]string](slice, "Java")
	i2 := slices.Index[[]string](slice, "Javaa")
	fmt.Println(i)
	fmt.Println(i2)

	r := slices.Compact(slice5)
	r = slices.Insert[[]string](r, 0, "Hoge")
	r = append(r, "Hoge")
	fmt.Println(r)
	fmt.Println(slice6)

	p := newPerson("i", 2)
	g := p.intro("hloo")
	fmt.Println(g)
}

type Person struct {
	firstName string
	age       int
}

func newPerson(firstName string, age int) *Person {
	person := new(Person)
	person.firstName = firstName
	person.age = age
	return person
}

func (p *Person) intro(greetings string) string {
	return greetings + " I am " + p.firstName
}
