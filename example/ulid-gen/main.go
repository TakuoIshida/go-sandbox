package main

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

// Get ex. output: 0000XSNJG0MQJHBF4QX1EFD6Y3
func main() {
	fmt.Println(ulid.Make())
}
