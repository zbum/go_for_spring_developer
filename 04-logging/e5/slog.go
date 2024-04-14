package main

import (
	"context"
	"log/slog"
	"os"
)

func main() {
	// Top Level Function (with Default)
	slog.Info("hello world", "user", os.Getenv("USER"))

	// TextHandler
	textSLogger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	textSLogger.Info("hello world", "user", os.Getenv("USER"))

	// JSONHandler
	jsonSLogger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	jsonSLogger.Info("hello world", "user", os.Getenv("USER"))

	// LogAttr
	slog.LogAttrs(context.Background(), slog.LevelInfo, "hello, world", slog.String("user", os.Getenv("USER")))
}
