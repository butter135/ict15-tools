package api


import (
	"net/http"
	"strings"

	"github.com/butter135/ict15-tools/CloudRun/internal/shared/config"
	"github.com/butter135/ict15-tools/CloudRun/internal/shared/message_sender"

	"github.com/butter135/ict15-tools/CloudRun/internal/twiFix"
)

func HandleTwiFix(w http.ResponseWriter, r *http.Request) {
	if (r.FormValue("token") != config.TwiFixVerifyToken1 && r.FormValue("token") != config.TwiFixVerifyToken2){
		http.Error(w, "Unautorized", http.StatusUnauthorized)
		return
	}

	text := strings.TrimSpace(r.FormValue("text"))
	if text == "" {
		http.Error(w, "Message text is required", http.StatusBadRequest)
		return
	}

	message, err := twiFix.CreateMessage(text, r.FormValue("user_name"))
	if err != nil {
		http.Error(w, "failed to create message", http.StatusInternalServerError)
		return
	}

	payload := message_sender.Payload{
		Message: message,
		ChannelID: r.FormValue("channel_id"),
		AuthToken: config.TwiFixBotToken,
	}

	err = message_sender.NewAPISender(config.ApiBaseUrl).Send(payload)

	if err != nil{
		http.Error(w, "failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}