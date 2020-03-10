package main

import (
	"github.com/lxn/walk"
)

// WindowsApp for windows system
type WindowsApp struct {
	*walk.MainWindows
	Icon *walk.NotifyIcon
}
