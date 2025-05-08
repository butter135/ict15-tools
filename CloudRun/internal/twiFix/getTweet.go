package twiFix

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func GetTweet(tweetURL string) (string, string, string, error) {
	var twiName, twiDesc, twiImage string

	parsedURL, err := url.Parse(tweetURL)
	if err != nil {
		return "", "", "", fmt.Errorf("invalid tweet URL: %w", err)
	}
	parsedURL.Host = "fxtwitter.com"

	resp, err := http.Get(parsedURL.String())
	if err != nil {
		return "", "", "", fmt.Errorf("failed to fetch tweet: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", "", "", fmt.Errorf("failed to parse HTML: %w", err)
	}

	// 抽出
	twiName, _ = doc.Find("meta[property='twitter:title']").Attr("content")
	twiDesc, _ = doc.Find("meta[property='og:description']").Attr("content")
	twiImage, _ = doc.Find("meta[property='og:image']").Attr("content")

	if twiName == "" && twiDesc == "" && twiImage == "" {
		return "", "", "", errors.New("failed to extract tweet metadata")
	}

	return twiName, twiDesc, twiImage, nil
}
