package html

/*
 * Handle logs action for jobs.html
 * Generate logs.html first
 */

import (
	"fmt"
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

// SetID set id for detail and delete
func (l *JobLog) SetID(i string) {
	l.ID = i
	l.Delete = "Delete-" + i 
	l.Detail = "Detail-" + i
}

// GenerateJobLogs automatically generate
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
func HandleLogs(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	PrintHTMLInfo(req)
	// TODO: check client cookie here

	// Get shell.html page
	if req.Method == "GET" {
		fmt.Println("handle Logs method get")
		http.ServeFile(w, req, "html/logs.html")
	}

	// Read form
	// TODO: do exec here
	if req.Method == "POST" {
		fmt.Println("need log action")
	}
}
