package main

import (
	"flag"
	"fmt"
	configUtils "github.com/easysoft/zmanager/pkg/config"
	"github.com/easysoft/zmanager/pkg/program"
	commonUtils "github.com/easysoft/zmanager/pkg/utils/common"
	constant "github.com/easysoft/zmanager/pkg/utils/const"
	i118Utils "github.com/easysoft/zmanager/pkg/utils/i118"
	"github.com/easysoft/zmanager/pkg/utils/vari"
	"github.com/kardianos/service"
	"log"
	"os"
)

var (
	flagSet *flag.FlagSet
	action  string
)

func main() {
	flagSet = flag.NewFlagSet("zmanager", flag.ContinueOnError)
	flagSet.StringVar(&action, "a", "", "")
	flagSet.BoolVar(&vari.StartZTFService, "ztf", false, "")
	flagSet.BoolVar(&vari.StartZDService, "zd", true, "")
	flagSet.StringVar(&vari.Language, "l", "", "")
	flagSet.Parse(os.Args[1:])

	log.Println(fmt.Sprintf("StartZTFService=%t, StartZDService=%t",
		vari.StartZTFService, vari.StartZDService))

	configUtils.Init()

	if action != "" && !commonUtils.StrInArr(action, constant.Actions) {
		log.Println(i118Utils.I118Prt.Sprintf("invalid_actions", action, service.ControlAction))
		return
	}

	options := make(service.KeyValue)
	options["Restart"] = "on-success"
	options["SuccessExitStatus"] = "1 2 8 SIGKILL"
	config := &service.Config{
		Name:        constant.AppName,
		DisplayName: constant.AppName,
		Description: constant.AppName + " service.",
		Dependencies: []string{
			"Requires=network.target",
			"After=network-online.target syslog.target"},
		Option: options,
	}

	prg := &program.Program{}
	srv, err := service.New(prg, config)
	if err != nil {
		log.Fatal(err)
	}
	errs := make(chan error, 5)
	program.Logger, err = srv.Logger(errs)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			err := <-errs
			if err != nil {
				log.Print(err)
			}
		}
	}()

	if action != "" {
		err := service.Control(srv, action)
		if err != nil {
			log.Println(i118Utils.I118Prt.Sprintf("valid_actions", service.ControlAction))
			log.Fatal(err)
		}
		return
	}

	err = srv.Run()
	if err != nil {
		program.Logger.Error(err)
	}
}
