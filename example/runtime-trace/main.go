package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"sync"
	"time"

	"golang.org/x/exp/rand"
)

// 並行処理で役立つデバッグ&分析手法
// https://zenn.dev/hsaki/books/golang-concurrency/viewer/analysis#%E6%94%B9%E5%96%84%E5%89%8D%E3%81%AE%E5%87%A6%E7%90%86%E3%82%92trace%E3%81%A7%E3%81%8D%E3%82%8B%E3%82%88%E3%81%86%E3%81%AB%E3%81%99%E3%82%8B
func RandomWait(ctx context.Context, i int) {
	// regionを始める
	defer trace.StartRegion(ctx, "randomWait").End()
	fmt.Printf("No.%d start\n", i+1)
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	fmt.Printf("No.%d done\n", i+1)
}

func _main() {
	// タスクを定義
	ctx, task := trace.NewTask(context.Background(), "main")
	defer task.End()

	rand.Seed(uint64(time.Now().UnixNano()))
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			RandomWait(ctx, i)
		}(i)
	}
	wg.Wait()
}

func main() {
	// トレースを始める
	// 結果出力用のファイルもここで作成
	f, err := os.Create("tseq.out")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalln("Error:", err)
	}
	defer trace.Stop()

	_main()
}
