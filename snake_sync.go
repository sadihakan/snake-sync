package snake_sync

import (
	"github.com/sadihakan/snake-sync/internal"
)

// SnakeSync ...
type SnakeSync struct {
	sync  internal.Sync
	Error error
}

func New() *SnakeSync {
	ss := new(SnakeSync)
	ss.sync = &internal.FSNotify{}
	return ss
}

func (ss *SnakeSync) NewWatcher() *SnakeSync {
	err := ss.sync.NewWatcher()
	ss.Error = err
	return ss
}

func (ss *SnakeSync) AddFilePath(file string) *SnakeSync {
	err := ss.sync.AddFilePath(file)
	ss.Error = err
	return ss
}

func (ss *SnakeSync) Chase(done chan bool) {
	ss.sync.Chase(done)
}
