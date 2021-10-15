package internal

// Sync ...
type Sync interface {
	NewWatcher() error
	AddFilePath(file string) error
	Chase(done chan bool)
}
