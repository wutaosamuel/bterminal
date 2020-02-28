package bterminal

import (
	"os"
	"fmt"
	"time"
	"net/http"
	"path/filepath"

	"./conf"
	ht "./html"
)

// Main func for bterminal
func Main(config *conf.Config, appPath string) {
	// set default(cmd prefer), when there is no value
	// if path is nil, no file to load
	// if port is nil, port is 5122
	// if password is nil, no password needed
	// if LogDir is nil, save log into local dir
	// TODO: log dir path should depend on system
	fmt.Println("start")
	c := conf.NewConfig()
	if config.Path != "" {
		c.Load(config.Path)
	}
	if config.Port == "" {
		if c.Port != "" {
			config.Port = c.Port
		}
		if c.Port == "" {
			config.Port = "5122"
		}
	}
	if config.Password == "" {
		if c.Password != "" {
			config.Password = c.Password
		}
	}
	if config.LogDir == "" {
		if c.LogDir != "" {
			config.LogDir = c.LogDir
		}
		if c.LogDir == "" {
			if err := os.MkdirAll("log", os.ModePerm); err != nil {
				panic(err)
			}
			logDir, err := filepath.Abs("log")
			if err != nil {
				panic(err)
			}
			config.LogDir = logDir
		}
	}
	fmt.Println(config)

	// setting up config html
	fmt.Println("start html")
	configHTML := ht.NewConfigHTML(6 * time.Hour)
	configHTML.Config = config
	configHTML.AppPath = appPath
	// TODO: restore jobs
	//configHTML.JobID
	//configHTML.Jobs

	//generate jobs.html & logs.html
	configHTML.Start()

	// start http/web server
	// set static assert
	fmt.Println("start web service")
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir(filepath.Join(appPath, "html", "CSS")))))
	// hand func
	http.HandleFunc("/", configHTML.HandleIndex)
	http.HandleFunc("/shell.html", configHTML.HandleShell)
	http.HandleFunc("/jobs.html", configHTML.HandleJobs)
	http.HandleFunc("/logs.html", configHTML.HandleLogs)
	err := http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		panic(err)
	}
}
