package constant

import (
	"fmt"
	"os"
	"os/user"
)

const (
	AppName = "zmanager"
)

var (
	PthSep = string(os.PathSeparator)

	userProfile, _ = user.Current()

	LanguageDefault = "en"
	LanguageEN      = "en"
	LanguageZH      = "zh"

	EnRes = fmt.Sprintf("res%sen%smessages.json", string(os.PathSeparator), string(os.PathSeparator))
	ZhRes = fmt.Sprintf("res%szh%smessages.json", string(os.PathSeparator), string(os.PathSeparator))

	LogDir = fmt.Sprintf("log%s", string(os.PathSeparator))

	Actions = []string{"start", "stop", "restart", "install", "uninstall"}

	QiNiuURL           = "https://dl.cnezsoft.com/"
	VersionDownloadURL = QiNiuURL + "%s/version.txt"
	PackageDownloadURL = QiNiuURL + "%s/%s/%s/%s.zip"

	ZTF  = "ztf"
	ZD   = "zd"
	Apps = []string{ZTF, ZD}
)
