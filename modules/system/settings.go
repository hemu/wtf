package system

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "system"

type Settings struct {
	common *cfg.Common
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config) *Settings {
	settings := Settings{
		common: cfg.NewCommonSettingsFromYAML(name, configKey, ymlConfig),
	}

	return &settings
}
