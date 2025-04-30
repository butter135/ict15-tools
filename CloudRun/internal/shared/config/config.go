package config

import "os"

var (
	DiceVerifyToken string // 受信時の検証用トークン
	DiceWebhookURL  string // 送信用Webhookのエンドポイント
)

func Init(){
	DiceVerifyToken = mustGetenv("DICE_VERIFY_TOKEN")
	DiceWebhookURL = mustGetenv("DICE_WEBHOOK_URL")
}

func mustGetenv(key string) string{
	val := os.Getenv(key)
	if val == ""{
		panic("missing required env: " + key)
	}
	return val
}