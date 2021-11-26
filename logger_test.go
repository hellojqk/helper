package helper

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hellojqk/helper/logger"
	_ "github.com/hellojqk/helper/logger/zerolog"
	"github.com/hellojqk/helper/util"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	viper.SetConfigType("yml")
	err := viper.ReadConfig(strings.NewReader(`
logger:
  console:
    noColor: true
    #日志等级 trace,debug,info,warn,error,fatal,panic 默认info
    level: debug
  files:
    - name: app.log
      #日志路径
      path: ./log/
      level: debug
      # mb
      maxSize: 3
      maxBackups: 3
      maxAge: 28
      compress: true
`))
	if err != nil {
		fmt.Println("ReadConfig Error", err)
		return
	}
	util.WaitInitFuncsExec()
	m.Run()
}

func TestLogger(t *testing.T) {
	t.Log(logger.Conf)
	log.Debug().Msg("来自zerolog")
}
