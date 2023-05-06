package ui

import (
	"fmt"
	"log"
	"time"

	"code.rocketnine.space/tslocum/cview"
	"github.com/cakturk/go-netstat/netstat"
	"github.com/gdamore/tcell/v2"
)

func DisplaySocks(cviewApp *cview.Application) *cview.TextView {
	view := cview.NewTextView()
    view.SetBorder(true)
    view.SetBorderColor(tcell.ColorYellow)
    view.SetTextColor(tcell.ColorMaroon)
    view.SetTextAlign(cview.AlignLeft)

	// Start a goroutine to periodically update the view with listening sockets
	go func() {
		for {
			// Get only listening TCP sockets
			tabs, err := netstat.TCPSocks(func(s *netstat.SockTabEntry) bool {
				return s.State == netstat.Listen
			})

			if err != nil {
				log.Fatal(err)
            }

			// Prepare the string to display in the view
			var result string
			for _, v := range tabs {
				// result += fmt.Sprintf("%d    %s    %s\n", v.UID, v.State, v.LocalAddr)
                result += fmt.Sprintf("%-10d%-10s%-20s\n", v.UID, v.State, v.LocalAddr)

			}

			// Update the view in the main UI thread
			viewWrapper := func() {
				view.SetText(result)
			}
			cviewApp.QueueUpdateDraw(viewWrapper)

			// Sleep for a while before fetching the sockets again
			time.Sleep(time.Second)
		}
	}()

	return view
}

