package file

import (
	"log"
)

type FileEventLogger struct {}

func NewFileEventLogger() *FileEventLogger {
	return &FileEventLogger{}
}


func (fel FileEventLogger) writeDelete(key string) {
	log.Println("Del")
}
func (fel FileEventLogger) writePut(key, value string) {
	log.Println("Put")
}

func (fel FileEventLogger) writeGet(key string) {
	log.Println("Get")
}

