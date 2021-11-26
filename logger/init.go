package logger

import (
	"io"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/hellojqk/helper/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	viper.SetDefault("logger", Config{
		Console: ConsoleLogger{
			NoColor: false,
			Level:   "info",
		},
		Files: []FileLogger{{
			Name:       "app.log",
			Path:       "./log/",
			Level:      "info",
			MaxSize:    100,
			MaxBackups: 10,
			MaxAge:     30,
			Compress:   true,
		}},
	})
	util.WaitInitFuncsAdd(InitLogger)
}

func InitLogger() (err error) {
	err = load()
	if err != nil {
		return
	}
	viper.OnConfigChange(func(in fsnotify.Event) {
		load()
	})
	return nil
}

func load() error {
	conf := &Config{}
	err := viper.UnmarshalKey("logger", conf)
	if err != nil {
		log.Error().Err(err).Msg("Logger配置加载失败")
		return err
	}
	initLogger(conf)
	log.Info().Msg("Logger配置加载成功")
	return nil
}

func initLogger(conf *Config) {

	writers := make([]io.Writer, 0, 2)

	consoleLevel, _ := zerolog.ParseLevel(conf.Console.Level)

	consoleLoggerWriter := &FilteredWriter{
		Level: consoleLevel,
		Writer: zerolog.MultiLevelWriter(zerolog.ConsoleWriter{
			TimeFormat: time.RFC3339,
			Out:        os.Stdout,
			NoColor:    conf.Console.NoColor}),
	}

	writers = append(writers, consoleLoggerWriter)

	for _, loggerFile := range conf.Files {
		os.MkdirAll(loggerFile.Path, 0731)
		fileLogger := &lumberjack.Logger{
			Filename:   loggerFile.Path + loggerFile.Name,
			MaxSize:    loggerFile.MaxSize,
			MaxBackups: loggerFile.MaxBackups,
			MaxAge:     loggerFile.MaxAge,
			Compress:   loggerFile.Compress,
		}

		fileLevel, _ := zerolog.ParseLevel(loggerFile.Level)
		fileLoggerWriter := &FilteredWriter{
			Level: fileLevel,
			Writer: zerolog.MultiLevelWriter(zerolog.ConsoleWriter{
				TimeFormat: time.RFC3339,
				Out:        fileLogger,
				NoColor:    true}),
		}

		writers = append(writers, fileLoggerWriter)
	}

	log.Logger = log.Output(zerolog.MultiLevelWriter(writers...))
}
