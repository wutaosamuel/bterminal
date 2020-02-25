package html

/*
 * Handle jobs action for jobs.html
 * Generate jobs.html first
 */

import (
	"fmt"
	"net/http"

	"../utils"
)

// TODO: Restore Jobs

// Job for job html
type Job struct {
	Name    string // job name
	ID      string // job UUID
	Command string // Command required to run
	Crontab string // cron schedule
	Stop    string // button name
}

// NewJob create new job
func NewJob() *Job {
	return &Job{}
}

// SetID set id and button
func (j *Job) SetID(i string) {
	j.ID = i
	j.Stop = "Stop-"+i
}

// GenerateJobs automatically generate
func GenerateJobs(jobs []Job, template, pattern string) string {
	// if num of jobs is 0,
	// replace {{{ 1 }}} and output template only
	if len(jobs) == 0 {
		html, _ := utils.ReplaceHTML(template, 1, "")
		return html
	}

	// process pattern first
	var p string
	for _, job := range jobs {
		tmp, _ := utils.ReplacePattern(pattern, job)
		p += tmp
	}

	// replease job html
	html, _ := utils.ReplaceHTML(template, 1, p)
	return html
}

// HandleJobs handle jobs
func (c *ConfigHTML) HandleJobs(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	PrintHTMLInfo(req)

	// authentication is login
	if !c.authentication(w, req, "/html/jobs.html") {
		return
	}

	// Read form
	if req.Method == "POST" {
		PrintHTMLInfo(req)
		fmt.Println("need job action")
	}
}