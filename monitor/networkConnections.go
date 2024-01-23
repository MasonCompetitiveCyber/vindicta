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

type ProcessConfig struct {
	Cwd  string
	Time time.Time
}

var processConfigMap = make(map[int]ProcessConfig)

func getResultString(process pstree.Process, v *netstat.SockTabEntry) string {
	var pid = process.Stat.PID
	var pidString = fmt.Sprintf("%d", pid)
	var cwd = process.Stat.Cwd
	var processName = process.Name

	var colorStyle string

	if processConfig, ok := processConfigMap[pid]; ok {
		var prevCwd = processConfig.Cwd
		if prevCwd != cwd && time.Since(processConfig.Time) <= 5*time.Second {
			colorStyle = "[black:yellow]"
		} else {
			colorStyle = "[black:blue]"
		}
	} else {
		colorStyle = "[black:blue]"
	}

	processConfigMap[pid] = ProcessConfig{
		Cwd:  cwd,
		Time: time.Now(),
	}

	return fmt.Sprintf("%s%-20s%-20s%-30s%-20s%-30s%-10d%-30s\n", colorStyle, processName, pidString, v.LocalAddr, v.State, v.RemoteAddr, v.UID, cwd)
}

func getChildProcess(result *string, tree *pstree.Tree, childPid []int, v *netstat.SockTabEntry) {

	if len(childPid) == 0 {
		return
	}

	for _, children := range childPid {
		if _, err := os.Stat(tree.Procs[children].Stat.Cwd); err == nil {
			*result += getResultString(tree.Procs[children], v)
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
	view.SetTitle("[black:violet:br]Name                Pid                 Local                         Status              Remote                        Uid       Cwd                           ")
	view.SetTitleAlign(cview.AlignLeft)
	view.SetTitleColor(tcell.ColorYellow)
	view.SetDynamicColors(true)

	go func() {
		for {

			// TCP Established Connections
			tcpTabs, err := netstat.TCPSocks(func(s *netstat.SockTabEntry) bool {
				return s.State == netstat.Established
			})
			if err != nil {
				log.Fatal(err)
			}

			// UDP Established Connections
			udpTabs, err := netstat.UDPSocks(func(s *netstat.SockTabEntry) bool {
				return s.State == netstat.Established
			})
			if err != nil {
				log.Fatal(err)
			}

			// All Established Connections
			tabs := append(tcpTabs, udpTabs...)

			var result string

			for _, v := range tabs {

				//process Information
				var pidInt int
				var tree *pstree.Tree
				var childrens []int

				if v.Process != nil {

					pidInt = v.Process.Pid
					tree, err = pstree.New()
					if err != nil {
						continue
					}

					childrens = tree.Procs[pidInt].Children

					if _, err := os.Stat(tree.Procs[pidInt].Stat.Cwd); err == nil {

						result += getResultString(tree.Procs[pidInt], &v)

					}

				}

				if len(childrens) != 0 {
					getChildProcess(&result, tree, childrens, &v)
				}

			}

			viewWrapper := func() {
				view.SetText(result)
			}
			cviewApp.QueueUpdateDraw(viewWrapper)

			time.Sleep(time.Second)
		}
	}()

	return view
}
