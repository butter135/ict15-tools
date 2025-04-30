package message_sender

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type WebhookSender struct {

}

func NewWebhookSeder() *WebhookSender {
	return &WebhookSender{}
}

func (w *WebhookSender)Send(payload Payload) error{
	body, _ := json.Marshal(map[string]string{
		"text": payload.Message,
	})
	req, err := http.NewRequest("POST", payload.WebhookURL, bytes.NewBuffer(body))
	if err != nil{
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil{
		return err
	}

	defer resp.Body.Close()

	return nil
}