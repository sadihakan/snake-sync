package snake_sync

import (
	"fmt"
	"github.com/sadihakan/snake-sync/notify"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"
)

type TestNotifyCallback struct {
	notify.Callback
	events []notify.Notify
}

func (tt *TestNotifyCallback) Notify(notify notify.Notify) {
	fmt.Println("Came here: ", notify)
	tt.events = append(tt.events, notify)
}

func TestSnakeSync_New(t *testing.T) {
	dir, _ := os.Getwd()
	defer removeTestFiles(filepath.Join(dir, "test"))
	err := os.MkdirAll(filepath.Join(dir, "test"), os.ModePerm)
	assert.Nil(t, err)
	cb := new(TestNotifyCallback)
	cb.events = make([]notify.Notify, 0)

	ss := New()

	err = ss.NewWatcher()
	assert.Nil(t, err)
	ss.SetNotifyCallback(cb)
	err = ss.AddFilePath(path.Join(dir, "test"))
	assert.Nil(t, err)

	time.AfterFunc(time.Second*5, func() {
		_, err := os.Create(filepath.Join(dir, "test", "testfile.jpg"))
		assert.Nil(t, err)
		time.AfterFunc(time.Second*2, func() {
			ss.Stop()
		})
	})

	if ss.Error != nil {
		panic(ss.Error)
	}

	go ss.Chase()

	<-ss.Chan()

	assert.Len(t, cb.events, 1)
	fmt.Println("test finished")
}

func removeTestFiles(dir string) {
	_ = filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if info.Name() != "README.md" {
			_ = os.Remove(path)
		}
		return nil
	})
}
