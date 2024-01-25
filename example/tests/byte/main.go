package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	c := bytes.NewBuffer([]byte("Hello "))
	fmt.Print(c.Len())
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "World!")
	b.Truncate(4)
	if _, err := b.WriteTo(os.Stdout); err != nil {
		fmt.Println(err)
	}
}
