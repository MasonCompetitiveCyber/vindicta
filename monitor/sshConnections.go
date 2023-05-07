package monitor

import (
	"bufio"
	"io"
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

// Check For Logs
func monitorLogs(app *cview.Application, successLogs *cview.TextView, errorLogs *cview.TextView) {

    for{
    // Check ssh logs using journalctl command
    cmd := exec.Command("journalctl", "-u", "sshd")
    out, err := cmd.Output()

    if err == nil {
        // Update successLogs and errorLogs with logs from journalctl output
        lines := strings.Split(string(out), "\n")
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

        return
    }

    // If there was an error with journalctl, check /var/log/auth.log
    logFile, err := os.Open("/var/log/auth.log")
    if err != nil {
        errorLogs.SetText("ERROR: Could not open /var/log/auth.log")
        return
    }
    defer logFile.Close()

    reader := bufio.NewReader(logFile)
    successText := ""
    errorText := ""

    for {
        line, err := reader.ReadString('\n')
        if err != nil && err != io.EOF {
            errorLogs.SetText("ERROR: Problem reading logs from /var/log/auth.log")
            return
        }

        if strings.Contains(line, "sshd") {
            if strings.Contains(line, "Accepted") {
                successText += line
            } else if strings.Contains(line, "Failed") {
                errorText += line
            }
        }

        if err == io.EOF {
            successLogs.SetText(successText)
            errorLogs.SetText(errorText)

            // Schedule an update and redraw of the application's screen
            app.QueueUpdateDraw(func() {
                successLogs.ScrollToEnd()
                errorLogs.ScrollToEnd()
            })
            break
        }
    }
    }

}
