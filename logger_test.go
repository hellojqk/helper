package helper

import (
	"fmt"
	"strings"
	"testing"

	_ "github.com/hellojqk/helper/logger"
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
    level: info
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
	log.Debug().Msg("来自zerolog")
	log.Info().Msg("来自zerolog")
	log.Warn().Msg("来自zerolog")
}
