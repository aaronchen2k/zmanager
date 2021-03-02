package manageService

import (
	commonUtils "github.com/easysoft/zmanager/pkg/utils/common"
	constant "github.com/easysoft/zmanager/pkg/utils/const"
	shellUtils "github.com/easysoft/zmanager/pkg/utils/shell"
	"github.com/easysoft/zmanager/pkg/utils/vari"
	"log"
	"strings"
)

func CheckStatus(app string) {
	output, _ := shellUtils.GetProcess(app)
	output = strings.TrimSpace(output)

	if output != "" {
		return
	}

	startApp(app)
}

func startApp(app string) (err error) {
	appDir := vari.WorkDir + app + constant.PthSep

	newExePath := appDir + "latest" + constant.PthSep + app + constant.PthSep + app
	if commonUtils.IsWin() {
		newExePath += ".exe"
	}

	log.Println("Before StartProcess " + newExePath)
	shellUtils.StartProcess(newExePath, app)

	return
}
