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
func (c *ConfigHTML) HandleIndex(w http.ResponseWriter, req *http.Request) {
	//http.ServeFile(w, req, "html/")
	req.ParseForm()
	PrintHTMLInfo(req)

	// First time access server
	// check user if is login
	// redirect to shell.html
	// if not, serve html/index.html
	if c.isLogIn(req) || c.Config.Password == "" {
		http.Redirect(w, req, "/shell.html", http.StatusFound)
		return
	}
	c.setToken(w)
	http.ServeFile(w, req, "html/index.html")

	// Form will require by POST
	// check token, only allow to submit form once
	if req.Method == "POST" {
		if !c.isToken(w, req) {
			return
		}
		password, _ := base64.StdEncoding.DecodeString(FormToString(req, "password"))
		fmt.Println(string(password))
		if string(password) != c.Config.Password {
			// TODO: display password error info
			http.Redirect(w, req, "/shell.html", http.StatusFound)
		}
		// set session to user
		c.setSession(w)
		http.Redirect(w, req, "/shell.html", http.StatusFound)
	}
}
