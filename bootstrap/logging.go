package bootstrap

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

// init installs a text handler before env is loaded so slog calls made
// during NewEnv() do not fall back to the stdlib log bridge and end up
// with an inconsistent format. InitSlog overrides this later based on
// the configured LogLevel.
func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))
}

// InitSlog configures the process-wide slog default.
// Production uses the JSON handler; other environments use the text
// handler for human readability. The level is taken from env.LogLevel
// (mapstructure key LOG_LEVEL).
func InitSlog(env *Env) {
	level := parseLevel(env.LogLevel)
	opts := &slog.HandlerOptions{Level: level, AddSource: false}

	var h slog.Handler
	if strings.EqualFold(env.AppEnv, "production") {
		h = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		h = slog.NewTextHandler(os.Stdout, opts)
	}

	slog.SetDefault(slog.New(h))
	slog.Info("slog initialized", "app_env", env.AppEnv, "configured_level", env.LogLevel, "effective_level", level.String())
}

func parseLevel(s string) slog.Level {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "debug":
		return slog.LevelDebug
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

// Fatal is equivalent to slog.Error followed by os.Exit(1). slog has no Fatal level of its own.
func Fatal(msg string, args ...any) {
	slog.Error(msg, args...)
	os.Exit(1)
}

// FatalContext is the same as Fatal but accepts a context.
func FatalContext(ctx context.Context, msg string, args ...any) {
	slog.ErrorContext(ctx, msg, args...)
	os.Exit(1)
}
