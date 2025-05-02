package picChan

import (
	"fmt"
)

func CreateMessage(emojiName, userName, token, name2IdUrl, emojiBaseURL string) (string, error) {
	id, err := Name2Id(emojiName, name2IdUrl, token)
	if err != nil {
		return "", err
	}
	message := fmt.Sprintf("![%s](%s%s/image) \nfrom: %s", emojiName, emojiBaseURL, id, userName)
	return message, nil
}
