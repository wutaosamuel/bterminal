package job

import (
	"log"
	"os"

	// FIXME:
	// "github.com/wutaosamuel/bterminal/utils"
	"../utils"
)

// LogFunction interface
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
func (e *Exec) WriteLog(logInfo interface{}) {
	e.Lock()
	utils.WriteLog(e.Logger, e.LogName, logInfo)
	e.Unlock()
	return
}

// WriteLogFunc write into log by func
func (e *Exec) WriteLogFunc(logFunc LogActCallback) {
	// TODO: file lock
	// open log file
	f, err := os.OpenFile(
		e.LogName,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	e.Logger = log.New(f, "", log.Ldate|log.Ltime|log.LUTC)
	logFunc(e.Logger)
	return
}

// ReadLog read log
func (e *Exec) ReadLog() (string, error) {
	// TODO: file lock
	// avoiding if log is not exist
	isFile, err := utils.IsFile(e.LogName)
	if err != nil {
		return "", err
	}
	if !isFile {
		return "", nil
	}
	str, err := utils.ReadLog(e.LogName)
	return str, err
}
