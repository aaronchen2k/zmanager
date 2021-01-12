package service

import (
	"errors"
	"fmt"
	configUtils "github.com/easysoft/zmanager/pkg/config"
	commonUtils "github.com/easysoft/zmanager/pkg/utils/common"
	constant "github.com/easysoft/zmanager/pkg/utils/const"
	fileUtils "github.com/easysoft/zmanager/pkg/utils/file"
	i118Utils "github.com/easysoft/zmanager/pkg/utils/i118"
	shellUtils "github.com/easysoft/zmanager/pkg/utils/shell"
	"github.com/easysoft/zmanager/pkg/utils/vari"
	"github.com/mholt/archiver/v3"
	"log"
	"strconv"
	"strings"
)

func CheckUpgrade(app string) {
	versionFile := vari.WorkDir + "version.txt"
	versionUrl := fmt.Sprintf(constant.VersionDownloadURL, app)
	commonUtils.Download(versionUrl, versionFile)

	content := strings.TrimSpace(fileUtils.ReadFile(versionFile))
	newVersion, _ := strconv.ParseFloat(content, 64)
	if (app == "ztf" && vari.Config.ZTFVersion < newVersion) ||
		(app == "ztf" && vari.Config.ZDVersion < newVersion) {

		log.Println(i118Utils.I118Prt.Sprintf("find_new_ver", content))

		newVersionStr := fmt.Sprintf("%.1f", newVersion)
		pass, err := downloadApp(app, newVersionStr)
		if pass && err == nil {
			restartApp(app, newVersionStr)
		}
	}
}

func downloadApp(app string, version string) (pass bool, err error) {
	os := commonUtils.GetOs()
	if commonUtils.IsWin() {
		os = fmt.Sprintf("%s%d", os, strconv.IntSize)
	}
	url := fmt.Sprintf(constant.PackageDownloadURL, app, version, os, app)

	dir := vari.WorkDir + version

	pth := dir + ".zip"
	err = commonUtils.Download(url, pth)
	if err != nil {
		return
	}

	md5Url := url + ".md5"
	md5Pth := pth + ".md5"
	err = commonUtils.Download(md5Url, md5Pth)
	if err != nil {
		return
	}

	pass = checkMd5(pth, md5Pth)
	if !pass {
		msg := i118Utils.I118Prt.Sprintf("fail_md5_check", pth)
		log.Println(msg)
		err = errors.New(msg)
		return
	}

	fileUtils.RmDir(dir)
	fileUtils.MkDirIfNeeded(dir)
	err = archiver.Unarchive(pth, dir)

	if err != nil {
		log.Println(i118Utils.I118Prt.Sprintf("fail_unzip", pth))
		return
	}

	return
}

func restartApp(app string, newVersion string) (err error) {
	currExePath := vari.ExeDir + constant.AppName
	bakExePath := currExePath + "_bak"
	newExePath := vari.WorkDir + newVersion + constant.PthSep + constant.AppName + constant.PthSep + constant.AppName
	if commonUtils.IsWin() {
		currExePath += ".exe"
		bakExePath += ".exe"
		newExePath += ".exe"
	}

	var oldVersion float64
	if app == "ztf" {
		oldVersion = vari.Config.ZTFVersion
		vari.Config.ZTFVersion, _ = strconv.ParseFloat(newVersion, 64)
	} else if app == "zd" {
		oldVersion = vari.Config.ZDVersion
		vari.Config.ZTFVersion, _ = strconv.ParseFloat(newVersion, 64)
	}
	log.Println(i118Utils.I118Prt.Sprintf("success_upgrade", oldVersion, newVersion))

	// update config file
	configUtils.SaveConfig(vari.Config)

	return
}

func checkMd5(filePth, md5Pth string) (pass bool) {
	expectVal := fileUtils.ReadFile(md5Pth)
	actualVal, _ := shellUtils.ExeSysCmd("md5sum " + filePth + " | awk '{print $1}'")

	return strings.TrimSpace(actualVal) == strings.TrimSpace(expectVal)
}
