package utils

import (
	"log"
	"os"
)

func InitLogger() {
	log.SetOutput((os.Stderr))
	log.SetFlags((log.LstdFlags | log.Lshortfile))
}
