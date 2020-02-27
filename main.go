package bterminal

import (
	"os"
	"path/filepath"

	"./conf"
)

// Main func for bterminal
func Main(config *conf.Config) {
	// set default(cmd prefer), when there is no value
	// if path is nil, no file to load
	// if port is nil, port is 5122
	// if password is nil, no password needed
	// if LogDir is nil, save log into local dir
	// TODO: log dir path should depend on system
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

	//
}
