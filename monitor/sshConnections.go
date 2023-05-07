package monitor

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"code.rocketnine.space/tslocum/cview"
	"github.com/gdamore/tcell/v2"
)

func SshLogs(app *cview.Application) *cview.Grid {

	// Define grid
	grid := cview.NewGrid()

	// Set number of rows
	grid.SetRows(1, 1, 0)

	// Set Border to be true
	grid.SetBorders(true)

	// Set Border Color
	grid.SetBordersColor(tcell.ColorPurple)

	// 1st row -> Index 0
	// titleBar := setUI("SSH Logs", tcell.ColorDarkBlue)
	titleBar := setView("SSH LOGS", tcell.ColorSkyblue)

	// 2nd row -> Index 1
	successTitle := setView("Successful SSH Logins", tcell.ColorYellow)
	deniedTitle := setView("Failed SSH Attempts", tcell.ColorYellow)

	// 3rd row -> Index 0 -> Success Logs
	// 3rd row -> Index 1 -> Failure Logs
	successLogs := setViewLogs("", tcell.ColorGreen)
	errorLogs := setViewLogs("", tcell.ColorRed)

	// Add column contents to the grid
	// Name, row index, column index, row span, column span, min grid height, min grid width, focus
	grid.AddItem(titleBar, 0, 0, 1, 4, 0, 0, false)
	grid.AddItem(successTitle, 1, 0, 1, 2, 0, 0, false)
	grid.AddItem(deniedTitle, 1, 2, 1, 2, 0, 0, false)
	grid.AddItem(successLogs, 2, 0, 1, 2, 0, 0, false)
	grid.AddItem(errorLogs, 2, 2, 1, 2, 0, 0, false)

	// Call monitorLogs in a goroutine to update logs
	go monitorLogs(app, successLogs, errorLogs)

	return grid
}

// Setting color and properties for UI components
func setView(text string, color tcell.Color) *cview.TextView {
	myView := cview.NewTextView()
	myView.SetTextAlign(cview.AlignCenter)
	myView.SetText(text)
	myView.SetTextColor(color)
	return myView
}

// Set color and properties from access and error logs
func setViewLogs(text string, color tcell.Color) *cview.TextView {
	myView := cview.NewTextView()
	myView.SetTextAlign(cview.AlignLeft)
	myView.SetText(text)
	myView.SetTextColor(color)
	myView.SetScrollable(true)
	return myView
}

func monitorLogs(app *cview.Application, successLogs *cview.TextView, errorLogs *cview.TextView) {
	// Read the contents of /etc/os-release to determine the appropriate log file to monitor
	content, err := ioutil.ReadFile("/etc/os-release")
	if err != nil {
		// handle error
	}

	// Convert byte slice to string
	strContent := string(content)

	// Extract the value of ID field
	var osType string
	for _, line := range strings.Split(strContent, "\n") {
		if strings.HasPrefix(line, "ID=") {
			osType = strings.TrimPrefix(line, "ID=")
			break
		}
	}

	// Use the appropriate log file to monitor SSH logs
	for {
		var logFile string
		switch osType {
		case "garuda", "arch", "\"rocky\"", "\"rhel\"", "fedora", "\"centos\"", "\"opensuse-leap\"", "\"ol\"", "\"almalinux\"":
			cmd := exec.Command("journalctl", "-u", "sshd")
			out, err := cmd.Output()
			if err != nil {
				errorLogs.SetText(fmt.Sprintf("ERROR: Could not read logs from %s: %s", "journalctl -u sshd", err.Error()))
				return
			}
			logFile = string(out)
		case "debian", "kali", "ubuntu", "alpine":
			file, err := os.Open("/var/log/auth.log")
			if err != nil {
				errorLogs.SetText(fmt.Sprintf("ERROR: Could not read logs from %s: %s", "/var/log/auth.log", err.Error()))
				return
			}
			defer file.Close()
			fileInfo, err := file.Stat()
			if err != nil {
				errorLogs.SetText(fmt.Sprintf("ERROR: Could not read logs from %s: %s", "/var/log/auth.log", err.Error()))
				return
			}
			size := fileInfo.Size()
			buffer := make([]byte, size)
			_, err = file.Read(buffer)
			if err != nil {
				errorLogs.SetText(fmt.Sprintf("ERROR: Could not read logs from %s: %s", "/var/log/auth.log", err.Error()))
				return
			}
			logFile = string(buffer)
		default:
			errorLogs.SetText("ERROR: Unknown or unsupported OS type")
			return
		}

		// Monitor the SSH logs in the selected log file
		lines := strings.Split(logFile, "\n")
		successText := ""
		errorText := ""
		for _, line := range lines {
			if strings.Contains(line, "Accepted") {
				successText += line + "\n"
			} else if strings.Contains(line, "Failed") {
				errorText += line + "\n"
			}
		}
		successLogs.SetText(successText)
		errorLogs.SetText(errorText)

		// Schedule an update and redraw of the application's screen
		app.QueueUpdateDraw(func() {
			successLogs.ScrollToEnd()
			errorLogs.ScrollToEnd()
		})
	}
}
