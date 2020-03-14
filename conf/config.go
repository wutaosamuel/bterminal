package conf

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"os"

	"github.com/wutaosamuel/bterminal/utils"
)

// Config contain settings
type Config struct {
	Path     string
	Port     string
	Password string
	LogDir   string
}

/////////////////// Setter && Getter ///////////////////

// NewConfig is a constructor
func NewConfig() *Config {
	return &Config{}
}

// Init to check config
func (c *Config) Init() {
	if c.Path != "" {
		c.SetPath(c.Path)
	}
	c.SetPort(c.Port)
	c.SetLogDir(c.LogDir)
}

// SetPath set config path
// check config file is exist
func (c *Config) SetPath(name string) {
	isFile, err := utils.IsFile(name)
	if err != nil {
		panic(utils.Errs("Set Path Error: ", err))
	}
	if !isFile {
		panic(utils.Err("Set Path Error: Not a file"))
	}
	c.Path = name
}

// SetPort set port
// check port can work
// TODO: display port is busy on panic function
func (c *Config) SetPort(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(utils.Errs("Set Port Error: ", err))
	}
	listener.Close()
	c.Port = port
}

// SetLogDir set log path
// if no dir path, it will set default at app/log directory
func (c *Config) SetLogDir(dir string) {
	if dir == "" {
		return
	}
	isDir, err := utils.IsDir(dir)
	if err != nil {
		panic(utils.Errs("Set LogDir Error: ", err))
	}
	if !isDir {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			panic(utils.Errs("Create LogDir Error: ", err))
		}
	}

}

/////////////////// Main ///////////////////

// JConfig for reading or writing json config
type JConfig struct {
	Port     string `json:"port,omitempty"`
	Password string `json:"password,omitempty"`
	LogDir   string `json:"logDir,omitempty"`
}

// Load load config file
func (c *Config) Load(name string) {
	config := &JConfig{}
	c.SetPath(name)
	buffer, err := ioutil.ReadFile(c.Path)
	if err != nil {
		panic(utils.Errs("Config Error: Cannot read file.", err))
	}
	if err = json.Unmarshal(buffer, config); err != nil {
		panic(utils.Errs("Config Error: Cannot decode json", err))
	}
	c.Port = config.Port
	c.Password = config.Password
	c.LogDir = config.LogDir
	// check config
	c.Init()
}

// UpdateJSON update json
func (c *Config) UpdateJSON() {
	config := &JConfig{}
	isFile, err := utils.IsFile(c.Path)
	if err != nil {
		panic(utils.Errs("Open Config Error: ", err))
	}
	if !isFile {
		os.Create(c.Path)
	}
	buffer, err := ioutil.ReadFile(c.Path)
	if err != nil {
		panic(utils.Errs("Config Error: Cannot read file.", err))
	}
	err = json.Unmarshal(buffer, config)
	if config.Port != c.Port || config.Password != c.Password {
		config.Port = c.Port
		config.Password = c.Password
		config.LogDir = c.LogDir
		buffer, err = json.Marshal(&config)
		if err != nil {
			// FIXME: possible use log
			panic(utils.Errs("Config Error: Cannot marshal file.", err))
		}
		err = ioutil.WriteFile(c.Path, buffer, 0777)
		if err != nil {
			panic(utils.Errs("Config Error: Cannot write file.", err))
		}
	}
}

// ChangePassword change password
func (c *Config) ChangePassword(current string, new1 string, new2 string) error {
	if new1 != new2 {
		return errors.New("password not match")
	}
	if current != c.Password {
		return errors.New("password not correct")
	}
	c.Password = new1
	c.UpdateJSON()
	return nil
}
