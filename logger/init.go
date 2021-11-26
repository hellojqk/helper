package logger

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/hellojqk/helper/util"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var onChange []func()

func OnChange(f func()) {
	onChange = append(onChange, f)
}

var Conf *Config

func init() {
	onChange = make([]func(), 0)
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
		log.Println(errors.WithMessage(err, "Logger配置加载失败"))
		return err
	}
	Conf = conf
	log.Println("Logger配置加载成功")
	for _, change := range onChange {
		change()
	}
	return nil
}
