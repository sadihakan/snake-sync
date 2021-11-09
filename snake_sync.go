package snake_sync

import (
	"github.com/sadihakan/snake-sync/internal"
	"github.com/sadihakan/snake-sync/notify"
)

// SnakeSync ...
type SnakeSync struct {
	sync  internal.Sync
	Error error
	ch chan struct{}
}

// New ...
func New() *SnakeSync {
	ss := new(SnakeSync)
	ss.sync = new(internal.FSNotify)
	ss.ch = make(chan struct{})
	ss.sync.SetChan(ss.ch)
	return ss
}

// Chan ..
func (ss SnakeSync) Chan() <-chan struct{} {
	return ss.ch
}

// SetChan ..
func (ss *SnakeSync) SetChan(ch chan struct{})  {
	ss.ch = ch
}

// Stop ..
func (ss *SnakeSync) Stop() {
	ss.ch<- struct{}{}
	close(ss.ch)
}

// NewWatcher ..
func (ss *SnakeSync) NewWatcher() error {
	err := ss.sync.NewWatcher()
	ss.Error = err
	return nil
}

// AddFilePath ..
func (ss *SnakeSync) AddFilePath(file string) error {
	err := ss.sync.AddFilePath(file)
	ss.Error = err
	return nil
}

// SetNotifyCallback ..
func (ss *SnakeSync) SetNotifyCallback(callback notify.Callback) {
	ss.sync.SetNotifyCallback(callback)
}

// Chase ..
func (ss *SnakeSync) Chase() {
	ss.sync.Chase()
}
