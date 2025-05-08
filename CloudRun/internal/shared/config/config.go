package config

import "os"

var (
	ApiBaseUrl string
	EmojiBaseUrl string
	Name2IdUrl string

	DiceVerifyToken string // 受信時の検証用トークン
	DiceWebhookURL  string // 送信用Webhookのエンドポイント

	PicChanVerifyToken string
	PicChanBotToken string

	TokumeiVerifyToken1 string
	TokumeiVerifyToken2 string
	TokumeiChannelID string
	TokumeiBotToken string

	TwiFixVerifyToken1 string
	TwiFixVerifyToken2 string
	TwiFixBotToken string
)

func Init(){
	ApiBaseUrl = mustGetenv("API_BASE_URL")
	EmojiBaseUrl = mustGetenv("EMOJI_BASE_URL")
	Name2IdUrl = mustGetenv("NAME_TO_ID_URL")
	DiceVerifyToken = mustGetenv("DICE_VERIFY_TOKEN")
	DiceWebhookURL = mustGetenv("DICE_WEBHOOK_URL")
	PicChanVerifyToken = mustGetenv("PICCHAN_VERIFY_TOKEN")
	PicChanBotToken = mustGetenv("PICCHAN_BOT_TOKEN")
	TokumeiVerifyToken1 = mustGetenv("TOKUMEI_VERIFY_TOKEN1")
	TokumeiVerifyToken2 = mustGetenv("TOKUMEI_VERIFY_TOKEN2")
	TokumeiChannelID = mustGetenv("TOKUMEI_CHANNEL_ID")
	TokumeiBotToken = mustGetenv("TOKUMEI_BOT_TOKEN")
	TwiFixVerifyToken1 = mustGetenv("TWIFIX_VERIFY_TOKEN1")
	TwiFixVerifyToken2 = mustGetenv("TWIFIX_VERIFY_TOKEN2")
	TwiFixBotToken = mustGetenv("TWIFIX_BOT_TOKEN")
}

func mustGetenv(key string) string{
	val := os.Getenv(key)
	if val == ""{
		panic("missing required env: " + key)
	}
	return val
}