package ui

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"github.com/fsnotify/fsnotify"
	"log"
	"strings"
	"sync"
)

func FileWatcher(path string, tab *widget.Label) {

	// Define watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Create a slice to store the previous events
	// Only records latest 1337 events
	maxEvents := 1337
	previousEvents := make([]string, 0, maxEvents)

	// Create a mutex to synchronize access to the slice of events
	var mutex sync.Mutex

	// Make Channel
	done := make(chan bool)

	// Monitor Here
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				// Lock the mutex before accessing the slice of events
				mutex.Lock()
				// Add the latest event to the beginning of the slice
				previousEvents = append([]string{event.String()}, previousEvents...)
				// Remove the last event if the slice exceeds the maximum size
				if len(previousEvents) > maxEvents {
					previousEvents = previousEvents[:maxEvents]
				}
				// Build the string from the slice using a string builder
				var sb strings.Builder
				for _, event := range previousEvents {
					sb.WriteString(event)
					sb.WriteString("\n")
				}

				// Set the string as the label text
				tab.SetText(sb.String())

				// Unlock the mutex after modifying the slice of events
				mutex.Unlock()
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
