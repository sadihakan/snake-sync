package internal

import (
	fsnotify "github.com/fsnotify/fsnotify"
	"github.com/sadihakan/snake-sync/notify"
	"log"
)

type FSNotify struct {
	Sync
	Watcher  *fsnotify.Watcher
	callback notify.Callback
	ch       chan struct{}
}

func (f *FSNotify) NewWatcher() error {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	f.Watcher = w
	return nil
}

func (f *FSNotify) SetChan(ch chan struct{}) {
	f.ch = ch
}

func (f *FSNotify) AddFilePath(file string) error {
	return f.Watcher.Add(file)
}

func (f *FSNotify) SetNotifyCallback(callback notify.Callback) {
	f.callback = callback
}

func (f *FSNotify) Chase() {
	defer f.Watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-f.Watcher.Events:
				if !ok {
					return
				}
				//bb, _ := json.Marshal(event)
				//fmt.Println(string(bb))
				if f.callback != nil {
					f.callback.Notify(notify.Notify{
						EventType: event.Op.String(),
						Path:      event.Name,
					})
				}
			case err, ok := <-f.Watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	<-f.ch
}
