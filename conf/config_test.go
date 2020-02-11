package conf

import (
	"os"
	"testing"
)

// TestLoad test config load
func TestConfig(t *testing.T) {
	tmpLog := "./testload.log"
	tmpConfig := &Config{}
	tmpConfig.SetPassword("1234567")
	tmpConfig.LogPath = "tmp-config"
	_, err := os.Create(tmpLog)
	if err != nil {
		t.Fatal("create tmp log fail")
	}
	err = tmpConfig.UpdateJSON()
	if err != nil {
		t.Error("update config error")
	}
	config := &Config{}
	err = config.Load(tmpLog)
	if err != nil {
		t.Fatal("load log fail")
	}
	err = config.ChangePassword("1234567", "TestPass", "TestPass")
	if err != nil {
		t.Fatal("change password failed")
	}
	if err := os.Remove(tmpLog); err != nil {
		t.Fatal("remove file fail")
	}
}
