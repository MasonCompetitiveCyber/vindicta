package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/MasonCompetitiveCyber/vindicta/ui"
)

func main() {

    // Define App
	myApp := app.New()

    // Define Windows
	myWindow := myApp.NewWindow("Vindicta")

    // Widget
    successLogs := widget.NewLabel("")
    errorLogs := widget.NewLabel("")


    // Define Tabs for the windows
	tabs := container.NewAppTabs(
		container.NewTabItem("Successful SSH Logins", successLogs),
		container.NewTabItem("Failed SSH Logins", errorLogs),
		container.NewTabItem("Firewall Configuration", widget.NewLabel("Firewall")),
		container.NewTabItem("Network Connections", widget.NewLabel("Network")),
		container.NewTabItem("Webserver Logs", widget.NewLabel("Webserver")),
		container.NewTabItem("Running Services", widget.NewLabel("Services")),
		container.NewTabItem("Suspicious Processes", widget.NewLabel("Processes")),
		container.NewTabItem("FileSystem Notifications", widget.NewLabel("FileSystem")),
	)

    // Sets tab location; looks best on the top
	tabs.SetTabLocation(container.TabLocationTop)

    // Sets window's content to those defined tabs
	myWindow.SetContent(tabs)


    // Go Routines

    // Successful SSH Logins
    go func() {
        for range time.Tick(time.Second / 2) {
            ui.AccessLog(successLogs)
        }
    }()


    // Failed SSH Logins
    go func() {
        for range time.Tick(time.Second / 2) {
            ui.ErrorLog(errorLogs)
        }
    }()

    // Shows and Runs the windows
	myWindow.ShowAndRun()
}

