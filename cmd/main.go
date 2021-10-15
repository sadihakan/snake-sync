package main

import (
	"errors"
	"flag"
	"fmt"
	snakesync "github.com/sadihakan/snake-sync"
	"github.com/sadihakan/snake-sync/notify"
)

var (
	scanStyle bool
	path      string
)

type NotifyCallback struct {
	notify.Callback
}

func (NotifyCallback) Notify(notify notify.Notify) {
	fmt.Println("Something happened: ", notify)
}

func main() {

	flag.BoolVar(&scanStyle, "scan", false, "")
	flag.StringVar(&path, "path", "", "File path")
	flag.Parse()

	if scanStyle {
		fmt.Println("Add file path: ")
		if _, err := fmt.Scanln(&path); err != nil {
			panic(errors.New("Path cannot be nil"))
		}
	}

	// Make callback
	cb := new(NotifyCallback)

	// Inıtılıze snake sync
	ss := snakesync.New()

	//Create Watcher
	ss.NewWatcher()

	//Set Callback
	ss.SetNotifyCallback(cb)

	//Add file path
	ss.AddFilePath(path)

	if ss.Error != nil {
		panic(ss.Error)
	}

	go ss.Chase()

	<-ss.Chan()
}
