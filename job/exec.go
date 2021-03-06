package job

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	cron "github.com/robfig/cron"
	uuid "github.com/satori/go.uuid"
)

// TODO: cycle job is not work on done
// TODO: avoiding cmd windows pop up on Windows 10
// TODO: replace RWLocker to file locker

// Exec type for do exec
type Exec struct {
	Name    string // the name of jobs
	NameID  string // the unique ID for each job
	Command string // command required to execute
	LogName string // log path/name.log

	*sync.RWMutex             // Read & write lock
	*os.File								  // File lock
	Logger        *log.Logger // logger for exec

	Cron *cron.Cron // does job need to schedule
	Time string     // schedule time of job
}

/////////////////// Setter&&Getter ///////////////////

// NewExecS create a new Exec struct
// Init required after create a new Exec
func NewExecS() *Exec {
	return &Exec{
		"",
		"",
		"",
		"",
		&sync.RWMutex{},
		&os.File{},
		&log.Logger{},
		cron.New(),
		"",
	}
}

// NewExec create a new Exec struct
func NewExec() *Exec {
	uuid := uuid.Must(uuid.NewV4()).String()
	return &Exec{
		uuid,
		"",
		"",
		"",
		&sync.RWMutex{},
		&os.File{},
		&log.Logger{},
		cron.New(),
		""}
}

// RecoverDat recover exec from dat
func RecoverDat(d Dat) map[string]Exec {
	e := make(map[string]Exec)
	for id, job := range d.Jobs {
		e[id] = Exec{
			job.Name,
			id,
			job.Command,
			job.LogName,
			&sync.RWMutex{},
			&os.File{},
			&log.Logger{},
			cron.New(),
			job.Time,
		}
	}
	return e
}

// Init init exec
func (e *Exec) Init() error {
	e.NameID = uuid.Must(uuid.NewV4()).String()
	// set log path
	e.SetLogName()
	return nil
}

// SetCronTime by cron like format schedule
func (e *Exec) SetCronTime(m string, h string, d string, mon string, w string) {
	e.Time = m + " " + h + " " + d + " " + mon + " " + w
}

// SetLogName set log path
func (e *Exec) SetLogName() {
	if len(e.LogName) < 4 {
		e.LogName = filepath.Join(e.LogName, e.CreateLogName())
		return
	}
	if e.LogName[len(e.LogName)-4:] != ".log" {
		e.LogName = filepath.Join(e.LogName, e.CreateLogName())
		return
	}
	return
}

// GetNameID get uuid for job
func (e *Exec) GetNameID() string { return e.NameID }

// GetNameID8b get the first 8 bits uuid string
func (e *Exec) GetNameID8b() string { return e.NameID[:8] }

/////////////////// Main ///////////////////

// CreateLogName get log name
func (e *Exec) CreateLogName() string { return e.Name + "_" + e.GetNameID() + ".log" }

// DoExec execute with Exec struct
func (e *Exec) DoExec() {
	// check
	if e.Name == "" {
		e.Name = "BTerminal-" + e.GetNameID8b()
	}
	if e.Command == "" {
		// do nothing
		return
	}

	// exec
	go DoExecute(e.LogName, e.Command)
}

// StartCron do schedule of Exec
func (e *Exec) StartCron() {
	// start cron
	e.Cron = cron.New()
	if _, err := e.Cron.AddFunc(e.Time, func() { e.DoExec() }); err != nil {
		e.WriteLog(err)
	}
	e.WriteLog(e.Name + " cron start!")
	e.Cron.Start()
}

// StopCron to stop job
func (e *Exec) StopCron() {
	e.Cron.Stop()
	e.WriteLog(e.Name + " cron has stopped!")
}

// DeleteLog to delete log
func (e *Exec) DeleteLog() error {
	e.Lock()
	err := os.RemoveAll(e.LogName)
	e.Unlock()
	return err
}
