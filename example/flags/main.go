package main

import (
	"flag"
	"fmt"
	"strings"
)

type StringSlice []string

// String implements the flag.Value interface
func (s *StringSlice) String() string {
	return fmt.Sprint(*s)
}

// Set implements the flag.Value interface
func (s *StringSlice) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

// go run main.go -slices=hogehoge
func main() {
	var sliceFlag StringSlice

	flag.Var(&sliceFlag, "slices", "Comma-separated list of strings")

	flag.Parse()

	fmt.Println("Parsed slices:", sliceFlag)
}
