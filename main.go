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
	panels.AddTab("ssh", "[1] SSH", ssh)
	panels.AddTab("NetAndProc", "[2] Network and Processes", netproc)
	panels.AddTab("filesystem", "[3] Filesystem", file)
	panels.AddTab("firewall", "[4] Firewall", cview.NewTextView())
	panels.AddTab("webserver", "[5] Webserver", cview.NewTextView())
	panels.AddTab("services", "[6] Services", cview.NewTextView())
	panels.AddTab("kill", "[7] Kill Process", cview.NewTextView())

    app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
        if event.Rune() == 113 {  // 113 means q
            app.Stop()
        } else if event.Rune() == 49 { // 49 means 1
            panels.SetCurrentTab("ssh")
        } else if event.Rune() == 50 { // 50 means 2
            panels.SetCurrentTab("NetAndProc")
        } else if event.Rune() == 51 {
            panels.SetCurrentTab("filesystem")
        } else if event.Rune() == 52 {
            panels.SetCurrentTab("firewall")
        } else if event.Rune() == 53 {
            panels.SetCurrentTab("webserver")
        } else if event.Rune() == 54 {
            panels.SetCurrentTab("services")
        } else if event.Rune() == 55 {
            panels.SetCurrentTab("kill")
        }

        return event
    })

	// Set Panels as Root
	app.SetRoot(panels, true)

	if err := app.Run(); err != nil {
		panic(err)
	}

}
