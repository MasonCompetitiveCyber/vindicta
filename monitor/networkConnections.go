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

// Track the previous Cwd and time for each PID
var prevCwdMap = make(map[int]string)
var prevTimeMap = make(map[int]time.Time)

// recursive function to trace child PID
func getChildProcess(result *string, tree *pstree.Tree, childPid []int, v *netstat.SockTabEntry) {

	if len(childPid) == 0 {
		return
	}

	for _, children := range childPid {
		if _, err := os.Stat(tree.Procs[children].Stat.Cwd); err == nil {
			childPid := fmt.Sprintf("%d", children)
			childName := tree.Procs[children].Name
			childPwd := tree.Procs[children].Stat.Cwd

			*result += fmt.Sprintf("[black:blue]%-20s%-20s%-30s%-20s%-30s%-10d%-30s\n", childName, childPid, v.LocalAddr, v.State, v.RemoteAddr, v.UID, childPwd)

		}
		getChildProcess(result, tree, tree.Procs[children].Children, v)
	}

}

// Display established network connections and process information related to it
func DisplaySocks(cviewApp *cview.Application) *cview.TextView {
	view := cview.NewTextView()
	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorOrange)
	view.SetTextColor(tcell.ColorMaroon)
	view.SetTextAlign(cview.AlignLeft)
	view.SetTitle("[black:red]Name                Pid                 Local                         Status              Remote                        Uid       Cwd                           ")
	view.SetTitleAlign(cview.AlignLeft)
	view.SetTitleColor(tcell.ColorYellow)
	view.SetDynamicColors(true)

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

				var tree *pstree.Tree

				var childrens []int

				if v.Process != nil {
					pid = fmt.Sprintf("%d", v.Process.Pid)
					name = v.Process.Name
					pidInt = v.Process.Pid

					tree, err = pstree.New()
					if err != nil {
						continue
					}

					childrens = tree.Procs[pidInt].Children

					// Check if the Cwd path exists before accessing it
					if _, err := os.Stat(tree.Procs[pidInt].Stat.Cwd); err == nil {
						cwd = tree.Procs[pidInt].Stat.Cwd

						// Check if Cwd changed in the last 5 seconds
						if prevCwd, ok := prevCwdMap[pidInt]; ok && prevCwd != cwd {
							if time.Since(prevTimeMap[pidInt]) <= 5*time.Second {
								// Highlight the line
								result += fmt.Sprintf("[black:yellow]%-20s%-20s%-30s%-20s%-30s%-10d%-30s\n", name, pid, v.LocalAddr, v.State, v.RemoteAddr, v.UID, cwd)
							}
						}

						// Update the previous Cwd and time for the PID
						prevCwdMap[pidInt] = cwd
						prevTimeMap[pidInt] = time.Now()
					} else {
						cwd = "N/A"
					}

				} else {
					pid = "N/A"
				}

				result += fmt.Sprintf("[black:blue]%-20s%-20s%-30s%-20s%-30s%-10d%-30s\n", name, pid, v.LocalAddr, v.State, v.RemoteAddr, v.UID, cwd)
				if len(childrens) != 0 {
					getChildProcess(&result, tree, childrens, &v)
				}

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
