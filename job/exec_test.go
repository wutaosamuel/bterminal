package job

import (
	"io/ioutil"
	"testing"
	"time"
)

// TestDoExec test
func TestDoExec(t *testing.T) {
	testExec := NewExec()
	testExec.Name = "tmp"
	logName := testExec.Name + testExec.GetNameID8b() + ".log"
	testExec.Command = "ls -l"
	testExec.SetCronTime("*/1", "*", "*", "*", "*")
	t.Log(testExec)
	testExec.StartCron()
	t.Log("Cron Start")
	time.Sleep(time.Duration(121)*time.Second)
	testExec.StopCron()
	t.Log("Cron Stop")

	f, err := ioutil.ReadFile(logName)
	if err != nil {
		t.Fatal("Read log error")
	}
	//str := str(f)  // bytes convert to "string"
	t.Log(f)

	if err := testExec.DeleteLog(); err != nil {
		t.Fatal("remove file fail")
	}
}
