package main

import (
	"fmt"
	"sync"
	"time"
)

// go run -race main.go で競合チェックができる
func main() {
	runRaceCondition()
	runMutexLock()
}

// データ競合
func runRaceCondition() {
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	// srcの要素毎にある何か処理をして、結果をdstにいれる
	for _, s := range src {
		go func(s int) {
			// 何か(重い)処理をする
			result := s * 2

			// 結果をdstにいれる
			dst = append(dst, result)
		}(s)
	}
	time.Sleep(time.Second)
	fmt.Println("runRaceCondition", dst)
}

func runMutexLock() {
	mux := sync.Mutex{}
	var wg sync.WaitGroup

	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	// srcの要素毎にある何か処理をして、結果をdstにいれる
	wg.Add(len(src))
	for _, s := range src {
		go func(s int) {
			// 何か(重い)処理をする
			result := s * 2

			// 結果をdstにいれる
			mux.Lock()
			dst = append(dst, result)
			mux.Unlock()
			wg.Done()
		}(s)
	}
	wg.Wait()
	fmt.Println("runMutexLock", dst) // 必ず、srcの要素数と同じ数の要素が入っている
}
