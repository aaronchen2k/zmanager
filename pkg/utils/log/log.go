package logUtils

import (
	"log"
	"os"
)

func Init(file *os.File) {
	log.SetOutput(file) // will log as root user under service
	log.SetPrefix("[Error]")
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
}
