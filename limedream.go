package limedream

import (
	"log/slog"
	"os"

	"github.com/taylormonacelli/littlecow"
)

func Main() int {
	var opts *slog.HandlerOptions
	var handler slog.Handler

	opts = littlecow.OpinionatedHandlerOptions(slog.LevelDebug, littlecow.RemoveTimestampAndTruncateSource)
	handler = slog.NewJSONHandler(os.Stderr, opts)
	run(handler)

	opts = littlecow.OpinionatedHandlerOptions(slog.LevelInfo, littlecow.TruncateSourcePath)
	handler = slog.NewTextHandler(os.Stderr, opts)
	run(handler)

	opts = littlecow.OpinionatedHandlerOptions(slog.LevelDebug, littlecow.TruncateSourcePath)
	handler = slog.NewTextHandler(os.Stderr, opts)
	run(handler)

	return 0
}

func run(handler slog.Handler) {
	logger := slog.New(handler)
	slog.SetDefault(logger)
	slog.Debug("debug", "debug", "debug")
	slog.Error("error", "error", "error")
}
