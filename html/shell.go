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
	// TODO: check client cookie here

	// Read form
	// TODO: do exec here
	fmt.Println(req.Form["command"])
	fmt.Println(req.Form["crontab"])
}
