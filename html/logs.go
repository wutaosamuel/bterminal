package html

/*
 * Handle logs action for jobs.html
 * Generate logs.html first
 */

import (
	"fmt"
	"log"
	"net/http"

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
	// if num of jobs is 0,
	// replace {{{ 1 }}} and output template only
	if len(logs) == 0 {
		html, _ := utils.ReplaceHTML(template, 1, "")
		return html
	}

	// process pattern first
	var p string
	for _, l := range logs {
		tmp, _ := utils.ReplacePattern(pattern, l)
		p += tmp
	}

	// replease job html
	html, _ := utils.ReplaceHTML(template, 1, p)
	return html
}

// HandleLogs handle logs.html action
func (c *ConfigHTML) HandleLogs(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	PrintHTMLInfo(req)

	// authentication is login
	if !c.authentication(w, req, "/html/logs.html") {
		return
	}

	// Read form
	// TODO:
	// Generate detail page here
	// Or delete log
	if req.Method == "POST" {
		if !c.isToken(w, req) {
			return
		}
		for key := range req.Form {
		}
	}
}

// logDetail read a log
func (c *ConfigHTML) logDetail(key string) {
	detail := c.setLogDetail(key)

}

// deleteLog delete a log
func (c *ConfigHTML) deleteLog(key string) {
	if key[:7] == "Delete-" {
		c.Lock()
		job := c.Jobs[key[7:]]
		err := job.DeleteLog()
		// TODO: need to improve delete into html
		log.Println(err)
		c.Unlock()
	}
}