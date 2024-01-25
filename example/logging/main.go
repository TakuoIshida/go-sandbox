package main

import (
	"log/slog"
	"os"
)

func main() {
	// シンプルなLog出力
	slog.Info("Hello, World!")

	ops := slog.HandlerOptions{
		Level: slog.LevelInfo,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &ops))

	// 構造化ログ出力
	logger.Info("Hello, World!")

}
