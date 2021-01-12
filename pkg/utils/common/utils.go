package commonUtils

import (
	"os"
	"os/user"
	"runtime"
)

func GetOs() string {
	osName := runtime.GOOS

	if osName == "darwin" {
		return "mac"
	} else {
		return osName
	}
}
func IsWin() bool {
	return GetOs() == "windows"
}
func IsLinux() bool {
	return GetOs() == "linux"
}
func IsMac() bool {
	return GetOs() == "mac"
}

func IsRelease() bool {
	if _, err := os.Stat("res"); os.IsNotExist(err) {
		return true
	}

	return false
}

func GetUserHome() string {
	userProfile, _ := user.Current()
	home := userProfile.HomeDir
	return home
}

func FindInArr(str string, arr []string) (bool, int) {
	for index, s := range arr {
		if str == s {
			return true, index
		}
	}

	return false, -1
}

func StrInArr(str string, arr []string) bool {
	found, _ := FindInArr(str, arr)
	return found
}
