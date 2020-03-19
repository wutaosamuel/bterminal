package html

import (
	"bufio"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/wutaosamuel/bterminal/conf"
	"github.com/wutaosamuel/bterminal/job"
	"github.com/wutaosamuel/bterminal/utils"
)

// ConfigHTML for HTML use
// TODO:
// Improve save Jobs as file for restore jobs
// Improve RW locker
type ConfigHTML struct {
	*sync.RWMutex      // read & write locker for execs
	*utils.CookieUtils // store session and token in cookie

	AppPath string              // set AppPath for read html file in ./bterminal/html
	Config  *conf.Config        // local config process
	Jobs    map[string]job.Exec // keep cron jobs TODO: merge DataStorage
}

// NewConfigHTML create new one
func NewConfigHTML(defaultExpiration time.Duration) *ConfigHTML {
	return &ConfigHTML{
		&sync.RWMutex{},
		utils.NewCookie(defaultExpiration),
		"",
		conf.NewConfig(),
		make(map[string]job.Exec),
	}
}

// ConfigHTMLInit init ConfigHTML
// DefaultExpiration is 6 hours
// every cookie is kept whin 6 hours
func (c *ConfigHTML) ConfigHTMLInit() *ConfigHTML {
	c = NewConfigHTML(6 * time.Hour)
	return NewConfigHTML(6 * time.Hour)
}

/////////////////// Public ////////////////

// Start setting up html
// recover jobs && logs
// generate jobs.html & logs.html
func (c *ConfigHTML) Start() {
	var (
		jobs    []Job
		jobLogs []JobLog
	)
	// read log dir
	logFiles, err := filepath.Glob(filepath.Join(c.Config.LogDir, "*"))
	utils.CheckPanic(err)

	// recover jobs & logs in terms of c.jobs, should be configured before starting
	// if time is "", it has been executed. not allow to recover jobs
	// if log not exist, not allow to recover joblogs
	for _, e := range c.Jobs {
		// add cron jobs
		if e.Time != "" {
			jobs = append(jobs, *c.setJob(&e))
			go e.StartCron()
		}
		// add jog logs
		for _, logFile := range logFiles {
			if strings.Contains(logFile, e.GetNameID()) {
				jobLogs = append(jobLogs, *c.setJobLog(&e))
				break
			}
		}
	}

	// generate jobs.html
	jobsHTML := GenerateJobs(
		jobs,
		filepath.Join(c.AppPath, "html", "template", "jobs.html"),
		filepath.Join(c.AppPath, "html", "pattern", "job_pattern1.html"))
	err = utils.SaveHTML(
		filepath.Join(c.AppPath, "html", "jobs.html"),
		jobsHTML)
	utils.CheckPanic(err)
	// generate jobLogs.html
	jobLogsHTML := GenerateJobLogs(
		jobLogs,
		filepath.Join(c.AppPath, "html", "template", "logs.html"),
		filepath.Join(c.AppPath, "html", "pattern", "log_pattern1.html"))
	err = utils.SaveHTML(
		filepath.Join(c.AppPath, "html", "logs.html"),
		jobLogsHTML)
	utils.CheckPanic(err)
}

// RecoverDat recover exec from dat
// more recover jobs & html on Start()
func (c *ConfigHTML) RecoverDat(d *job.Dat) {
	e := make(map[string]job.Exec)
	for id, j := range d.Jobs {
		tmp := job.NewExecS()
		tmp.Name = j.Name
		tmp.NameID = id
		tmp.Command = j.Command
		tmp.LogName = j.LogName
		tmp.Time = j.Time
		e[id] = *tmp
	}
	c.Jobs = e
}

// FormToString to string
func FormToString(req *http.Request, attribute string) string {
	return strings.Join(req.Form[attribute], "")
}

/////////////////// Private ////////////////

// authentication check security
// not working on index page
func (c *ConfigHTML) authentication(w http.ResponseWriter, req *http.Request, name string) bool {
	if !c.isLogIn(req) {
		// do login
		if c.Config.Password == "" {
			c.setSession(w)
			redirectLink := "/" + name
			http.Redirect(w, req, redirectLink, http.StatusFound)
			return true
		}
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return false
	}
	c.setToken(w)
	http.ServeFile(w, req, filepath.Join(c.AppPath, "html", name))
	return true
}

