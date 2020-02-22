package job

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
// TODO: set log path

// Exec type for do exec
type Exec struct {
	Name    string // the name of jobs
	nameID  string // the unique ID for each job
	Command string // command required to execute

	// TODO: mux lock for writing in log
	Cron *cron.Cron // does job need to schedule
	time string     // schedule time of job
	done bool       // does it finish
}

/////////////////// Setter&&Getter ///////////////////

// NewExec crete a new Exec struct
func NewExec() *Exec {
	uuid := uuid.Must(uuid.NewV4()).String()
	return &Exec{nameID: uuid}
}

// SetCronTime by cron like format schedule
func (e *Exec) SetCronTime(m string, h string, d string, mon string, w string) {
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

// GetLogName get log name
func (e *Exec) GetLogName() string { return e.Name + "_" + e.GetNameID8b() + ".log" }

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
	logName := e.GetLogName()
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
		log.Fatalln("no time")
	}

	// start cron
	e.Cron = cron.New()
	if _, err := e.Cron.AddFunc(e.time, func() {e.DoExec()}); err != nil {
		log.Fatalln(err)
	}
	log.Println(e.Name + " cron start!")
	e.Cron.Start()
}

// StopCron to stop job
func (e *Exec) StopCron() {
	e.Cron.Stop()
	e.done = false
	log.Println(e.Name + " cron has stopped!")
}

// DeleteLog to delete log
func (e *Exec) DeleteLog() error {
	logName := e.GetLogName()
	return os.RemoveAll(logName)
}

// DoExecute execute command and log recording
// phrase command to args[] first
// exec comand
// output log
// TODO: set default system log path
func DoExecute(logName string, command string) {
	args := strings.Fields(strings.TrimSpace(command))
	cmd := exec.Command(args[0], args[1:]...)
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)

	log.Println("do execute")
	f, err := os.OpenFile(
		logName,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	out, err := cmd.CombinedOutput()
	if err == nil {
		log.Println(command)
		log.Println("done")
	}
	log.Printf("The output:\n\n%s\n", out)
}
