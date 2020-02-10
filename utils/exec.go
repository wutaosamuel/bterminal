package utils

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// DoExec execute command and log recording
// phrase command to args[] first
// exec comand
// output log
func DoExec(logName string, command string) {
	args := strings.Fields(strings.TrimSpace(command))
	cmd := exec.Command(args[0], args[1:]...)
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)

	f, err := os.OpenFile(
		logName,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalf("Error: open log file - %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error: cannot run output - %v", err)
	}
	log.Printf("The output - %s\n", out)
}
