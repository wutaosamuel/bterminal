package conf

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"os"

	"../utils"
)

// Config contain settings
type Config struct {
	path     string
	port     string
	password string
}

// NewConfig is a constructor
func NewConfig() *Config {
	return &Config{}
}

/////////////////// setter && getter ///////////////////

// SetPath set path
// check config file is exist
func (c *Config) SetPath(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return utils.Errs("Config Error: no such file.", err)
		}
	}
	c.path = path
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
	c.port = port
	return nil
}

// SetPassword set password
func (c *Config) SetPassword(password string) error {
	if c.password != "" {
		return utils.Err("Config Error: password is not empty")
	}
	c.password = password
	return nil
}

// GetPath get config path
func (c *Config) GetPath() string { return c.path }

// GetPort get port
func (c *Config) GetPort() string { return c.port }

/////////////////// main ///////////////////

// Load load config file
func (c *Config) Load(path string) error {
	type JConfig struct {
		port     string
		password string
	}
	config := &JConfig{}
	if c.path == "" && path == "" {
		path = "config.json"
	}
	if err := c.SetPath(path); err != nil {
		return err
	}
	var buffer []byte
	buffer, err := ioutil.ReadFile(c.path)
	if err != nil {
		return utils.Errs("Config Error: Cannot read file.", err)
	}
	err = json.Unmarshal(buffer, &config)
	if err != nil {
		return utils.Errs("Config Error: Cannot decode json", err)
	}
	c.port = config.port
	c.password = config.password
	return nil
}

// UpdateJSON update json
func (c *Config) UpdateJSON() error {
	type JConfig struct {
		port     string
		password string
	}
	var buffer []byte
	config := &JConfig{}
	buffer, err := ioutil.ReadFile(c.path)
	if err != nil {
		return utils.Errs("Config Error: Cannot read file.", err)
	}
	err = json.Unmarshal(buffer, &config)
	if config.port != c.port || config.password != c.password {
		config.port = c.port
		config.password = c.password
		buffer, err = json.Marshal(&config)
		if err != nil {
			return utils.Errs("Config Error: Cannot marshal file.", err)
		}
		err = ioutil.WriteFile(c.path, buffer, 0644)
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
