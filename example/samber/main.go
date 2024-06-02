package main

import (
	"fmt"
	"strconv"

	"github.com/samber/lo"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	println("Hello, Samber!")
	// Unique
	names := lo.Uniq([]string{"Alice", "Bob", "Alice", "Charlie", "Bob"})

	// Filter
	filteredNames := lo.Filter(names, func(item string, index int) bool {
		return item != "Alice"
	})
	fmt.Println(filteredNames) // [Bob Charlie]

	// Map
	mappedNames := lo.Map(filteredNames, func(item string, index int) string {
		return item + strconv.Itoa(index)
	})
	fmt.Println(mappedNames) // [Bob0 Charlie1]

	// Flatten
	flattenNames := lo.Flatten([][]string{{"Alice", "Bob", "Alice", "Charlie", "Bob"}})
	fmt.Println(flattenNames) // [Alice Bob Charlie Bob]

	// Chunk
	chunkedNames := lo.Chunk(flattenNames, 2)
	fmt.Println(chunkedNames) // [[Alice Bob] [Alice Charlie Bob]]

	// Compact
	compactedNames := lo.Compact([]string{"Alice", "", "", "Bob"})
	fmt.Println(compactedNames) // [Alice Bob]

	// IsNil
	var nilSlice []string
	fmt.Println(lo.IsNil(nilSlice)) // true

	// ToPtr
	ptr := lo.ToPtr("Alice")
	fmt.Println(ptr) // 0xc0000b8000
}
