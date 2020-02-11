package utils

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

// TestDoExec test
func TestDoExec(t *testing.T) {
	testExec := NewExec()
	testExec.Name = "tmp"
	logName := testExec.Name + testExec.GetNameID8b() + ".log"
	testExec.Command = "ls -l"
	testExec.DoExec()
	testExec.SetTime("*/2", "*", "*", "*", "*")
	testExec.CronOP = CronStart
	testExec.StartCron()
	t.Log(testExec.done)
	time.Sleep(60)
	testExec.CronOP = CronEnd
	testExec.StopCron()
	t.Log(testExec.done)

	f, err := ioutil.ReadFile(logName)
	if err != nil {
		t.Fatal("Read log error")
	}
	//str := str(f)  // bytes convert to "string"
	t.Log(f)

	if err := os.Remove(logName); err != nil {
		t.Fatal("remove file fail")
	}
}
