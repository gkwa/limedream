package limedream

import (
	"log/slog"
	"os"

	"github.com/taylormonacelli/littlecow"
)

func Main() int {
	slog.Debug("limedream", "test", true)

	return 0
}

func run(level slog.Level) {
	logLevel := &slog.LevelVar{} // INFO
	logLevel.Set(level)
	opts := slog.HandlerOptions{
		AddSource: true,
		Level:     logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey && len(groups) == 0 {
				return slog.Attr{}
			}
			if a.Key == slog.SourceKey {
				source, _ := a.Value.Any().(*slog.Source)
				if source != nil {
					setPartialPath(source)
				}
			}
			return a
		},
	}
	handler := slog.NewTextHandler(os.Stderr, &opts)
	logger := slog.New(handler)

	slog.SetDefault(logger)
}
