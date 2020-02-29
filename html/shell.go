package html

/*
 * Handle function of shell.html
 */

import (
	"log"
	"fmt"
	"net/http"
	"path/filepath"

	"../utils"
)

// HandleShell shell.html func
// do not care value pass by GET or POST
// Check client cookie first
func (c *ConfigHTML) HandleShell(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	PrintHTMLInfo(req)

	// authentication is login
	if req.Method == "GET" {
		if !c.authentication(w, req, "shell.html") {
			return
		}
	}

	// Read form
	if req.Method == "POST" {
		if !c.isToken(w, req) {
			return
		}
		name := FormToString(req, "name")
		command := FormToString(req, "command")
		crontab := FormToString(req, "crontab")
		fmt.Println(name)
		fmt.Println(command)
		fmt.Println(crontab)
		// command needed
		if command == "" {
			// TODO: display info
			http.Redirect(w, req, "/shell.html", http.StatusNotModified)
			return
		}
		// do exec then generate log html
		// TODO: display info
		if crontab == "" {
			c.execAction(name, command, crontab)
			fmt.Println("exec done")
			http.Redirect(w, req, "/logs.html", http.StatusSeeOther)
		}
		// do cron then generate job & log
		if crontab != "" {
			c.cronAction(name, command, crontab)
			http.Redirect(w, req, "/jobs.html", http.StatusSeeOther)
		}
	}
	return
}

// execAction action for executing a job
func (c *ConfigHTML) execAction(name, command, crontab string) {
	fmt.Println("start execAction")
	e := c.setExec(name, command, crontab)
	l := c.setJobLog(e)
	e.DoExec()
	c.Lock()
	c.JobID[e.GetNameID()] = 1
	c.Jobs[e.GetNameID()] = *e
	// update logs.html
	fmt.Println("append page")
	err := utils.AppendPage(
		&l,
		filepath.Join(c.AppPath, "html", "logs.html"),
		filepath.Join(c.AppPath, "html", "pattern", "log_pattern1.html"))
	if err != nil {
		log.Println(err)
	}
	c.Unlock()
	fmt.Println("done")
}

// cronAction action for a crontab job
func (c *ConfigHTML) cronAction(name, command, crontab string) {
	e := c.setExec(name, command, crontab)
	j := c.setJob(e)
	l := c.setJobLog(e)
	e.StartCron()
	c.Lock()
	c.JobID[e.GetNameID()] = 1
	c.Jobs[e.GetNameID()] = *e
	// update logs.html and jobs.html
	err := utils.AppendPage(
		l,
		filepath.Join(c.AppPath, "html", "logs.html"),
		filepath.Join(c.AppPath, "html", "pattern", "log_pattern1.html"))
	if err != nil {
		log.Println(err)
	}
	err = utils.AppendPage(
		j,
		filepath.Join(c.AppPath, "html", "jobs.html"),
		filepath.Join(c.AppPath, "html", "pattern", "job_pattern1.html"))
	if err != nil {
		log.Println(err)
	}
	c.Unlock()
}
