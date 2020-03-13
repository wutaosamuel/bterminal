package main

import (
	"log"

	"github.com/lxn/walk"

	"github.com/wutaosamuel/bterminal/utils"
)

// WindowsApp for windows system
type WindowsApp struct {
	MainWindows *walk.MainWindow
	Icon        *walk.NotifyIcon
}

// NewWindowsApp create a new windows
func NewWindowsApp() *WindowsApp {
	mainWindows, err := walk.NewMainWindow()
	utils.CheckPanic(err)
	defer mainWindows.Dispose()
	icon, err := walk.NewNotifyIcon(mainWindows)
	utils.CheckPanic(err)
	defer icon.Dispose()
	return &WindowsApp{
		mainWindows,
		icon,
	}
}

/////////////////// Public ///////////////////

// AddNotifyIcon add a notify with icon
func (w *WindowsApp) AddNotifyIcon(action func()) {
	// variables
	var err error
	// init mainwindows & icon
	w.MainWindows, err = walk.NewMainWindow()
	defer w.MainWindows.Dispose()
	utils.CheckPanic(err)
	w.Icon, err = walk.NewNotifyIcon(w.MainWindows)
	defer w.Icon.Dispose()
	utils.CheckPanic(err)
	// set icon
	//icon, err := walk.Resources.Icon(filepath.Join(appPath, "cmd", "bterminalWin", "icon.ico"))
	icon, err := walk.NewIconFromResourceId(3)
	utils.CheckPanic(err)
	if err := w.Icon.SetIcon(icon); err != nil {
		panic(err)
	}
	// set mouse button, display info by windows notify system
	// left button: display extra info
	//leftButtonAction := func(info string, button walk.MouseButton) {
		//if button != walk.LeftButton {
			//return
		//}
		//if err := w.Icon.ShowCustom(
			//"bterminal",
			//info,
			//icon); err != nil {
			//panic(err)
		//}
	//}
	//w.Icon.MouseDown().Attach(
		//func(x, y int, button walk.MouseButton) {
			//leftButtonAction("Status: pendding", button)
		//})

	// set actions
	// TODO: setting action
	// TODO: stop action
	startAction := w.addAction(nil, "Start")
	exitAction := w.addAction(nil, "Exit")
	// set start action
	startAction.Triggered().Attach(
		func() {
			go func() {
				action()
			}()
			startAction.SetChecked(true)
			startAction.SetEnabled(false)
		})
	// set exit action
	exitAction.Triggered().Attach(
		func() {
			walk.App().Exit(0)
		})

	// initially, notify icon is hidden, make it visible
	if err := w.Icon.SetVisible(true); err != nil {
		log.Fatal(err)
	}
	// TODO: display tip info of running or padding
	// display tip info, when mouse @ systray
	if err := w.Icon.SetToolTip("bterminal"); err != nil {
		panic(err)
	}
	w.MainWindows.Run()
}

/////////////////// Private ///////////////////

// addAction add action into menu
func (w *WindowsApp) addAction(menu *walk.Menu, name string) *walk.Action {
	action := walk.NewAction()
	action.SetText(name)
	if menu != nil {
		menu.Actions().Add(action)
	}
	if menu == nil {
		w.Icon.ContextMenu().Actions().Add(action)
	}
	return action
}
