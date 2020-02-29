package html

/*
 * Handle jobs action for jobs.html
 * Generate jobs.html first
 */

import (
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

// Init init job
func (j *Job) Init() {
	j.SetID(j.ID)
}

// SetID set id and button
func (j *Job) SetID(i string) {
	j.ID = i
	j.Stop = "Stop-" + i
}

// GenerateJobs automatically generate
// if jobs.html is not at html directory
// or force replace jobs.html
func GenerateJobs(jobs []Job, template, pattern string) string {
	templateS, err := utils.ReadHTML(template)
	utils.CheckPanic(err)
	patternS, err := utils.ReadHTML(pattern)
	utils.CheckPanic(err)
	// if num of jobs is 0,
	// replace {{{ 1 }}} and output template only
	if len(jobs) == 0 {
		html, _ := utils.ReplaceHTML(templateS, 1, "")
		return html
	}

	// process pattern first
	var p string
	for _, job := range jobs {
		tmp, _ := utils.ReplacePattern(patternS, job)
		p += tmp
	}

	// replease job html
	html, _ := utils.ReplaceHTML(templateS, 1, p)
	return html
}

// HandleJobs handle jobs
func (c *ConfigHTML) HandleJobs(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	PrintHTMLInfo(req)

	// authentication is login
	if req.Method == "GET" {
		if !c.authentication(w, req, "jobs.html") {
			return
		}
	}

	// Read form
	// stop a cron job
	if req.Method == "POST" {
		if !c.isToken(w, req) {
			return
		}
		c.jobsAction(req)
	}
	return
}

// JobsAction do jobs action
// general stop action
func (c *ConfigHTML) jobsAction(req *http.Request) {
	// read ID for stop
	for key := range req.Form {
		if key[:5] == "Stop-" {
			c.Lock()
			job := c.Jobs[key[5:]]
			job.StopCron()
			c.Unlock()
		}
	}
}
