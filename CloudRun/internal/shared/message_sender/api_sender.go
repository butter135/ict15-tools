package message_sender

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type APISender struct {
	BaseURL string
}

func NewAPISender(baseURL string) *APISender {
	return &APISender{BaseURL: baseURL}
}

func (a *APISender) Send(p Payload) error {
	body, _ := json.Marshal(map[string]interface{}{
		"message":   p.Message,
		"channelId": p.ChannelID,
	})

	req, err := http.NewRequest("POST", a.BaseURL+"/messages", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+p.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
