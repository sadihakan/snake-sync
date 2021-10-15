package main

import (
	"errors"
	"flag"
	"fmt"
	snakesync "github.com/sadihakan/snake-sync"
)

var (
	scanStyle bool
	path string
)


func main() {

	flag.BoolVar(&scanStyle, "scan", false, "")
	flag.StringVar(&path, "path", "", "File path")

	if scanStyle {
		fmt.Println("Add file path: ")
		if _, err := fmt.Scanln(&path); err != nil {
			panic(errors.New("Path cannot be nil"))
		}
	}

	ss := snakesync.New()

	//Create Watcher
	ss.NewWatcher()

	//Add file path
	ss.AddFilePath(path)

	done := make(chan bool)

	if ss.Error != nil {
		panic(ss.Error)
	}

	ss.Chase(done)

	<-done
}
