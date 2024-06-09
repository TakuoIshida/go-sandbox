package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 引数にメソッドを持つインターフェース
type Callback[T any, U any] interface {
	Do(param T) (U, error)
}

func main() {
	fmt.Println("Hello, Exponential Backoff!")

	// Exponential Backoff
	// 指数関数的に待ち時間を設定した指数バックオフ再実行関数（ランダム要素なし）

	// 引数にメソッドを渡して実行
	cb := PrinterCallback{}
	res, err := retryWithExponentialBackoff(cb, 3)
	if err != nil {
		fmt.Println("エラー:", err)
	} else {
		fmt.Println("成功!")
		fmt.Println(res)
	}
}

// https://zenn.dev/nobishii/articles/type_param_intro
func retryWithExponentialBackoff[T any, U any](fn Callback[T, U], param T) (U, error) {
	// デフォルト値
	maxRetries := 5             // 最大再試行回数
	initialIntervalSeconds := 1 // 初期待機時間(s)
	maxIntervalSeconds := 16    // 最大待機時間(s) [optional]
	var res U
	var err error
	for i := 0; i < maxRetries; i++ {
		res, err = fn.Do(param)
		if err == nil {
			return res, nil
		}
		// 指数関数的に待機時間を計算
		waitSeconds := int(math.Pow(float64(2.0), float64(i)) * float64(initialIntervalSeconds))
		// 最大待機時間と比較して制限[optional]
		if waitSeconds > maxIntervalSeconds {
			waitSeconds = maxIntervalSeconds
		}
		time.Sleep(time.Duration(waitSeconds) * time.Second)
		fmt.Printf("リトライ #%d回目 失敗: %v, 待機時間: %vs\n", i+1, err, waitSeconds)
	}
	return res, fmt.Errorf("最大回数 (%d) まで再試行しましたが、失敗しました", maxRetries)
}

type PrinterCallback struct{}

func (p PrinterCallback) Do(i int) (string, error) {
	s := fmt.Sprintf("hello: %d回目", i)
	randomInt := rand.Intn(100)
	if randomInt%2 == 0 {
		return "", fmt.Errorf("エラー: %d回目", i)
	}
	return s, nil
}
