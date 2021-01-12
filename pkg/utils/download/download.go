package download

import (
	"archive/zip"
	"fmt"
	errUtils "github.com/easysoft/zmanager/pkg/utils/err"
	i118Utils "github.com/easysoft/zmanager/pkg/utils/i118"
	"github.com/mholt/archiver/v3"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func Download(uri string, dst string) error {
	if strings.Index(uri, "?") < 0 {
		uri += "?"
	} else {
		uri += "&"
	}
	uri += fmt.Sprintf("&r=%d", time.Now().Unix())

	res, err := http.Get(uri)
	if err != nil {
		log.Println(i118Utils.I118Prt.Sprintf("download_fail", uri, err.Error()))
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(i118Utils.I118Prt.Sprintf("download_read_fail", uri, err.Error()))
	}

	err = ioutil.WriteFile(dst, bytes, 0666)
	if err != nil {
		log.Println(i118Utils.I118Prt.Sprintf("download_write_fail", dst, err.Error()))
	} else {
		log.Println(i118Utils.I118Prt.Sprintf("download_success", uri, dst))
	}

	return err
}

func GetZipSingleDir(path string) string {
	folder := ""
	z := archiver.Zip{}
	err := z.Walk(path, func(f archiver.File) error {
		if f.IsDir() {
			zfh, ok := f.Header.(zip.FileHeader)
			if ok {
				//logUtils.PrintTo("file: " + zfh.Name)

				if folder == "" && zfh.Name != "__MACOSX" {
					folder = zfh.Name
				} else {
					if strings.Index(zfh.Name, folder) != 0 {
						return errUtils.New("found more than one folder")
					}
				}
			}
		}
		return nil
	})

	if err != nil {
		return ""
	}

	return folder
}
