package gateway

// Gateways holds 3rd party gateways.
type Gateways struct {
	Twilio TwilioGateway
}

// New returns a new Gateways instance.
func New() *Gateways {
	return &Gateways{
		Twilio: &twilioGateway{},
	}
}
