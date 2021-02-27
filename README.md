# gomysmtp

A simple wrapper around the [ohmysmtp](https://www.ohmysmtp.com) API allowing you to send messages through the service.

## Example
```go
emailClient := gomysmtp.NewClient("OHMYSMTP_TOKEN")
emailPayload := gomysmtp.Payload{
    From:     "service@example.com",
    To:       "user@example.com",
    Subject:  "Welcome",
    Textbody: "Hi",
}

err := emailClient.Send(emailPayload)
if err != nil {
    // handle err
}
```