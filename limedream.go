package limedream

import (
	"log/slog"
	"os"

	"github.com/taylormonacelli/littlecow"
)

func Main() int {
	run(slog.LevelDebug)
	run(slog.LevelInfo)
	return 0
}

func run(level slog.Level) {
	opts := slog.HandlerOptions{
		AddSource:   true,
		Level:       level,
		ReplaceAttr: littlecow.Replace,
	}
	handler := slog.NewTextHandler(os.Stderr, &opts)
	logger := slog.New(handler)

	slog.SetDefault(logger)
	slog.Debug("my debug message", "test", "test")
	slog.Error("my error message", "test", "test")
}
