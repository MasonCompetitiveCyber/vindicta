package monitor

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sbinet/pstree"

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
			var results string

			// //Get all TCP sockets

			tcp6Tabs, err := netstat.TCP6Socks(func(s *netstat.SockTabEntry) bool {
				return s.State == netstat.Established
			})

			if err != nil {
				log.Fatal(err)
			}

			tcpTabs, err := netstat.TCPSocks(func(s *netstat.SockTabEntry) bool {
				return s.State == netstat.Established
			})

			if err != nil {
				log.Fatal(err)
			}

			results += parseConnectionView(tcp6Tabs)
			results += parseConnectionView(tcpTabs)

			// Update the view in the main UI thread
			viewWrapper := func() {
				view.SetText(results)
			}
			cviewApp.QueueUpdateDraw(viewWrapper)

			// Sleep for a while before fetching the sockets again
			time.Sleep(time.Second)
		}
	}()

	return view
}

func parseConnectionView(tabs []netstat.SockTabEntry) string {

	var result string
	var pid string
	var name string
	var pidInt int
	var cwd string

	for _, v := range tabs {
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

	return result

}
