package ui

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"github.com/fsnotify/fsnotify"
	"log"
)

func FileWatcher(path string, tab *widget.Label) {

	// Define watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Make Channel
	done := make(chan bool)

	// Monitor Here
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					wrote := fmt.Sprintf("Modified file: %s", event.Name)
					tab.SetText(wrote)
				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
					renamed := fmt.Sprintf("Renamed file: %s", event.Name)
					tab.SetText(renamed)
				} else if event.Op&fsnotify.Create == fsnotify.Create {
					created := fmt.Sprintf("Created file: %s", event.Name)
					tab.SetText(created)
				} else if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					perm := fmt.Sprintf("Permission changed on file: %s", event.Name)
					tab.SetText(perm)
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					removed := fmt.Sprintf("Removed file: %s", event.Name)
					tab.SetText(removed)
				}
			case err := <-watcher.Errors:
				criticalError := fmt.Sprintf("%s", err)
				tab.SetText(criticalError)
			}
		}
	}()

	// Add Path to Monitor
	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	// Receive from done channel
	<-done
}
