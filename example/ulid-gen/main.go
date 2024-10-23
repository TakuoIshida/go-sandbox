package main

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

// Get ex. output: 0000XSNJG0MQJHBF4QX1EFD6Y3
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(ulid.Make())
	}
}
