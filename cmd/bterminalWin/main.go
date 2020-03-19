package main

import (
	"os"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/spf13/pflag"

	"github.com/wutaosamuel/bterminal/conf"
	bt "github.com/wutaosamuel/bterminal"
)

// main func
func main() {
	// read from cli
	var (
		helpFlag       = pflag.BoolP("help", "h", false, "display usage")
		configPathFlag = pflag.StringP("config", "c", "", "Json config file")
		portFlag       = pflag.IntP("port", "p", 5122, "TCP port for web service")
		passwordFlag   = pflag.StringP("password", "P", "", "password for protecting web service")
		logDirFlag     = pflag.StringP("log", "l", "", "log directory for keeping job logs")
	)

	pflag.Parse()

	if *helpFlag {
		pflag.Usage()
		os.Exit(0)
	}

	appPath := filepath.Join("C:\\ProgramData", "bterminal")
	fmt.Println(appPath)

	// set config
	config := conf.NewConfig()
	if *configPathFlag == "" {
		config.Path = filepath.Join(appPath, "config.json")
	} else {
		config.Path = *configPathFlag
	}
	fmt.Println(config.Path)
	config.Port = strconv.Itoa(*portFlag)
	config.Password = *passwordFlag
	config.LogDir = *logDirFlag
	config.Init()

	// set windows systray
	windowsApp := NewWindowsApp()
	windowsApp.AddNotifyIcon(
		func() {
			bt.Main(config, appPath)
		})
}
