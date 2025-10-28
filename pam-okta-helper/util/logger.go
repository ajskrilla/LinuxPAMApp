// util/logger.go
package util

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	multi := io.MultiWriter(os.Stdout, file)
	Logger = log.New(multi, "", log.LstdFlags|log.Lshortfile)
	return nil
}
