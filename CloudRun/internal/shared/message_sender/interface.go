package message_sender

type Payload struct {
	Message string
	WebhookURL string
	ChannelID string
	AuthToken string
}

type Sender interface {
	Send(payload Payload) error
}