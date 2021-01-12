package stdinUtils

import (
	"bufio"
	i118Utils "github.com/easysoft/zmanager/pkg/utils/i118"
	"log"
	"os"
	"regexp"
	"strings"
)

func GetInput(regx string, defaultVal string, fmtStr string, params ...interface{}) string {
	var ret string

	msg := i118Utils.I118Prt.Sprintf(fmtStr, params...)

	for {
		log.Println("\n" + msg)
		Scanf(&ret)

		if strings.TrimSpace(ret) == "" && defaultVal != "" {
			ret = defaultVal

			log.Println(ret)
		}

		temp := strings.ToLower(ret)
		if temp == "exit" {
			os.Exit(0)
		}

		if regx == "" {
			return ret
		}

		pass, _ := regexp.MatchString("^"+regx+"$", temp)
		msg := "invalid_input"

		if pass {
			return ret
		} else {
			ret = ""
			log.Println(i118Utils.I118Prt.Sprintf(msg))
		}
	}
}

func Scanf(a *string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	*a = string(data)
}
