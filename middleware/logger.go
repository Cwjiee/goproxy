package middleware

import (
	"log"
	"os"
)

func InitLogger() {
	log.SetOutput(os.Stdout)
	log.SetPrefix("Goproxy: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
}
