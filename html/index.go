package html

/*
 * Handle index.html action
 */

import (
	"fmt"
	"net/http"
)

// HandleIndex is process functions of index.html
// TODO:
// Cookie
// Config
func HandleIndex(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Println(req.Form)
	fmt.Println("path ", req.URL.Path)
	fmt.Println("scheme ", req.URL.Scheme)

	// not all value pass by GET URL
	if req.Method == "GET" {
		// FIXME: potential fail
		http.ServeFile(w, req, "./index.html")
	}

	fmt.Println(req.Form["password"])
	http.ServeFile(w, req, "./jobs.html")
}
