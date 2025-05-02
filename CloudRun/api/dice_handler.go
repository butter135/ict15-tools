package api

import (
	"net/http"

	"github.com/butter135/ict15-tools/CloudRun/internal/shared/config"
	"github.com/butter135/ict15-tools/CloudRun/internal/shared/message_sender"

	"github.com/butter135/ict15-tools/CloudRun/internal/dice"
)

func HandleDiceRoll(w http.ResponseWriter, r *http.Request){
	if r.FormValue("token") != config.DiceVerifyToken{
		http.Error(w, "Unautorized", http.StatusUnauthorized)
		return
	}

	message, err := dice.PlayDice(r.FormValue("text"))
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