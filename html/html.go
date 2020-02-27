package html

import "net/http"

// StartHTTP start http setting
func StartHTTP(c *ConfigHTML) {
	// set static assert
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir("./html/CSS"))))
	// hand func
	http.HandleFunc("/", c.HandleIndex)
	http.HandleFunc("/shell.html", c.HandleShell)
	http.HandleFunc("/jobs.html", c.HandleJobs)
	http.HandleFunc("/logs.html", c.HandleLogs)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}
}