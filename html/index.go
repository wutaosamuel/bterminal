package html

/*
 * Handle index.html action
 */

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"path/filepath"
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
	if req.Method == "GET" {
		if c.Config.Password == "" {
			c.setSession(w)
			http.Redirect(w, req, "/shell.html", http.StatusFound)
			return
		}
		if c.isLogIn(req) {
			http.Redirect(w, req, "/shell.html", http.StatusFound)
			return
		}
		c.setToken(w)
		http.ServeFile(w, req, filepath.Join(c.AppPath, "html", "index.html"))
	}

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
			http.Redirect(w, req, "/", http.StatusFound)
			return
		}
		// set session to user
		c.setSession(w)
		http.Redirect(w, req, "/shell.html", http.StatusFound)
	}
	return
}
