package zerolog

import (
	"io"
	"os"
	"time"

	"github.com/hellojqk/helper/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	logger.OnChange(initLogger)
}

func initLogger() {

	writers := make([]io.Writer, 0, 2)

	consoleLevel, _ := zerolog.ParseLevel(logger.Conf.Console.Level)

	consoleLoggerWriter := &FilteredWriter{
		Level: consoleLevel,
		Writer: zerolog.MultiLevelWriter(zerolog.ConsoleWriter{
			TimeFormat: time.RFC3339,
			Out:        os.Stdout,
			NoColor:    logger.Conf.Console.NoColor}),
	}

	writers = append(writers, consoleLoggerWriter)

	for _, loggerFile := range logger.Conf.Files {
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
