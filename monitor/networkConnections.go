package monitor

import (
	"fmt"
	"github.com/sbinet/pstree"
	"log"
	"os"
	"time"

	"code.rocketnine.space/tslocum/cview"
	"github.com/cakturk/go-netstat/netstat"
	"github.com/gdamore/tcell/v2"
)

// Display established network connnections and process information related to it
func DisplaySocks(cviewApp *cview.Application) *cview.TextView {
	view := cview.NewTextView()
	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorAquaMarine)
	view.SetTextColor(tcell.ColorMaroon)
	view.SetTextAlign(cview.AlignLeft)

	// Start a goroutine to periodically update the view with listening sockets
	go func() {
		for {
			// Get only listening TCP sockets
			tabs, err := netstat.TCPSocks(func(s *netstat.SockTabEntry) bool {
				return s.State == netstat.Established
			})

			if err != nil {
				log.Fatal(err)
			}

			// Prepare the string to display in the view
			var result string
			for _, v := range tabs {

				var pid string
				var name string
				var pidInt int
				var cwd string

				if v.Process != nil {
					pid = fmt.Sprintf("%d", v.Process.Pid)
					name = fmt.Sprintf("%s", v.Process.Name)
					pidInt = v.Process.Pid

					tree, err := pstree.New()
					if err != nil {
						continue
					}

					// Check if the Cwd path exists before accessing it
					if _, err := os.Stat(tree.Procs[pidInt].Stat.Cwd); err == nil {
						cwd = tree.Procs[pidInt].Stat.Cwd
					} else {
						cwd = "N/A"
					}

				} else {
					pid = "N/A"
				}

				result += fmt.Sprintf("%-20s%-20s%-20s%-20s%-10d%-30s\n", name, pid, v.LocalAddr, v.State, v.UID, cwd)

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
