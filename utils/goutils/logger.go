package goutils

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
)

func SetLogger() {
	w := os.Stdout
	logger := slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			NoColor: !isatty.IsTerminal(w.Fd()),
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				// Remove time.
				if a.Key == slog.TimeKey && len(groups) == 0 {
					return slog.Attr{}
				}
				// Remove the directory from the source's filename.
				if a.Key == slog.SourceKey {
					source := a.Value.Any().(*slog.Source)
					source.File = filepath.Base(source.File)
				}
				return a
			},
		}),
	)

	slog.SetDefault(logger)
}
