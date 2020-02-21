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
func HandleJobs(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	PrintHTMLInfo(req)
	// TODO: check client cookie here

	// Get shell.html page
	if req.Method == "GET" {
		fmt.Println("handle Jobs method get")
		http.ServeFile(w, req, "html/jobs.html")
	}

	// Read form
	// TODO: do exec here
	if req.Method == "POST" {
		PrintHTMLInfo(req)
		fmt.Println("need job action")
	}
}