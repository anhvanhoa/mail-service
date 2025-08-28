package bootstrap

import (
	"strings"

	"github.com/anhvanhoa/service-core/boostrap/config"
)

type Env struct {
	NODE_ENV string

	URL_DB string

	NAME_SERVICE   string
	PORT_GRPC      int
	HOST_GRPC      string
	INTERVAL_CHECK string
	TIMEOUT_CHECK  string
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
	return strings.ToLower(env.NODE_ENV) == "production"
}
