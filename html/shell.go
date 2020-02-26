package html

/*
 * Handle function of shell.html
 */

import (
	"net/http"

	"../utils"
)

// HandleShell shell.html func
// do not care value pass by GET or POST
// Check client cookie first
func (c *ConfigHTML) HandleShell(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	PrintHTMLInfo(req)

	// authentication is login
	if !c.authentication(w, req, "/html/shell.html") {
		return
	}

	// Read form
	if req.Method == "POST" {
		if !c.isToken(w, req) {
			return
		}
		name := FormToString(req, "name")
		command := FormToString(req, "command")
		crontab := FormToString(req, "crontab")
		// command needed
		if command == "" {
			// TODO: display info
			http.Redirect(w, req, "/html/shell.html", http.StatusNotModified)
			return
		}
		// do exec then generate log html
		// TODO: display info
		if crontab == "" {
			c.execAction(name, command, crontab)
		}
		// do cron then generate job & log
		if crontab != "" {
			c.cronAction(name, command, crontab)
		}
	}
}

// execAction action for executing a job
func (c *ConfigHTML) execAction(name, command, crontab string) {
	e := c.setExec(name, command, crontab)
	l := c.setJobLog(e)
	e.DoExec()
	c.Lock()
	c.JobID = append(c.JobID, e.GetNameID())
	utils.UpdatePage(l, "./html/logs.html", "./html/pattern/log_pattern1.html")
	c.Unlock()
}

// cronAction action for a crontab job
func (c *ConfigHTML) cronAction(name, command, crontab string) {
	e := c.setExec(name, command, crontab)
	j := c.setJob(e)
	l := c.setJobLog(e)
	c.Lock()
	c.JobID = append(c.JobID, e.GetNameID())
	c.CronJobs[e.GetNameID()] = *e
	utils.UpdatePage(l, "./html/logs.html", "./html/pattern/log_pattern1.html")
	utils.UpdatePage(j, "./html/jobs.html", "./html/pattern/job_pattern1.html")
	c.Unlock()
}