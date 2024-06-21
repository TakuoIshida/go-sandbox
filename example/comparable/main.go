package main

import "fmt"

func main() {
	// s := []int{1, 2}
	// s2 := []int{1, 3}
	// // slice can only be compared to nil
	// fmt.Println(s == s2)

	// var aは、*int32型の変数
	var a *int32 = nil
	// var bは、*int64型の変数
	// var b *int64 = nil
	var b *int32 = nil
	// aとbは、異なる型のnilなので、等しくないと判断される
	fmt.Println(a == b) // false

	// var x [2]func()
	// var y struct{ s []int }     // User.ApplicationAccess []string
	// fmt.Println(x == x, y == y) // NG

	// ## 比較できる
	type Single struct {
		Name string
		Age  int
	}
	type Parent struct {
		Name        string
		Age         int
		ChildrenIDs []int
	}
	single := Single{
		Name: "Alice",
		Age:  25,
	}
	single2 := Single{
		Name: "Alice",
		Age:  25,
	}

	// ## 比較できない
	parent := Parent{
		Name:        "Alice",
		Age:         25,
		ChildrenIDs: []int{1, 2, 3},
	}

	fmt.Println(parent)            // NG
	fmt.Println(single == single2) // OK

}
