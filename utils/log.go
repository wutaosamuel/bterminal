package utils

/*
 * log.go recode info by logging
 * the log file will be close and open frequently
 */

import (
	"io/ioutil"
	"log"
	"os"
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
type LogActCallback func(*log.Logger)

// WriteLog write into log
// Println only
func WriteLog(logger *log.Logger, logName string, logInfo interface{}) {
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
	logger = log.New(f, "", log.Ldate|log.Ltime|log.LUTC)
	logger.Println(logInfo)
	return
}

// WriteLogFunc write into log by func
func WriteLogFunc(logger *log.Logger, logName string, logFunc LogActCallback) {
	// open log file
	f, err := os.OpenFile(
		logName,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	logger = log.New(f, "", log.Ldate|log.Ltime|log.LUTC)
	logFunc(logger)
	return
}

// ReadLog read log
func ReadLog(logName string) (string, error) {
	f, err := ioutil.ReadFile(logName)
	if err != nil {
		return "", err
	}
	return string(f), nil
}
