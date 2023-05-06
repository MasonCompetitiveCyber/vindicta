package ui

import (
    "time"
	"fmt"
	"log"
	"strings"

	"code.rocketnine.space/tslocum/cview"
	"github.com/fsnotify/fsnotify"
	"github.com/gdamore/tcell/v2"
)

func FileSystemPanel(cviewApp *cview.Application) *cview.TextView {
	view := cview.NewTextView()
	view.SetTitle("Suspicious File System Activities")
	view.SetTitleColor(tcell.ColorGreen)
	view.SetBorder(true)
	view.SetBorderColor(tcell.ColorPurple)
	view.SetTextAlign(cview.AlignLeft)
	view.SetPadding(2, 2, 2, 2)
	view.SetTextColor(tcell.ColorRed)

	// Define watcher to watch for file system events
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	// Add a watch to the watcher for the specified directory
	err = watcher.Add("/tmp")
	if err != nil {
		log.Fatal(err)
	}

	// Keep track of the latest 50 events
	var events []string

	// Start a goroutine to watch for file system changes
	go func() {
        defer watcher.Close()
		for {
			select {
			case event := <-watcher.Events:

				// Prepare the string to display in the view
                result := fmt.Sprintf("%s: %s", time.Now().Format("2006-01-02 15:04:05"), event.String())


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
