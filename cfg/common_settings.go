package cfg

import (
	"github.com/olebedev/config"
)

type Colors struct {
	Background      string
	BorderFocusable string
	BorderFocused   string
	BorderNormal    string
	Checked         string
	HighlightFore   string
	HighlightBack   string
	Text            string
}

type Module struct {
	ConfigKey string
	Name      string
}

type Position struct {
	Height int
	Left   int
	Top    int
	Width  int
}

type Common struct {
	Colors
	Module
	Position

	Enabled         bool
	RefreshInterval int
	Title           string
}

func NewCommonSettingsFromYAML(name, configKey string, ymlConfig *config.Config) *Common {
	common := Common{
		Colors: Colors{
			Background:      ymlConfig.UString("wtf.colors.background", "black"),
			BorderFocusable: ymlConfig.UString("wtf.colors.border.focusable"),
			BorderFocused:   ymlConfig.UString("wtf.colors.border.focused"),
			BorderNormal:    ymlConfig.UString("wtf.colors.border.normal"),
			Checked:         ymlConfig.UString("wtf.colors.checked"),
			HighlightFore:   ymlConfig.UString("wtf.colors.highlight.fore"),
			HighlightBack:   ymlConfig.UString("wtf.colors.highlight.back"),
			Text:            ymlConfig.UString("wtf.colors.text", "white"),
		},

		Module: Module{
			ConfigKey: configKey,
			Name:      name,
		},

		Position: Position{},
	}

	return &common
}
