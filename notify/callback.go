package notify

type Notify struct {
	EventType string
	Path      string
}

type Callback interface {
	Notify(notify Notify)
}
