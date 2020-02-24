package html

import (
	"fmt"
	"net/http"
	"strings"

	"../utils"
)

// CookieHTML for HTML use
type CookieHTML struct {
	*utils.CookieUtils
}

// authentication check security
// not working on index page
func (c *CookieHTML) authentication(w http.ResponseWriter, req *http.Request, html string) {
	if !c.isLogIn(w, req) {
		http.Redirect(w, req, "/", http.StatusFound)
		return
	}
	c.setToken(w)
	http.ServeFile(w, req, html)
	return
}

// isLogIn check whether user has login in
// Check cookie has session ID
// if not, return -> redirect to index page for login
// or -> serve html page
func (c *CookieHTML) isLogIn(w http.ResponseWriter, req *http.Request) bool {
	// avoiding multiple cookies with same name
	for _, cookie := range req.Cookies() {
		if cookie.Name == utils.CookieSession {
			if c.IsSession(cookie.Value) {
				return true
			}
		}
	}
	return false
}

// setToken allows user single submit form
func (c *CookieHTML) setToken(w http.ResponseWriter) {
	// generate and store a token
	token := c.SetToken()
	tokenCookie := &http.Cookie{
		Name:     utils.CookieToken,
		Value:    token,
		MaxAge:   10800, // 3 hours
		HttpOnly: true}
	http.SetCookie(w, tokenCookie)
}

// setSession set session to user's browser
func (c *CookieHTML) setSession(w http.ResponseWriter) {
	// generate and store a session
	session := c.SetSession()
	sessionCookie := &http.Cookie{
		Name: utils.CookieSession,
		Value: session,
		MaxAge: 43200,	// 12 hours
		HttpOnly: true}
  http.SetCookie(w, sessionCookie)
}

// PrintHTMLInfo infomation
func PrintHTMLInfo(req *http.Request) {
	fmt.Println(req.Form)
	fmt.Println("path: ", req.URL.Path)
	fmt.Println("scheme: ", req.URL.Scheme)
	fmt.Println("method: ", req.Method)
}

// FormToString to string
func FormToString(req *http.Request, attribute string) string {
	return strings.Join(req.Form[attribute], "")
}