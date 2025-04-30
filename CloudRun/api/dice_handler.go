package api

import (
	"encoding/json"
	"github.com/butter135/ict15-tools/CloudRun/internal/shared/message_sender"
	"net/http"

	"github.com/butter135/ict15-tools/CloudRun/internal/shared/config"
	"github.com/butter135/ict15-tools/CloudRun/internal/dice"
)

type DiceRequest struct{
	Text string `json:"text"`
	Token string `json:"token"`
}

func HandleDiceRoll(w http.ResponseWriter, r *http.Request){
	var req DiceRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Token != config.DiceVerifyToken{
		http.Error(w, "Unautorized", http.StatusUnauthorized)
	}

	message, err := dice.PlayDice(req.Text)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	payload := message_sender.Payload{
		Message: message,
		WebhookURL: config.DiceWebhookURL,
	}

	err = message_sender.NewWebhookSeder().Send(payload)

	if err != nil{
		http.Error(w, "failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}