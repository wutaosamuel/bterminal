package html

/*
 * Handle index.html action
 */

import (
	"encoding/base64"
	"fmt"
	"net/http"
)

// HandleIndex is process functions of index.html
// TODO:
// Config
func (c *ConfigHTML) HandleIndex(w http.ResponseWriter, req *http.Request) {
	//http.ServeFile(w, req, "html/")
	req.ParseForm()
	PrintHTMLInfo(req)

	// First time access server
	// check user if is login
	// if not, serve html/index.html
	if !c.isLogIn(w, req) {
		http.ServeFile(w, req, "html/index.html")
		c.setSession(w)
	}

	// Form will require by POST
	if req.Method == "POST" {
		password, _ := base64.StdEncoding.DecodeString(FormToString(req, "password"))
		fmt.Println(FormToString(req, "password"))
		fmt.Println(string(password))
		// TODO:
		// Read Config and allow to sign in
		// Cookie for Auth
		http.Redirect(w, req, "/shell.html", 307)
	}
}
