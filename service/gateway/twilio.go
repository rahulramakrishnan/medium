package gateway

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type (
	// TwilioGateway is a gateway to the Twilio API.
	TwilioGateway interface {
		SendSMS(mobile, msg string) (*http.Response, error)
	}
	twilioGateway struct {
		client    *http.Client
		accountID string
		authToken string
	}
)

// NewTwilioGateway returns a new twilio gateway instance.
func NewTwilioGateway() TwilioGateway {
	return &twilioGateway{
		client:    &http.Client{},
		accountID: os.Getenv("TWILIO_ACCOUNT_ID"),
		authToken: os.Getenv("TWILIO_AUTH_TOKEN"),
	}
}

func (t *twilioGateway) SendSMS(mobile, msg string) (resp *http.Response, err error) {
	endpoint := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", t.accountID)
	values := url.Values{
		"From": []string{"+16503530259"},
		"Body": []string{msg},
		"To":   []string{mobile},
	}
	req, err := http.NewRequest("POST", endpoint, strings.NewReader(values.Encode()))
	req.SetBasicAuth(t.accountID, t.authToken)
	resp, err = t.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
