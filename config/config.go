package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type state string

const (
	stateLocal state = "dev"
	stateDEV   state = "dev"
	stateSIT   state = "sit"
	statePROD  state = "prod"
)

type ConstantViper struct {
	State       state
	ProjectCode string
	Maintenance bool
}

func (cv *ConstantViper) SetState(s *string) {
	switch *s {
	case "local", "localhost", "l":
		cv.State = stateLocal
	case "dev", "develop", "development", "d":
		cv.State = stateDEV
	case "sit", "staging", "s":
		cv.State = stateSIT
	case "prod", "production", "p":
		cv.State = statePROD
	default:
		cv.State = stateLocal
	}
}

func (cv *ConstantViper) Init() {
	viper.SetConfigFile("config.yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	cv.binding()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		//log.Infoln("config file changed:", e.Name)
		cv.binding()
	})
}

func (cv *ConstantViper) binding() {
	sub := viper.Sub(string(cv.State))

	cv.ProjectCode = sub.GetString("project_code")
	cv.Maintenance = sub.GetBool("maintenance")
}
