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
	"regexp"
	"strconv"
	"strings"
)

func ExeSysCmd(cmdStr string) (string, error) {
	var cmd *exec.Cmd
	if commonUtils.IsWin() {
		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	output := out.String()

	return output, err
}

func GetPrecess(app string) (string, error) {
	var cmd *exec.Cmd

	tmpl := ""
	cmdStr := ""
	if commonUtils.IsWin() {
		tmpl = `tasklist`
		cmdStr = fmt.Sprintf(tmpl)

		cmd = exec.Command("cmd", "/C", cmdStr)
	} else {
		tmpl = `ps -ef | grep "%s" | grep -v "grep" | awk '{print $2}'`
		cmdStr = fmt.Sprintf(tmpl, app)

		cmd = exec.Command("/bin/bash", "-c", cmdStr)
	}

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	output := ""
	if commonUtils.IsWin() {
		arr := strings.Split(out.String(), "\n")
		for _, line := range arr {
			if strings.Index(line, app+".exe") > -1 {
				arr2 := regexp.MustCompile(`\s+`).Split(line, -1)
				output = arr2[1]
				break
			}
		}
	} else {
		output = out.String()
	}

	return output, err
}

func KillPrecess(app string) (string, error) {
	var cmd *exec.Cmd

	tmpl := ""
	cmdStr := ""
	if commonUtils.IsWin() {
		// tasklist | findstr ztf.exe
		tmpl = `taskkill.exe /f /im %s.exe`
		cmdStr = fmt.Sprintf(tmpl, app)

		cmd = exec.Command("cmd", "/C", cmdStr)
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
		tmpl = `start %s -%s %d > %snohup.%s.log 2>&1`
		cmdStr = fmt.Sprintf(tmpl, execPath, portTag, portNum, vari.WorkDir, app)

		cmd = exec.Command("cmd", "/C", cmdStr)

		//log := filepath.Join(vari.WorkDir, "nohup."+app+".log")
		//f, _ := os.Create(log)
		//
		//cmd.Stdout = f
		//cmd.Stderr = f

	} else {
		cmd = exec.Command("nohup", execPath, "-"+portTag, strconv.Itoa(portNum))

		log := filepath.Join(vari.WorkDir, "nohup."+app+".log")
		f, _ := os.Create(log)

		cmd.Stdout = f
		cmd.Stderr = f
	}

	cmd.Dir = path.Dir(execPath)
	err := cmd.Start()
	return "", err
}
