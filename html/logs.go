package html

/*
 * Handle logs action for jobs.html
 * Generate logs.html first
 */

import (
	"fmt"
	"net/http"
	"path/filepath"

	"../utils"
)

// JobLog for logs html
type JobLog struct {
	Name    string // job name
	ID      string // job UUID
	Command string // Command required to run
	Crontab string // cron schedule
	Detail  string // button for detail
	Delete  string // button for delete
}

// NewJobLog create new job
func NewJobLog() *JobLog {
	return &JobLog{}
}

// Init init joblog
func (l *JobLog) Init() {
	l.SetID(l.ID)
}

// SetID set id for detail and delete
func (l *JobLog) SetID(i string) {
	l.ID = i
	l.Delete = "Delete-" + i
	l.Detail = "Detail-" + i
}

// GenerateJobLogs automatically generate
// if logs.html is not at html directory
// or force replace logs.html
func GenerateJobLogs(logs []JobLog, template, pattern string) string {
	templateS, err := utils.ReadHTML(template)
	utils.CheckPanic(err)
	patternS, err := utils.ReadHTML(pattern)
	utils.CheckPanic(err)
	// if num of jobs is 0,
	// replace {{{ 1 }}} and output template only
	if len(logs) == 0 {
		html, _ := utils.ReplaceHTML(templateS, 1, "")
		return html
	}

	// process pattern first
	var p string
	for _, l := range logs {
		tmp, _ := utils.ReplacePattern(patternS, l)
		p += tmp
	}

	// replease job html
	html, _ := utils.ReplaceHTML(templateS, 1, p)
	return html
}

// HandleLogs handle logs.html action
func (c *ConfigHTML) HandleLogs(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	PrintHTMLInfo(req)

	// authentication is login
	if req.Method == "GET" {
		if !c.authentication(w, req, "logs.html") {
			return
		}
	}

	// Read form
	if req.Method == "POST" {
		if !c.isToken(w, req) {
			return
		}
		for key := range req.Form {
			c.deleteLog(key)
			c.logDetail(w, key)
		}
	}
	return
}

// logDetail read a log
func (c *ConfigHTML) logDetail(w http.ResponseWriter, key string) {
	if key[:7] == "Detail-" {
		j := c.Jobs[key[7:]]
		detail, err := c.setLogDetail(key[7:])
		if err != nil {
			j.WriteLog(j)
		}
		c.RLock()
		html, err := detail.GenerateDetail(
			filepath.Join(c.AppPath, "html", "template", "detail.html"),
			filepath.Join(c.AppPath, "html", "pattern", "detail_pattern1.html"))
		if err != nil {
			j.WriteLog(j)
		}
		fmt.Fprintf(w, html)
		c.RUnlock()
	}
}

// deleteLog delete a log
func (c *ConfigHTML) deleteLog(key string) {
	if key[:7] == "Delete-" {
		c.Lock()
		j := c.Jobs[key[7:]]
		jobLog := c.setJobLog(&j)
		if err := j.DeleteLog(); err != nil {
			j.WriteLog(err)
		}
		err := utils.DeletePage(
			jobLog,
			filepath.Join(c.AppPath, "html", "logs.html"),
			filepath.Join(c.AppPath, "html", "pattern", "log_pattern1.html"))
		if err != nil {
			j.WriteLog(err)
		}
		delete(c.JobID, j.GetNameID())
		delete(c.Jobs, j.GetNameID())
		c.Unlock()
	}
}
