package handler

// Handlers holds available http handlers.
type Handlers struct {
	Notification *notification
}

// New returns a new instance of Handlers.
func New() *Handlers {
	return &Handlers{
		Notification: &notification{},
	}
}
