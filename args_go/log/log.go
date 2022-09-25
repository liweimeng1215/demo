package log

import "log"

func Fetal(format string, v ...any) {
	log.Fatalf("[FETAL] "+format+"\n", v)
}

func Error(format string, v ...any) {
	log.Printf("[ERROR] "+format+"\n", v)
}

func Warn(format string, v ...any) {
	log.Printf("[WARN] "+format+"\n", v)
}

func Info(format string, v ...any) {
	log.Printf("[INFO] "+format+"\n", v)
}

func Debug(format string, v ...any) {
	log.Printf("[DEBUG] "+format+"\n", v)
}
