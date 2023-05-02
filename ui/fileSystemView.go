package ui

import (
    "code.rocketnine.space/tslocum/cview"
    "github.com/gdamore/tcell/v2"
)

// To Be Completed
func FileSystemPanel() *cview.TextView {

	view := cview.NewTextView()
	view.SetTitle("Suspicious File System Activities")
	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorPurple)
	view.SetText("File MASONCC WAS CREATED\nFILE /etc/passwd WAS MODIFIED")
	view.SetTextAlign(cview.AlignLeft)
	view.SetPadding(2, 2, 2, 2)
	view.SetTextColor(tcell.ColorRed)

	return view
}
