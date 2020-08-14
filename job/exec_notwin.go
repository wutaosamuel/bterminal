// +build !windows

package job

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// DoExecute execute command and log recording
// phrase command to args[] first
// exec comand
// output log
// TODO: set default system log path
func DoExecute(logName string, command string) {
	commandToRun := os.ExpandEnv(command)
	args := strings.Fields(strings.TrimSpace(commandToRun))
	cmd := exec.Command(args[0], args[1:]...)
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)

	f, err := os.OpenFile(
		logName,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	log.SetOutput(f)

	out, err := cmd.CombinedOutput()
	if err == nil {
		log.Println(command)
		log.Println("done")
	}
	log.Printf("The output:\n\n%s\n", string(out))
}