// isLogIn check whether user has login in
// Check cookie has session ID
// if not, return -> redirect to index page for login
// or -> serve html page
func (c *ConfigHTML) isLogIn(req *http.Request) bool {
	// avoiding multiple cookies with same name
	for _, cookie := range req.Cookies() {
		if cookie.Name == utils.CookieSession {
			if c.IsSession(cookie.Value) {
				return true
			}
		}
	}
	return false
}

// isToken check token form cookie
func (c *ConfigHTML) isToken(w http.ResponseWriter, req *http.Request) bool {
	// avoiding multiple cookie with same name
	for _, cookie := range req.Cookies() {
		if cookie.Name == utils.CookieToken {
			if c.IsToken(cookie.Value) {
				// delete from cookie
				tokenCookie := &http.Cookie{
					Name:     utils.CookieToken,
					Value:    cookie.Value,
					MaxAge:   -1,
					HttpOnly: true}
				http.SetCookie(w, tokenCookie)
				return true
			}
		}
	}
	return false
}

// setToken allows user single submit form
func (c *ConfigHTML) setToken(w http.ResponseWriter) {
	// generate and store a token
	token := c.SetToken()
	tokenCookie := &http.Cookie{
		Name:     utils.CookieToken,
		Value:    token,
		MaxAge:   10800, // 3 hours
		HttpOnly: true}
	http.SetCookie(w, tokenCookie)
}

// setSession set session to user's browser
func (c *ConfigHTML) setSession(w http.ResponseWriter) {
	// generate and store a session
	session := c.SetSession()
	sessionCookie := &http.Cookie{
		Name:     utils.CookieSession,
		Value:    session,
		MaxAge:   43200, // 12 hours
		HttpOnly: true}
	http.SetCookie(w, sessionCookie)
}

// setExec set log struct
// config must be contain
func (c *ConfigHTML) setExec(name, command, crontab string) *job.Exec {
	e := job.NewExecS()
	e.Name = name
	e.Command = command
	e.LogName = c.Config.LogDir
	e.Time = crontab
	e.Init()
	return e
}

// setJob set job for job.html
func (c *ConfigHTML) setJob(e *job.Exec) *Job {
	j := NewJob()
	j.Name = e.Name
	j.ID = e.GetNameID()
	j.Command = e.Command
	j.Crontab = e.Time
	j.Init()
	return j
}

// setJobLog set job log for logs.html
func (c *ConfigHTML) setJobLog(e *job.Exec) *JobLog {
	jobLog := NewJobLog()
	jobLog.Name = e.Name
	jobLog.ID = e.GetNameID()
	jobLog.Command = e.Command
	jobLog.Crontab = e.Time
	jobLog.Init()
	return jobLog
}

// setLogDetail display log detail for every job
func (c *ConfigHTML) setLogDetail(id string) (*Detail, error) {
	d := NewDetail()
	e := c.Jobs[id]
	d.Name = e.Name
	d.ID = e.GetNameID()
	d.Command = e.Command
	d.Crontab = e.Time
	isLog, err := utils.IsFile(e.LogName)
	if err != nil {
		return d, err
	}
	var logString string
	if isLog {
		logF, err := os.Open(e.LogName)
		defer logF.Close()
		if err != nil {
			return d, err
		}
		logScanner := bufio.NewScanner(logF)
		// TODO: improve display
		for logScanner.Scan() {
			logString += "<p>" + logScanner.Text() + "</p>"
		}
	}
	if !isLog {
		logString = ""
	}
	d.Log = logString
	return d, nil
}

// updateDat update job data
func (c *ConfigHTML) updateDat() error {
	datPath := filepath.Join(c.AppPath, "GobData.dat")
	dat := job.NewDat()
	c.Lock()
	dat.SetDatS(c.Jobs)
	err := dat.SaveEncode(datPath)
	c.Unlock()
	if err != nil {
		return err
	}
	return nil
}
