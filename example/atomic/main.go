package main

import (
	"fmt"
	"sync"
)

// go run -race main.go で競合チェックができる
func main() {
	// i := 0
	var i int64

	// Race
	// go func() {
	// 	i++
	// }()

	// go func() {
	// 	i++
	// }()

	// atomic
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go func() {
	// atomic.AddInt64(&i, 1)
	// 	wg.Done()
	// }()
	// go func() {
	// 	atomic.AddInt64(&i, 2)
	// 	wg.Done()
	// }()
	// wg.Wait()

	// mutex Lock
	mux := sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		mux.Lock()
		i++
		mux.Unlock()
		wg.Done()
	}()
	go func() {
		mux.Lock()
		i++
		mux.Unlock()
		wg.Done()
	}()
	wg.Wait()

	fmt.Println(i)
}
