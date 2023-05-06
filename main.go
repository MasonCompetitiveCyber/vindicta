package main

import (
	"code.rocketnine.space/tslocum/cview"
	"fyne.io/fyne/v2/app"
	"github.com/MasonCompetitiveCyber/vindicta/ui"
	"maunium.net/go/tcell"
)

func main() {

	// Define App
	myApp := app.New()

	// Define Windows
	myWindow := myApp.NewWindow("Vindicta")

	// Properties of Panels
	panels.SetBorder(true)
	panels.SetBorderColor(tcell.ColorYellow)
	panels.SetTitle("Vindicta")
	panels.SetTitleColor(tcell.ColorBlue)
	panels.SetTabTextColor(tcell.ColorPurple)
	panels.SetBorderAttributes(tcell.AttrBold)
	panels.SetTabBackgroundColor(tcell.ColorBlueViolet)
	panels.SetTabTextColor(tcell.ColorWhite)
	panels.SetTabBackgroundColorFocused(tcell.ColorOrange)
	panels.SetChangedFunc(func() { app.Draw() })

	// Call UI Tabs for each
	ssh := ui.SshPanel()
	file := ui.FileSystemPanel(app)
	net := ui.DisplaySocks(app)
	procs := ui.ListProcesses(app)

	// Add Tabs For Panels
	panels.AddTab("ssh", "SSH", ssh)
	panels.AddTab("network", "Network", net)
	panels.AddTab("filesystem", "Filesystem", file)
	panels.AddTab("firewall", "Firewall", cview.NewTextView())
	panels.AddTab("webserver", "Webserver", cview.NewTextView())
	panels.AddTab("services", "Services", cview.NewTextView())
	panels.AddTab("processes", "Processes", procs)

	// Sets window's content to those defined tabs
	myWindow.SetContent(tabs)

	if err := app.Run(); err != nil {
		panic(err)
	}

}
