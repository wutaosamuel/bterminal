package html

/*
 * Handle function of shell.html
 */

import (
	"fmt"
	"net/http"

	"../job"
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
	// TODO: 
	// do exec then genrate log
	// do cron then genrate jon & log
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
			http.Redirect(w, req, "/html/shell.html", http.StatusNotModified)
			return
		}
	}
}

/////////////////// Private /////////////////

// setExec set log struct
// config must be contain
func (c *ConfigHTML) setExec(name, command, crontab string) {
	c.Lock()
	e := job.NewExecS()
	e.Name = name
	e.Command = command
	e.LogPath = c.Config.LogPath
	e.Time = crontab
}