package main

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/pflag"

	"github.com/wutaosamuel/bterminal/conf"
	"github.com/wutaosamuel/bterminal/utils"
	bt "github.com/wutaosamuel/bterminal"
)

func main() {
	// read from cli
	var (
		helpFlag       = pflag.BoolP("help", "h", false, "display usage")
		configPathFlag = pflag.StringP("config", "c", "/etc/bterminal/config.json", "Json config file")
		portFlag       = pflag.IntP("port", "p", 5122, "TCP port for web service")
		passwordFlag   = pflag.StringP("password", "P", "", "password for protecting web service")
		logDirFlag     = pflag.StringP("log", "l", "", "log directory for keeping job logs")
	)

	pflag.Parse()

	if *helpFlag {
		pflag.Usage()
		os.Exit(0)
	}

	// set config
	config := conf.NewConfig()
	isFile, _ := utils.IsFile(*configPathFlag)
	if isFile {
		config.Path = *configPathFlag
	}
	if !isFile {
		config.Path = ""
	}
	config.Port = strconv.Itoa(*portFlag)
	config.Password = *passwordFlag
	if *logDirFlag != "" {
		config.LogDir = *logDirFlag
	}
	if *logDirFlag == "" {
		config.LogDir = filepath.Join("/var", "log", "bterminal")
	}
	config.Init()

	// set App path
	appPath := "/usr/share/bterminal"

	// do main func
	bt.Main(config, appPath)
}
