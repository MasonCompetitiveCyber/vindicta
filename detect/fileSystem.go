package detect

import (
    "log"

    "github.com/fsnotify/fsnotify"
)


// To Be Completed

func FileWatcher(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Modified file:", event.Name)
				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
					log.Println("Renamed file:", event.Name)
				} else {
					log.Println("Event:", event)
				}
			case err := <-watcher.Errors:
				log.Println("Error:", err)
			}
		}
	}()

	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
