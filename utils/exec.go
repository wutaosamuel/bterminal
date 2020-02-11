package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"

	cron "github.com/robfig/cron"
	uuid "github.com/satori/go.uuid"
)

// TODO: recover jobs
// TODO: cycle job is not work on done

// Exec type for do exec
type Exec struct {
	Name    string // the name of jobs
	nameID  string // the unique ID for each job
	Command string // command required to execute

	Cron   *cron.Cron // does job need to schedule
	time   string     // schedule time of job
	CronOP CronOP     // cron operation for job
	done   bool       // does it finish
}

// CronOP is cron option
type CronOP uint8

// CronOP const
const (
	CronNull CronOP = 1 << iota
	CronStart
	CronEnd
)

/////////////////// Setter&&Getter ///////////////////

// NewExec crete a new Exec struct
func NewExec() *Exec {
	uuid := uuid.Must(uuid.NewV4()).String()
	return &Exec{nameID: uuid}
}

// SetTime by cron like format schedule
func (e *Exec) SetTime(m string, h string, d string, mon string, w string) {
	e.time = m + " " + h + " " + d + " " + mon + " " + w
}

// GetNameID get uuid for job
func (e *Exec) GetNameID() string { return e.nameID }

// GetNameID8b get the first 8 bits uuid string
func (e *Exec) GetNameID8b() string { return e.nameID[:8] }

// GetTime get time of job
func (e *Exec) GetTime() string { return e.time }

// GetDone get done signal
func (e *Exec) GetDone() bool { return e.done }

/////////////////// Main ///////////////////

// DoExec execute with Exec struct
func (e *Exec) DoExec() {
	// check
	if e.Name == "" {
		e.Name = "BackConsole-" + e.GetNameID8b()
	}
	if e.Command == "" {
		// do nothing
		return
	}

	// exec
	logName := e.Name + e.GetNameID8b() + ".log"
	DoExecute(logName, e.Command)
	e.done = true
}

// StartCron do schedule of Exec
func (e *Exec) StartCron() {
	// check
	//if e.CronOP != CronStart || e.CronOP != CronEnd {
	//return
	//}
	if e.time == "" {
		return
	}
	// FIXME:	cron op can not need

	// start cron
	e.Cron = cron.New()
	e.Cron.AddFunc(e.time, func() { e.DoExec() })
	e.Cron.Start()
	log.Println(e.Name + " cron start!")
	e.done = true
}

// StopCron to stop job
func (e *Exec) StopCron() {
	if e.CronOP != CronEnd {
		return
	}
	e.Cron.Stop()
	e.done = false
	log.Println(e.Name + " cron has stopped!")
}

// DoExecute execute command and log recording
// phrase command to args[] first
// exec comand
// output log
func DoExecute(logName string, command string) {
	args := strings.Fields(strings.TrimSpace(command))
	cmd := exec.Command(args[0], args[1:]...)
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)

	f, err := os.OpenFile(
		logName,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalf("Error: open log file - %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error: cannot run output - %v", err)
	}
	log.Printf("The output - %s\n", out)
}
