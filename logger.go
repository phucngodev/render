package render

import (
	"log/slog"
	"os"
	"sync"
)

var (
	once    sync.Once
	_logger *slog.Logger
)

// logger return default slog logger for logging error detail in render.Error().
func logger() *slog.Logger {
	once.Do(func() {
		_logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	})

	return _logger
}

func logWithTrace(requestId string) *slog.Logger {
	return logger().With(slog.String("request_id", requestId))
}
