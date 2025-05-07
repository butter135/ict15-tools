package api

import (
	"net/http"
	"strings"

	"github.com/butter135/ict15-tools/CloudRun/internal/shared/config"
	"github.com/butter135/ict15-tools/CloudRun/internal/shared/message_sender"

)

func HandleTokumei(w http.ResponseWriter, r *http.Request) {
	if (r.FormValue("token") != config.TokumeiVerifyToken1 && r.FormValue("token") != config.TokumeiVerifyToken2){
		http.Error(w, "Unautorized", http.StatusUnauthorized)
		return
	}

	text := strings.TrimSpace(r.FormValue("text"))
	if text == "" {
		http.Error(w, "Message text is required", http.StatusBadRequest)
		return
	}


	payload := message_sender.Payload{
		Message: text,
		ChannelID: config.TokumeiChannelID,
		AuthToken: config.TokumeiBotToken,
	}

	err := message_sender.NewAPISender(config.ApiBaseUrl).Send(payload)

	if err != nil{
		http.Error(w, "failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}