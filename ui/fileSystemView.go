package ui

import (
    "code.rocketnine.space/tslocum/cview"
    "github.com/gdamore/tcell/v2"
)

func FileSystemPanel(cviewApp *cview.Application) *cview.TextView {
    view := cview.NewTextView()
    view.SetTitle("Suspicious File System Activities")
    view.SetBorder(true)
    view.SetBorderColor(tcell.ColorPurple)
    view.SetTextAlign(cview.AlignLeft)
    view.SetPadding(2, 2, 2, 2)
    view.SetTextColor(tcell.ColorRed)

    view.SetText("FILESSSSS")
    // watcher, err := fsnotify.NewWatcher()
    // if err != nil {
    //     log.Fatal(err)
    // }
    // defer watcher.Close()
    //
    // // Add a watch to the watcher for the specified directory
    // err = watcher.Add("/tmp")
    // if err != nil {
    //     log.Fatal(err)
    // }
    //
    // // Start a goroutine to watch for file system changes
    // go func() {
    //     for {
    //         select {
    //         case event := <-watcher.Events:
    //             // Prepare the string to display in the view
    //             result := event.String()
    //
    //             // Update the view in the main UI thread
    //             viewWrapper := func() {
    //                 view.SetText(result)
    //             }
    //             cviewApp.QueueUpdateDraw(viewWrapper)
    //         case err := <-watcher.Errors:
    //             if err != nil {
    //                 myErr := fmt.Sprintf("error: %v", err)
    //                 log.Fatal(myErr)
    //             }
    //         }
    //     }
    // }()

    return view
}

