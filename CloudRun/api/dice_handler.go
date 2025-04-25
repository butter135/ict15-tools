package api

import (
	"encoding/json"
	"net/http"

	"github.com/butter135/ict15-tools/CloudRun/internal/shared/config"
	"github.com/butter135/ict15-tools/CloudRun/internal/dice/usecase"
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

	if req.Token != config.TokenDice{
		http.Error(w, "Unautorized", http.StatusUnauthorized)
	}

	message, err := usecase.PlayDice(req.Text)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(message))
}