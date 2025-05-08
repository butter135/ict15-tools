package twiFix

import (
	"fmt"
)

func CreateMessage(tweetURL, username string) (string, error) {
	name, desc, image, err := GetTweet(tweetURL)

	if err != nil {
		return "", err
	}
	message := fmt.Sprintf(
		"**%s**\n%s\n[![image](%s)](%s)\n[ツイートを見る](%s)\nfrom: %s",
		name,
		desc,
		image,
		tweetURL,
		tweetURL,
		username,
	)
	return message, nil
}