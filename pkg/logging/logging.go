package logging

import "log"

func Info(message string) {
	log.Println("🔍 INFO: " + message)
}

func Done(message string) {
	log.Println("✅ DONE: " + message)
}

func Error(message string) {
	log.Println("❌ ERROR: " + message)
}

func Debug(message string) {
	log.Println("🐛 DEBUG: " + message)
}

func Warn(message string) {
	log.Println("⚠️ WARN: " + message)
}
