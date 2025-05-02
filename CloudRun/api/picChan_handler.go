package api

import (
	"net/http"

	"github.com/butter135/ict15-tools/CloudRun/internal/shared/config"
	"github.com/butter135/ict15-tools/CloudRun/internal/shared/message_sender"

	"github.com/butter135/ict15-tools/CloudRun/internal/picChan"
)

func HandlePicChan(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("token") != config.PicChanVerifyToken{
		http.Error(w, "Unautorized", http.StatusUnauthorized)
		return
	}
	text, err := picChan.NormalizeEmojiName(r.FormValue("text"))
	if err != nil{
		http.Error(w, "invalid emoji format", http.StatusBadRequest)
		return
	}

	message, err :=picChan.CreateMessage(
		text,
		r.FormValue("user_name"),
		config.PicChanBotToken,
		config.Name2IdUrl,
		config.EmojiBaseUrl)

	if err != nil {
		http.Error(w, "failed to create message", http.StatusInternalServerError)
		return
	}

	payload := message_sender.Payload{
		Message: message,
		ChannelID: r.FormValue("channel_id"),
		AuthToken: config.PicChanBotToken,
	}

	err = message_sender.NewAPISender(config.ApiBaseUrl).Send(payload)

	if err != nil{
		http.Error(w, "failed to send message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}