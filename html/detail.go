package html

/*
 * Handle detail action for jobs.html
 * Generate detail.html first
 */

import (
	"fmt"
	"net/http"

	"../utils"
)

// Detail for detail html
type Detail struct {
	Name    string // job name
	ID      string // job UUID
	Command string // Command required to run
	Crontab string // cron schedule
	Log     string // log
}

// NewDetail create new job
func NewDetail() *Detail {
	return &Detail{}
}

// GenerateDetail automatically generate
func GenerateDetail(details []Detail, template, pattern string) string {
	// if num of jobs is 0,
	// replace {{{ 1 }}} and output template only
	if len(details) == 0 {
		html, _ := utils.ReplaceHTML(template, 1, "")
		return html
	}

	// process pattern first
	var p string
	for _, detail := range details {
		tmp, _ := utils.ReplacePattern(pattern, detail)
		p += tmp
	}

	// replease job html
	html, _ := utils.ReplaceHTML(template, 1, p)
	return html
}

// HandleDetail handle detail action
func HandleDetail(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	PrintHTMLInfo(req)
	// TODO: check client cookie here

	// Get shell.html page
	if req.Method == "GET" {
		fmt.Println("handle Detail method get")
		http.ServeFile(w, req, "html/detail.html")
	}

	// Read form
	// TODO: do exec here
	if req.Method == "POST" {
		fmt.Println("need detail action")
	}
}
