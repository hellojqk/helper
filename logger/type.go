package logger

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
