package config

import "os"

var (
	TokenDice string
)

func Init(){
	TokenDice = mustGetenv("TOKEN_DICE")
}

func mustGetenv(key string) string{
	val := os.Getenv(key)
	if val == ""{
		panic("missing required env: " + key)
	}
	return val
}