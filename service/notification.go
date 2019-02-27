package service

type (
	Notification interface {
		Notify(username, msg string) error
	}
	notification struct {
		Gateway 
	}
)

func (n *notification) Notify(username, msg string) error {

}
