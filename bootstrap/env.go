package bootstrap

import (
	"strings"

	"github.com/anhvanhoa/service-core/boostrap/config"
)

type Env struct {
	NodeEnv string `mapstructure:"node_env"`

	UrlDb string `mapstructure:"url_db"`

	NameService   string `mapstructure:"name_service"`
	PortGprc      int    `mapstructure:"port_grpc"`
	HostGprc      string `mapstructure:"host_grpc"`
	IntervalCheck string `mapstructure:"interval_check"`
	TimeoutCheck  string `mapstructure:"timeout_check"`
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
