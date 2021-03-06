package conf

import (
	"os"
	"testing"
)

// TestLoad test config load
func TestConfig(t *testing.T) {
	tmpLog := "./testload.log"
	tmpConfig := &Config{}
	tmpConfig.Password = "1234567"
	tmpConfig.LogDir = "tmp-config"
	_, err := os.Create(tmpLog)
	if err != nil {
		t.Fatal("create tmp log fail")
	}
	tmpConfig.UpdateJSON()
	config := &Config{}
	config.Load(tmpLog)
	err = config.ChangePassword("1234567", "TestPass", "TestPass")
	if err != nil {
		t.Fatal("change password failed")
	}
	if err := os.Remove(tmpLog); err != nil {
		t.Fatal("remove file fail")
	}
}
