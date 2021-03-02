package vari

import (
	"github.com/easysoft/zmanager/pkg/model"
)

var (
	Config = model.NewConfig()

	ExeDir     string
	WorkDir    string
	ConfigFile string
	LogFile    string
	Language   string

	StartZTFService bool
	StartZDService  bool
)
