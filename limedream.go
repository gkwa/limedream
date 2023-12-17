package limedream

import (
	"log/slog"
	"os"

	"github.com/taylormonacelli/littlecow"
)

func Main() int {
	run(slog.LevelDebug)
	return 0
}

func run(level slog.Level) {
	logLevel := &slog.LevelVar{} // INFO
	logLevel.Set(level)
	opts := slog.HandlerOptions{
		AddSource:   true,
		Level:       logLevel,
		ReplaceAttr: littlecow.Replace,
	}
	handler := slog.NewTextHandler(os.Stderr, &opts)
	logger := slog.New(handler)

	slog.SetDefault(logger)
	slog.Debug("my debug message", "test", "test")
	slog.Error("my error message", "test", "test")
}
