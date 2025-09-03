package bootstrap

import (
	"strings"

	"github.com/anhvanhoa/service-core/boostrap/config"
)

type Env struct {
	NodeEnv string

	UrlDb string

	NameService   string
	PortGprc      int
	HostGprc      string
	IntervalCheck string
	TimeoutCheck  string
}

func NewEnv(env any) {
	setting := config.DefaultSettingsConfig()
	if setting.IsProduction() {
		setting.SetFile("prod.config")
	} else {
		setting.SetFile("dev.config")
	}
	config.NewConfig(setting, env)
}

func (env *Env) IsProduction() bool {
	return strings.ToLower(env.NodeEnv) == "production"
}
