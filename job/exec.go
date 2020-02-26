package job

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"

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
	LogPath string // log path

	*sync.RWMutex             // Read & write lock
	Logger        *log.Logger // logger for exec

	Cron *cron.Cron // does job need to schedule
	Time string     // schedule time of job
}

/////////////////// Setter&&Getter ///////////////////

// NewExecS create a new Exec struct
// Init required after create a new Exec
func NewExecS() *Exec {
	return &Exec{}
}

// NewExec create a new Exec struct
func NewExec() *Exec {
	uuid := uuid.Must(uuid.NewV4()).String()
	return &Exec{nameID: uuid}
}

// Init init exec
func (e *Exec) Init() error {
	e.nameID = uuid.Must(uuid.NewV4()).String()
	// set log path
	e.SetLogPath()
	return nil
}

// SetCronTime by cron like format schedule
func (e *Exec) SetCronTime(m string, h string, d string, mon string, w string) {
	e.Time = m + " " + h + " " + d + " " + mon + " " + w
}

// SetLogPath set log path
func (e *Exec) SetLogPath() {
	// FIXME: path not exist
	if len(e.LogPath) < 4 {
		e.LogPath = path.Join(e.LogPath, e.GetLogName())
		return
	}
	if e.LogPath[len(e.LogPath)-4:] != ".log" {
		e.LogPath = path.Join(e.LogPath, e.GetLogName())
		return
	}
	return
}

// GetNameID get uuid for job
func (e *Exec) GetNameID() string { return e.nameID }

// GetNameID8b get the first 8 bits uuid string
func (e *Exec) GetNameID8b() string { return e.nameID[:8] }

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
	e.Lock()
	DoExecute(logName, e.Command)
	e.Unlock()
}

// StartCron do schedule of Exec
func (e *Exec) StartCron() {
	// start cron
	e.Lock()
	e.Cron = cron.New()
	if _, err := e.Cron.AddFunc(e.Time, func() { e.DoExec() }); err != nil {
		e.WriteLogFunc(func(l *log.Logger){l.Fatalln(err)})
	}
	e.WriteLog(e.Name + " cron start!")
	e.Cron.Start()
	e.Unlock()
}

// StopCron to stop job
func (e *Exec) StopCron() {
	e.Lock()
	e.Cron.Stop()
	e.Unlock()
	e.WriteLog(e.Name + " cron has stopped!")
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
