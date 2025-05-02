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

)

func Init(){
	ApiBaseUrl = mustGetenv("API_BASE_URL")
	EmojiBaseUrl = mustGetenv("EMOJI_BASE_URL")
	Name2IdUrl = mustGetenv("NAME_TO_ID_URL")
	DiceVerifyToken = mustGetenv("DICE_VERIFY_TOKEN")
	DiceWebhookURL = mustGetenv("DICE_WEBHOOK_URL")
	PicChanVerifyToken = mustGetenv("PICCHAN_VERIFY_TOKEN")
	PicChanBotToken = mustGetenv("PICCHAN_BOT_TOKEN")


}

func mustGetenv(key string) string{
	val := os.Getenv(key)
	if val == ""{
		panic("missing required env: " + key)
	}
	return val
}