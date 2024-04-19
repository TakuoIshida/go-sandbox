package main

import (
	"context"
	"io"
	"net"
	"net/http"
	"testing"
)

func buildServer(body string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, body)
	})

	return &http.Server{
		Handler: mux,
	}
}

func Test_main(t *testing.T) {
	t.Helper()
	want := "Hello, World!"
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, want)
	})
	// Arrange
	srv := buildServer(want)

	// 動的のportを取得
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		if err := srv.Serve(l); err != http.ErrServerClosed {
			t.Errorf("Server error: %v", err)
		}
		// サーバーが終了したことを通知
		close(idleConnsClosed)
	}()

	// Act
	res, err := http.Get("http://" + l.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	res.Body.Close()

	// Assert
	if string(b) != want {
		t.Fatalf("want %q, but %q", want, b)
	}

	// Cleanup
	if err := srv.Shutdown(context.Background()); err != nil {
		t.Fatalf("HTTP server Shutdown: %v", err)
	}

	// サーバの終了を確認してからテストコードを終了する。
	<-idleConnsClosed

}
