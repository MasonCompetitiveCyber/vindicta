package monitor

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"code.rocketnine.space/tslocum/cview"
	"github.com/fsnotify/fsnotify"
	"github.com/gdamore/tcell/v2"
)

var fileSystemPath string
var watcher *fsnotify.Watcher
var err error

// on-change handler for input
func handleOnChange(text string) {
	fileSystemPath = text
}

func addNewWatcher() {
	//parse the fileSystemPath
	paths := strings.Split(fileSystemPath, ",")

	for _, val := range paths {
		val = strings.TrimSpace(val)
		err := watcher.Add(val)

		if err != nil {
			log.Fatal(err)
		}
	}
}

func getPermissions(filename string) string {
	// get info and potential errors
	info, err := os.Stat(filename)
	// panic if there is an error
	if err != nil {
		return "Deleted"
	}

	// grab permissions (very verbose)
	mode := info.Mode()
	temp := ""

	// Loop through verbose output to just get rwx
	for i := 0; i < 10; i++ {
		temp += string(mode.String()[i])
	}
	// return rwx permissions
	return temp
}

// Input Box to ask for paths for monitoring
func CreateInput(panels *cview.TabbedPanels, app *cview.Application) func() {
	form := cview.NewForm()
	form.AddInputField("File Paths to monitor", "", 100, nil, handleOnChange)
	form.AddButton("Start Monitoring", func() {
		addNewWatcher()
		panels.SetCurrentTab("filesystem") // Set the current tab to "filesystem"
		app.SetRoot(panels, true)          // Set the root view to the tabbed panels
	})
	form.SetBorder(true)
	form.SetBorderColor(tcell.ColorPurple)
	form.SetFieldBackgroundColorFocused(tcell.ColorBlack)
	form.SetFieldTextColorFocused(tcell.ColorBlue)
	form.SetTitle("Files and Directories")
	form.SetTitleAlign(cview.AlignCenter)
	form.SetRect(60, 10, 80, 30)

	return func() {
		app.SetRoot(form, false)
	}
}

// File System Monitoring
func FileSystemPanel(cviewApp *cview.Application) *cview.TextView {
	view := cview.NewTextView()
	view.SetTitle("File System Activities")
	view.SetTitleColor(tcell.ColorGreen)
	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorPurple)
	view.SetTextAlign(cview.AlignLeft)
	view.SetPadding(2, 2, 2, 2)
	view.SetTextColor(tcell.ColorRed)

	watcher, err = fsnotify.NewWatcher()

	//Default directory for watcher
	err = watcher.Add("/tmp")

	if err != nil {
		log.Fatal(err)
	}

	// Keep track of the latest 50 events
	var events []string
	var result string

	// Start a goroutine to watch for file system changes
	go func() {
		defer watcher.Close()
		for {
			select {
			case event := <-watcher.Events:
				// store filename (ex: CHMOD "/tmp/grass")
				filename := event.String()

				// parse filename to only have filename
				filename = strings.Split(filename, "\"")[1]

				// pass parsed filename to getPermissions
				perms := getPermissions(filename)

				// Prepare the string to display in the view
				if perms != "Deleted" {
					result = fmt.Sprintf("%s: %s %s", time.Now().Format("2006-01-02 15:04:05"), perms, event.String())
				} else {
					result = fmt.Sprintf("%s: %-10s %s", time.Now().Format("2006-01-02 15:04:05"), "", event.String())
				}

				// Add the new event to the beginning of the events slice
				events = append([]string{result}, events...)

				// Limit the number of events to 50
				if len(events) > 50 {
					events = events[:50]
				}

				// Concatenate all events into a single string
				text := strings.Join(events, "\n")

				// Update the view in the main UI thread
				viewWrapper := func() {
					view.SetText(text)
				}
				cviewApp.QueueUpdateDraw(viewWrapper)

			case err := <-watcher.Errors:
				if err != nil {
					myErr := fmt.Sprintf("error: %v", err)
					view.SetText(myErr)
				}
			}
		}
	}()

	return view
}
