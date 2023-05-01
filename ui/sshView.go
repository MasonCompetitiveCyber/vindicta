package ui

import (
	"code.rocketnine.space/tslocum/cview"
	"github.com/MasonCompetitiveCyber/vindicta/detect"
	"github.com/gdamore/tcell/v2"
)

func SshPanel() *cview.Grid {

    // Define UI Over here
    grid := cview.NewGrid()
    grid.SetRows(1,1,0)
    grid.SetBorders(true)
    grid.SetBordersColor(tcell.ColorPurple)

    // 1st row -> Index 0
    // titleBar := setUI("SSH Logs", tcell.ColorDarkBlue)
    titleBar := setUI("SSH LOGS", tcell.ColorSkyblue)

    // 2nd row -> Index 1
    success := setUI("Succcessful SSH Logins", tcell.ColorYellow)
    denied := setUI("Failed SSH Attempts", tcell.ColorYellow)


    // 3rd row -> Index 2
    // Get Logs
    accessLog := setUILogs(detect.AccessLog(), tcell.ColorGreen)
    deniedLog := setUILogs(detect.ErrorLog(), tcell.ColorRed)


    // Add column contents to the grid
    // Name, row index, column index, row span, column span, min grid height, min grid width, focus
	grid.AddItem(titleBar, 0, 0, 1, 4, 0, 0, false)
	grid.AddItem(success, 1, 0, 1, 2, 0, 0, false)
	grid.AddItem(denied, 1, 2, 1, 2, 0, 0, false)
	grid.AddItem(accessLog, 2, 0, 1, 2, 0, 0, false)
	grid.AddItem(deniedLog, 2, 2, 1, 2, 0, 0, false)

    return grid
}


// Settings for each UI components
func setUI(text string, color tcell.Color) *cview.TextView {
    myView := cview.NewTextView()
    myView.SetTextAlign(cview.AlignCenter)
    myView.SetText(text)
    myView.SetTextColor(color)
    return myView
}


func setUILogs(text string, color tcell.Color) *cview.TextView {
    myView := cview.NewTextView()
    myView.SetTextAlign(cview.AlignLeft)
    myView.SetText(text)
    myView.SetTextColor(color)
    return myView
}
