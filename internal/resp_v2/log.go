package respv2

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

var log *slog.Logger

func init() {
	file, err := os.OpenFile("./log.txt", os.O_RDWR|os.O_CREATE, 0644)
	var logOutput io.Writer

	if err != nil {
		slog.Default().Warn(fmt.Sprintf("Failed to create log file for resp parser, using only stdout %s", err.Error()))
		logOutput = os.Stdout
	} else {
		logOutput = io.MultiWriter(os.Stdout, file)
	}
	log = slog.New(slog.NewTextHandler(logOutput, nil)).With(slog.Group("context", "package", "resp-parser"))
}
