package snake_sync

import (
	"fmt"
	"github.com/sadihakan/snake-sync/notify"
	"github.com/stretchr/testify/assert"
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
	os.RemoveAll(path.Join(dir, "test"))
	err := os.Mkdir(path.Join(dir, "test"), os.ModePerm)
	assert.Nil(t, err)
	cb := new(TestNotifyCallback)
	cb.events = make([]notify.Notify, 0)

	ss := New()

	ss.NewWatcher()
	ss.SetNotifyCallback(cb)
	ss.AddFilePath(path.Join(dir, "test"))

	time.AfterFunc(time.Second * 5, func() {
		_, err := os.Create(filepath.Join(dir, "test", "testfile.jpg"))
		assert.Nil(t, err)
		time.AfterFunc(time.Second * 2, func() {
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
