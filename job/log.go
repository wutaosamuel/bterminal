package job

import (
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

// WriteLog write into log
func (e *Exec) WriteLog(logInfo string) {
	e.Lock()
	utils.WriteLog(e.Logger, e.LogName, logInfo)
	e.Unlock()
}

// WriteLogFunc write log with func
func (e *Exec) WriteLogFunc(logFunc utils.LogActCallback) {
	e.Lock()
	utils.WriteLogFunc(e.Logger, e.LogName, logFunc)
	e.Unlock()
}

// ReadLog read log
func (e *Exec) ReadLog() (string, error) {
	e.RLock()
	str, err := utils.ReadLog(e.LogName)
	e.RUnlock()
	return str, err
}
