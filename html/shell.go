package html

/*
 * Handle function of shell.html
 */

import (
	"fmt"
	"net/http"
)

// HandleShell shell.html func
// do not care value pass by GET or POST
// Check client cookie first
// TODO:
// Cookie
// Name & ID & Command & Crontab
func HandleShell(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	PrintHTMLInfo(req)
	// TODO: check client cookie here

	// Get shell.html page
	if req.Method == "GET" {
		fmt.Println("handle shell method get")
		http.ServeFile(w, req, "html/shell.html")
	}

	// Read form
	// TODO: do exec here
	if req.Method == "POST" {
		fmt.Println(req.Form["command"])
		fmt.Println(req.Form["crontab"])
	}
}
