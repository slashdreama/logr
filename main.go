package logr

import "log"

func Info(msg string) {
	log.Printf("INFO ✅ :> %v", msg)
}

func Warn(msg string) {
	log.Printf("WARN ❗️:> %v", msg)
}

func Eroor(msg string) {
	log.Printf("ERROR ❌ :> %v", msg)
}
