package html

/*
 * Handle jobs action for jobs.html
 * Generate jobs.html first
 */

import (
	"net/http"
	"path/filepath"

	// FIXME:
	"github.com/wutaosamuel/bterminal/utils"
	// "../utils"
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
// TODO: job restart action
func (c *ConfigHTML) HandleJobs(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()

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
		c.jobsAction(w, req)
	}
	return
}

// JobsAction do jobs action
// general stop action
// TODO: job restart action
func (c *ConfigHTML) jobsAction(w http.ResponseWriter, req *http.Request) {
	// read ID for stop
	for key := range req.Form {
		if key[:5] == "Stop-" {
			c.Lock()
			job := c.Jobs[key[5:]]
			j := c.setJob(&job)
			job.StopCron()
			// delete job from jobs.html
			// TODO: restart job action
			err := utils.DeletePage(
				j,
				filepath.Join(c.AppPath, "html", "jobs.html"),
				filepath.Join(c.AppPath, "html", "pattern", "job_pattern1.html"))
			if err != nil {
				job.WriteLog(err)
			}
			// set cron time is ""
			job.Time = "stopped" + job.Time
			c.Jobs[key[5:]] = job
			c.Unlock()
			http.Redirect(w, req, "/jobs.html", http.StatusSeeOther)
			return
		}
	}
}
