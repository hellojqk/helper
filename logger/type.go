package logger

import "github.com/rs/zerolog"

type Config struct {
	Console ConsoleLogger `json:"console" mapstructure:"console"`
	Files   []FileLogger  `json:"files" mapstructure:"files"`
}
type ConsoleLogger struct {
	Level   string
	NoColor bool
}
type FileLogger struct {
	Name       string
	Path       string
	Level      string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

type FilteredWriter struct {
	Writer zerolog.LevelWriter
	Level  zerolog.Level
}

func (w *FilteredWriter) Write(p []byte) (n int, err error) {
	return w.Writer.Write(p)
}
func (w *FilteredWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if level >= w.Level {
		return w.Writer.WriteLevel(level, p)
	}
	return len(p), nil
}
