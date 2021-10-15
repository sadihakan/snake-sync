package internal

import (
	snakeSync "github.com/sadihakan/snake-sync/notify"
)

// Sync ...
type Sync interface {
	NewWatcher() error
	SetChan(chan struct{})
	AddFilePath(file string) error
	Chase()
	SetNotifyCallback(callback snakeSync.Callback)
	Chan() <-chan struct{}
	Stop()
}
