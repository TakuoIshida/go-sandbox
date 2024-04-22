package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	type customer struct {
		id         string
		operations []float64
	}

	cust1 := customer{id: "1", operations: []float64{1.0, 2.0}}
	cust2 := customer{id: "1", operations: []float64{1.0, 2.0}}
	// fmt.Println((cust1 == cust2))
	fmt.Println(reflect.DeepEqual(cust1, cust2)) // == で独自に比較する関数よりも100倍遅い => testの場合は、testify/assert.EqualValuesを使う
	// byteなら、
	r := bytes.Compare([]byte{1, 2}, []byte{1, 2})
	fmt.Println(r == 0)
}
