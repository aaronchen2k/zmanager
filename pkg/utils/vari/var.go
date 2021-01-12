package vari

import (
	"github.com/easysoft/zmanager/pkg/model"
)

var (
	Config = model.Config{ZTFVersion: 1, ZDVersion: 1, Language: "en"}

	ExeDir     string
	WorkDir    string
	ConfigFile string
	LogFile    string
	Language   string
)
