package shellUtils

import (
	"bytes"
	"fmt"
	commonUtils "github.com/easysoft/zmanager/pkg/utils/common"
	constant "github.com/easysoft/zmanager/pkg/utils/const"
	"github.com/easysoft/zmanager/pkg/utils/vari"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
)

func ExeSysCmd(cmdStr string) (string, error) {
	var cmd *exec.Cmd
	if commonUtils.IsWin() {
		cmd = exec.Command(cmdStr)
	} else {
		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	return out.String(), err
}

func KillPrecess(app string) (string, error) {
	var cmd *exec.Cmd

	tmpl := ""
	cmdStr := ""
	if commonUtils.IsWin() {
		tmpl = `taskkill.exe /f /im %s.exe`
		cmdStr = fmt.Sprintf(tmpl, app)

		cmd = exec.Command(cmdStr)
		// cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		tmpl = `ps -ef | grep "%s" | grep -v "grep" | awk '{print $2}' | xargs kill -9`
		cmdStr = fmt.Sprintf(tmpl, app)

		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	output := out.String()

	return output, err
}

func StartPrecess(execPath string, app string) (string, error) {
	portTag := ""
	portNum := 0
	if app == constant.ZTF {
		portTag = "P"
		portNum = 8848
	} else if app == constant.ZenData {
		portTag = "p"
		portNum = 8849
	}

	tmpl := ""
	cmdStr := ""
	var cmd *exec.Cmd
	if commonUtils.IsWin() {
		tmpl = `start /b %s.exe -%s %d > %s\%s`
		cmdStr = fmt.Sprintf(tmpl, execPath, portTag, portNum, vari.WorkDir, app)

		cmd = exec.Command(cmdStr)
		// cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		cmd = exec.Command("nohup", execPath, "-"+portTag, strconv.Itoa(portNum))
		cmd.Dir = path.Dir(execPath)

		log := filepath.Join(vari.WorkDir, app+".nohup.log")
		f, _ := os.Create(log)

		cmd.Stdout = f
		cmd.Stderr = f
	}

	err := cmd.Start()
	return "", err
}
