package main

import (
	"errors"
	"flag"
	"fmt"
	snakesync "github.com/sadihakan/snake-sync"
	"github.com/sadihakan/snake-sync/notify"
	"io/ioutil"
	"log"
)

var (
	scanStyle bool
	path      string
)

type NotifyCallback struct {
	notify.Callback
}

func (NotifyCallback) Notify(notify notify.Notify) {
	fmt.Println("Something happened: ", notify.EventType)
}

func main() {

	flag.BoolVar(&scanStyle, "scan", false, "")
	flag.StringVar(&path, "path", "", "File path")
	flag.Parse()

	if scanStyle {
		fmt.Println("Add file path: ")
		if _, err := fmt.Scanln(&path); err != nil {
			panic(errors.New("path cannot be nil"))
		}
	}

	// Make callback
	cb := new(NotifyCallback)

	// Install snake sync
	ss := snakesync.New()

	//Create Watcher
	err := ss.NewWatcher()
	if err != nil {
		return
	}

	//Set Callback
	ss.SetNotifyCallback(cb)

	//Add file path
	err = ss.AddFilePath(path)
	if err != nil {
		return
	}

	if ss.Error != nil {
		panic(ss.Error)
	}

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.Size(), file.IsDir())
	}

	go ss.Chase()

	<-ss.Chan()
}
