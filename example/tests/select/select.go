package main

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecondTimeout = 10 * time.Second

// Racer は、2つのURLのレスポンスを比較して、より速いURLを返します。
// より早い =>をより適切な順番であることをtestしたいとする
func Racer(a, b string) (winner string, err error) {
	// aDuration := measureResponseTime(a)
	// bDuration := measureResponseTime(b)

	// if aDuration < bDuration {
	// 	return a
	// }
	// return b
	// selectでできることは、multiple チャネルで待機することです。 値を送信する最初のものは「勝ち」、caseの下のコードが実行されます。
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// NOTE: chan struct{}はメモリの観点から利用できる最小のデータ型のため、responseの型として定義しているがこだわりはない。
func ping(url string) chan struct{} {
	ch := make(chan struct{})
	// var ch chan struct{} の場合はnilで初期化されるためNG.
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
