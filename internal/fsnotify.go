package internal

import (
	fsnotify "github.com/fsnotify/fsnotify"
	"log"
)

type FSNotify struct {
	Sync
	Watcher *fsnotify.Watcher
}

func (f *FSNotify) NewWatcher() error {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	f.Watcher = w
	return nil
}

func (f *FSNotify) AddFilePath(file string) error {
	return f.Watcher.Add(file)
}

func (f *FSNotify) Chase(done chan bool) {
	defer f.Watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-f.Watcher.Events:
				if !ok {
					return
				}
				log.Println("modified file:", event.Name, "modified type:", event.Op)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-f.Watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	<-done
}
