package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/wutaosamuel/bterminal/conf"
	ht "github.com/wutaosamuel/bterminal/html"
	"github.com/wutaosamuel/bterminal/job"
)

// Main func for bterminal
// TODO: log dir path should depend on system
func Main(config *conf.Config, appPath string) {
	// set default(cmd prefer), when there is no value
	// if path is nil, no file to load
	// if port is nil, port is 5122
	// if password is nil, no password needed
	// if LogDir is nil, save log into local dir
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
			logDir := filepath.Join(appPath, "log")
			if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
				panic(err)
			}
			config.LogDir = logDir
		}
	}

	// setting up config html
	configHTML := ht.NewConfigHTML(6 * time.Hour)
	configHTML.Config = config
	configHTML.AppPath = appPath

	// recover jobs from dat
	datPath := filepath.Join(appPath, "GobData.dat")
	dat := job.NewDat()
	err := dat.ReadDecode(datPath)
	if err != nil {
		panic(err)
	}
	// set recover job in config html
	configHTML.RecoverDat(dat)

	//generate jobs.html & logs.html
	fmt.Println("start html")
	configHTML.Start()

	// start http/web server
	// set static assert
	fmt.Println("start web service")
	http.Handle("/CSS/", http.StripPrefix("/CSS/", http.FileServer(http.Dir(filepath.Join(appPath, "html", "CSS")))))
	http.Handle("/favicon.ico", http.FileServer(http.Dir(filepath.Join(appPath, "html", "image", "favicon.png"))))
	// hand func
	http.HandleFunc("/", configHTML.HandleIndex)
	http.HandleFunc("/shell.html", configHTML.HandleShell)
	http.HandleFunc("/jobs.html", configHTML.HandleJobs)
	http.HandleFunc("/logs.html", configHTML.HandleLogs)
	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		panic(err)
	}
}
