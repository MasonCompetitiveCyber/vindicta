package ui

import (
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

func ListProcesses(cviewApp *cview.Application) *cview.TextView {
    view := cview.NewTextView()
    view.SetBorder(true)
    view.SetBorderColor(tcell.ColorYellow)
    view.SetTextColor(tcell.ColorPurple)
    view.SetText("AMAZING PROCESSES")


    return view
}
