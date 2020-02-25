package html

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"../conf"
	"../job"
	"../utils"
)

// ConfigHTML for HTML use
type ConfigHTML struct {
	*sync.RWMutex      // read & write locker for execs
	*utils.CookieUtils // store session and token in cookie

	Config *conf.Config        // local config process
	//JobID		string  // id for each job
	Jobs  map[string]job.Exec // job execution
}

// NewConfigHTML create new one
func NewConfigHTML(defaultExpiration time.Duration) *ConfigHTML {
	return &ConfigHTML{
		&sync.RWMutex{},
		utils.NewCookie(defaultExpiration),
		conf.NewConfig(),
		make(map[string]job.Exec)}
}

// ConfigHTMLInit init ConfigHTML
// DefaultExpiration is 6 hours
// every cookie is kept whin 6 hours
func (c *ConfigHTML) ConfigHTMLInit() *ConfigHTML {
	c = NewConfigHTML(6 * time.Hour)
	return NewConfigHTML(6 * time.Hour)
}

/////////////////// Private ////////////////

// authentication check security
// not working on index page
func (c *ConfigHTML) authentication(w http.ResponseWriter, req *http.Request, html string) bool {
	if !c.isLogIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return false
	}
	c.setToken(w)
	http.ServeFile(w, req, html)
	return true
}

// isLogIn check whether user has login in
// Check cookie has session ID
// if not, return -> redirect to index page for login
// or -> serve html page
func (c *ConfigHTML) isLogIn(req *http.Request) bool {
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

// isToken check token form cookie
func (c *ConfigHTML) isToken(w http.ResponseWriter, req *http.Request) bool {
	// avoiding multiple cookie with same name
	for _, cookie := range req.Cookies() {
		if cookie.Name == utils.CookieToken {
			if c.IsToken(cookie.Value) {
				// delete from cookie
				tokenCookie := &http.Cookie{
					Name:     utils.CookieToken,
					Value:    cookie.Value,
					MaxAge:   -1,
					HttpOnly: true}
				http.SetCookie(w, tokenCookie)
				return true
			}
		}
	}
	return false
}

// setToken allows user single submit form
func (c *ConfigHTML) setToken(w http.ResponseWriter) {
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
func (c *ConfigHTML) setSession(w http.ResponseWriter) {
	// generate and store a session
	session := c.SetSession()
	sessionCookie := &http.Cookie{
		Name:     utils.CookieSession,
		Value:    session,
		MaxAge:   43200, // 12 hours
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
