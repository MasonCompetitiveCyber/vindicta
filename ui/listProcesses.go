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


	// // Start a goroutine to periodically update the view with listening sockets
	// go func() {
	// 	for {
	// 		// Get only listening TCP sockets
	// 		tabs, err := netstat.TCPSocks(func(s *netstat.SockTabEntry) bool {
	// 			return s.State == netstat.Listen
	// 		})
	//
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	//
	// 		// Prepare the string to display in the view
	// 		var result string
	// 		for _, v := range tabs {
	//
 //                procs, err := ps.FindProcess(int(v.UID))
 //                
 //                if err != nil {
 //                    continue
 //                }
	//
 //                result += fmt.Sprintf("%-10d%-10d%-20s\n", procs.Pid(), procs.PPid(), procs.Executable())
	//
	// 		}
	//
	// 		// Update the view in the main UI thread
	// 		newWrapper := func() {
	// 			view.SetText(result)
	// 		}
	// 		cviewApp.QueueUpdateDraw(newWrapper)
	//
	// 		// Sleep for a while before fetching the sockets again
	// 		time.Sleep(time.Second)
	// 	}
	// }()
	//

    return view
}
