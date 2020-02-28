package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/pflag"

	bt "../../bternimal"
	"../conf"
)

func main() {

	var (
		helpFlag       = pflag.BoolP("help", "h", false, "display usage")
		configPathFlag = pflag.StringP("config", "c", "", "Json config file")
		portFlag       = pflag.IntP("port", "p", 5122, "TCP port for web service")
		passwordFlag   = pflag.StringP("password", "P", "", "password for protecting web service")
		logDirFlag     = pflag.StringP("log", "l", "./", "log directory for keeping job logs")
	)

	pflag.Parse()

	if *helpFlag {
		pflag.Usage()
	}

	fmt.Println(*configPathFlag)
	fmt.Println(*portFlag)
	fmt.Println(*passwordFlag)
	fmt.Println(*logDirFlag)

	config := conf.NewConfig()
	config.Path = *configPathFlag
	config.Port = strconv.Itoa(*portFlag)
	config.Password = *passwordFlag
	config.LogDir = *logDirFlag
	config.Init()

	thisPath,_ := os.Getwd()
	appPath,_ := filepath.Abs(filepath.Dir(thisPath))
	fmt.Println(appPath)
	bt.Main(config, appPath)
}
