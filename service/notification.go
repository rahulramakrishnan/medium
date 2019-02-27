package service

type (
	Notification interface {
		Notify(username, msg string) error
	}
	notification struct {
		gateway TwilioGateway
	}
)

// NewNotificationService returns a new notification service instance.
func NewNotificationService(gateway TwilioGateway) {
	return &notification{
		gateway: gateway,
	}
}

func (n *notification) Notify(username, msg string) error {
	n.gateway.
}
