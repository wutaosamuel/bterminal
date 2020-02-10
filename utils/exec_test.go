package utils

import (
	"io/ioutil"
	"os"
	"testing"
)

// TestDoExec test
func TestDoExec(t *testing.T) {
	filename := "tmp.log"
	command := "ls -l"
	DoExec(filename, command)

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal("Read log error")
	}

	//str := str(f)  // bytes convert to "string"

	t.Log(f)

	if err := os.Remove(filename); err != nil {
		t.Fatal("remove file fail")
	}
}
