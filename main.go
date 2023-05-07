package main

import (
	"code.rocketnine.space/tslocum/cview"
	"github.com/MasonCompetitiveCyber/vindicta/monitor"
	"github.com/gdamore/tcell/v2"
)

func main() {
	// Define the Application
	app := cview.NewApplication()

	// Enable Using Mouse
	app.EnableMouse(true)

	// Define Tabbed Panels
	panels := cview.NewTabbedPanels()

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

	// Redraw on some change
	panels.SetChangedFunc(func() { app.Draw() })

	// Call UI Tabs for each
	// SSH Logs Tab
	ssh := monitor.SshLogs(app)
	// File System Monitoring Tab
	file := monitor.FileSystemPanel(app)
	// Network Connections and Process Monitoring Tab
	netproc := monitor.DisplaySocks(app)

	// Attach The Tabs Above To The Panels
	panels.AddTab("ssh", "SSH", ssh)
	panels.AddTab("NetAndProc", "Network and Processes", netproc)
	panels.AddTab("filesystem", "Filesystem", file)
	panels.AddTab("firewall", "Firewall", cview.NewTextView())
	panels.AddTab("webserver", "Webserver", cview.NewTextView())
	panels.AddTab("services", "Services", cview.NewTextView())

	// Set Panels as Root
	app.SetRoot(panels, true)

	if err := app.Run(); err != nil {
		panic(err)
	}

}
