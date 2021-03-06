package main

import (
	"os"
	"fmt"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/spf13/pflag"

	// FIXME:
	"github.com/wutaosamuel/bterminal/conf"
	bt "github.com/wutaosamuel/bterminal"
	// "../../conf"
	// bt "../../../bterminal"
)

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

	// set config
	config := conf.NewConfig()
	config.Path = *configPathFlag
	config.Port = strconv.Itoa(*portFlag)
	config.Password = *passwordFlag
	config.LogDir = *logDirFlag
	config.Init()

	// set App path
	_, thisPath, _, ok := runtime.Caller(0)
	if !ok {
		panic("Set App Path fail")
	}
	appPath := filepath.Dir(filepath.Dir(filepath.Dir(thisPath)))
	// FIXME: remove it
	fmt.Println("app path: " + appPath)

	// do main func
	bt.Main(config, appPath)
}
