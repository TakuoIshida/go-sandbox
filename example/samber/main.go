package main

import (
	"fmt"
	"slices"
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

	currentUserBranchIDs := []int{1, 2, 3, 4, 5}
	operationUserIDs := []int{1, 2}
	holdBranchIDs := lo.Filter(currentUserBranchIDs, func(currentUserBranchID int, index int) bool {
		return !lo.Contains(operationUserIDs, currentUserBranchID)
	})

	fmt.Println(holdBranchIDs)

	// Concat
	concatenatedNames := slices.Concat([]string{"Alice", "Bob"}, []string{"Charlie", "David"})
	fmt.Println(concatenatedNames)

	BranchIDs := []int{1, 2, 3}
	allBranchIDs := []int{1, 2, 3, 4, 5}
	exists := lo.EveryBy(BranchIDs, func(branchID int) bool {
		return slices.Contains(allBranchIDs, branchID)
	})

	fmt.Println(exists)

}
