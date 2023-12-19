package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// sample1()
	sample2()
}

func sample1() {
	start := time.Now()
	post := fetchPost()
	// channel の初期化
	// 2個のバッファを持った channel を作成

	resChan := make(chan any, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	go fetchPostLikes(post, resChan, &wg)
	go fetchPostComments(post, resChan, &wg)

	wg.Wait()
	// resChan channel への送信を終了し channel を閉じる
	close(resChan)

	// channel が閉じられるまでループする
	for res := range resChan {
		fmt.Println("res: ", res)
	}

	fmt.Println("took: ", time.Since(start))

}

// 投稿を一件取得する関数
func fetchPost() string {
	time.Sleep(time.Millisecond * 50)

	return "What programming languages do you prefer?"
}

// 投稿に紐づいたいいね数を取得する関数
func fetchPostLikes(post string, reschan chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 50)

	reschan <- 10
	wg.Done()
}

// 投稿に紐づいたコメントを全て取得する関数
func fetchPostComments(post string, reschan chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	reschan <- []string{"Golang", "Java", "Rust"}
	wg.Done()
}

func sample2() {
	ch := make(chan int, 2)

	var wg sync.WaitGroup
	wg.Add(2)

	// データを送信するgoroutine
	go func() {
		fmt.Println("Sending data...")
		time.Sleep(2 * time.Second) // saving
		ch <- 42
		ch <- 23
		wg.Done()
	}()

	// データを受信するgoroutine
	go func() {
		// ch からデータを取得するまでdead lockになる
		// 関数としては上の関数に依存してるので、可読性、メンテ性は良くない。
		value := <-ch
		value2 := <-ch
		fmt.Println("Received data:", value, value2)

		ch <- value + 1
		ch <- value2 + 1
		wg.Done()
	}()
	// 非同期処理が完了するまで待機する
	wg.Wait()
	// もう使用しないchを閉じる
	close(ch)

	// channel が閉じられるまでループする
	for res := range ch {
		fmt.Println("res: ", res)
	}

}
