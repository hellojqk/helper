package zerolog

import "github.com/rs/zerolog"

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
