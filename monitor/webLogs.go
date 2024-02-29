package monitor

import (
	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

func MonitorWebLogs(cviewApp *cview.Application) *cview.TabbedPanels {
	// Web Server Logs Filtering
	view := cview.NewTabbedPanels()

	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorYellow)
	view.SetBorderAttributes(tcell.AttrBold)
	view.SetTabBackgroundColor(tcell.ColorPurple)
	view.SetTabTextColor(tcell.ColorWhite)
	view.SetTabBackgroundColorFocused(tcell.ColorGreen)

	// Todo: Defined Text Primitive Below in this file and call it in their respective tabs
	view.AddTab("1xx", "1xx", cview.NewTextView())
	view.AddTab("2xx", "2xx", cview.NewTextView())
	view.AddTab("3xx", "3xx", cview.NewTextView())
	view.AddTab("4xx", "4xx", cview.NewTextView())
	view.AddTab("5xx", "5xx", cview.NewTextView())

	return view
}
