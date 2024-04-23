package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func doSomething(i int) error {
	if i%2 == 0 {
		return fmt.Errorf("error: %d", i)
	}
	fmt.Println("do something")
	return nil
}

func main() {
	eg, _ := errgroup.WithContext(context.Background())

	var errs []error
	for i := 0; i < 10; i++ {
		eg.Go(func() error {
			// return doSomething(i)
			err := doSomething(i)
			if err != nil {
				errs = append(errs, err)
			}
			return nil
		})
	}

	// ↓がないと、エラーが発生してもエラーが表示されない
	eg.Wait() //  すべてのgoroutineが終了するまで待つ.

	// 最初のerrorを返す or 全てのgoroutineが成功すればいい場合は、こちらを使う
	// err := eg.Wait()

	for _, err := range errs {
		fmt.Println(err)
	}
}
