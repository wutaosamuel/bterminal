package conf

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"

	"../utils"
)

// Config contain settings
type Config struct {
	Path     string
	Port     string
	password string
	LogPath  string
}

// NewConfig is a constructor
func NewConfig() *Config {
	return &Config{}
}

/////////////////// setter && getter ///////////////////

// SetPath set name
// check config file is exist
func (c *Config) SetPath(name string) error {
	isFile, err := utils.IsFile(name)
	if err != nil {
		return utils.Errs("Set Path Error: ", err)
	}
	if !isFile {
		return utils.Err("Set Path Error: Not a file")
	}
	c.Path = name
	return nil
}

// SetPort set port
// check port can work
func (c *Config) SetPort(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		listener.Close()
		panic(err)
	}
	listener.Close()
	c.Port = port
	return nil
}

// SetPassword set password
func (c *Config) SetPassword(password string) error {
	c.password = password
	return nil
}

// GetPath get config name
func (c *Config) GetPath() string { return c.Path }

// GetPort get port
func (c *Config) GetPort() string { return c.Port }

/////////////////// main ///////////////////

// JConfig for reading or writing json config
	type JConfig struct {
		Port     string		`json:"port,omitempty"`
		Password string		`json:"password,omitempty"`
		LogPath  string		`json:"logPath,omitempty"`
	}

// Load load config file
func (c *Config) Load(name string) error {
	config := &JConfig{}
	if err := c.SetPath(name); err != nil {
		return err
	}
	buffer, err := ioutil.ReadFile(c.Path)
	if err != nil {
		return utils.Errs("Config Error: Cannot read file.", err)
	}
	if err = json.Unmarshal(buffer, config); err != nil {
		return utils.Errs("Config Error: Cannot decode json", err)
	}
	c.Port = config.Port
	c.password = config.Password
	c.LogPath = config.LogPath
	return nil
}

// UpdateJSON update json
func (c *Config) UpdateJSON() error {
	config := &JConfig{}
	buffer, err := ioutil.ReadFile(c.Path)
	if err != nil {
		return utils.Errs("Config Error: Cannot read file.", err)
	}
	err = json.Unmarshal(buffer, config)
	if config.Port != c.Port || config.Password != c.password {
		config.Port = c.Port
		config.Password = c.password
		config.LogPath = c.LogPath
		buffer, err = json.Marshal(&config)
		if err != nil {
			return utils.Errs("Config Error: Cannot marshal file.", err)
		}
		err = ioutil.WriteFile(c.Path, buffer, 0777)
		if err != nil {
			return utils.Errs("Config Error: Cannot write file.", err)
		}
	}
	return nil
}

// ChangePassword change password
func (c *Config) ChangePassword(current string, new1 string, new2 string) error {
	if new1 != new2 {
		return errors.New("password not match")
	}
	// FIXME: no password
	if current != c.password {
		return errors.New("password not correct")
	}
	c.password = new1
	if err := c.UpdateJSON(); err != nil {
		return err
	}
	return nil
}
