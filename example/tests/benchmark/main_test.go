package main

import (
	"sync/atomic"
	"testing"
)

// benchmark の比較
// func BenchmarkXXX (b *testing.B)  で始める。出ないとbuild error
// 実行コマンド：go test -bench=. -count=10 | tee stats.txt | benchstat stats.txt
func BenchmarkAtomicStoreInt32(b *testing.B) {
	var i int32
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		atomic.StoreInt32(&i, 1)
	}
}

func BenchmarkAtomicStoreInt64(b *testing.B) {
	var i int64
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		atomic.StoreInt64(&i, 1)
	}
}
