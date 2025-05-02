package message_sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type APISender struct {
	BaseURL string
}

func NewAPISender(baseURL string) *APISender {
	return &APISender{BaseURL: baseURL}
}

func (s *APISender) Send(p Payload) error {
	if p.ChannelID == "" || p.AuthToken == "" {
		return fmt.Errorf("channel ID and auth token are required")
	}

	body, err := json.Marshal(map[string]string{
		"channel_id": p.ChannelID,
		"message":    p.Message,
	})
	if err != nil {
		return fmt.Errorf("failed to marshal api payload: %w", err)
	}

	req, err := http.NewRequest("POST", s.BaseURL+"posts", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create api request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+p.AuthToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("api send failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to send message: %s (%s)", resp.Status, string(respBody))
	}

	return nil
}
