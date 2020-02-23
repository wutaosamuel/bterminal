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
	utils.WriteLog(e.Logger, e.LogPath, logInfo)
}

// WriteLogFunc write log with func
func (e *Exec) WriteLogFunc(logFunc utils.LogActCallback) {
	utils.WriteLogFunc(e.Logger, e.LogPath, logFunc)
}

// ReadLog read log
func (e *Exec) ReadLog() (string, error) {
	return utils.ReadLog(e.LogPath)
}
