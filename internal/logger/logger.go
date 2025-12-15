package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"
)

type syncWriter struct {
	mu sync.Mutex
	w  io.Writer
}

func (s *syncWriter) write(p []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.w.Write(p)
}

var Log *slog.Logger

func init() {
	file, err := os.OpenFile("./log.json", os.O_RDWR|os.O_CREATE, 0644)
	var logOutput syncWriter

	if err != nil {
		slog.Default().Warn(fmt.Sprintf("Failed to create log file for resp parser, using only stdout %s", err.Error()))
		logOutput = syncWriter{w: os.Stdout}
	} else {
		logOutput = syncWriter{w: io.MultiWriter(os.Stdout, file)}
	}
	Log = slog.New(slog.NewJSONHandler(logOutput.w, nil)).With(slog.Group("context", "package", "resp-parser"))
}
