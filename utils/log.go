package utils

import (
	"log"
	"os"
	"io/ioutil"
)

// LogFunction for doing log action
type LogFunction interface {
	// write log
	WriteLog()
	// write log by extra func
	WriteLogFunc()
	// read log
	ReadLog()
}

// LogActCallback for log call back
type LogActCallback func(logger *log.Logger)

// WriteLog write into log
// Println only
func WriteLog(logger *log.Logger, logName, logInfo string) {
	// open log file
	logger.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	f, err := os.OpenFile(
		logName,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	logger = log.New(f, "", log.Ldate | log.Ltime | log.LUTC)
	logger.Println(logInfo)
	return
}

// WriteLogFunc write into log by func
func WriteLogFunc(logger *log.Logger, logName string, f LogActCallback) error {

}

// ReadLog read log
func ReadLog(logName string) (string, error) {
	f, err := ioutil.ReadFile(logName)
	if err != nil {
		return "", err
	}
	return string(f), nil
}
