package gomysmtp

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// sendURL V1 send URL
const sendURL = "https://app.ohmysmtp.com/api/v1/send"

// Client Handles sending of emails
type Client struct {
	Token string
}

// Send Make a HTTP request to ohmysmtp API
func (c *Client) Send(p Payload) error {
	pb, err := json.Marshal(p)
	if err != nil {
		return err
	}

	b := bytes.NewReader(pb)

	req, err := http.NewRequest("POST", sendURL, b)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Ohmysmtp-Server-Token", c.Token)

	resp, err := http.DefaultClient.Do(req)
	log.Println(resp)

	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusForbidden:
		return errors.New("403 Forbidden. Check token")
	}

	defer resp.Body.Close()

	return nil
}

// Payload Content of email
type Payload struct {
	From        string `json:"from"`
	To          string `json:"to"`
	CC          string `json:"cc"`
	BCC         string `json:"bcc"`
	Replyto     string `json:"replyto"`
	Subject     string `json:"subject"`
	Textbody    string `json:"textbody"`
	HTMLbody    string `json:"htmlbody"`
	Attachments string `json:"attachments"`
}

// NewClient Creates a new client
func NewClient(t string) *Client {
	return &Client{Token: t}
}
