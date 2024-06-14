package main

import "fmt"

func main() {
	emails := []string{"test@example.com", "test2@example.com"}
	emails2 := []string{"test2@example.com", "test3@example.com"}
	for _, email := range emails {
		for _, email2 := range emails2 {
			if email == email2 {
				fmt.Println("Found!")
				break
			}
		}
	}
}
