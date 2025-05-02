package picChan

import (
	"strings"
	"errors"
)

func NormalizeEmojiName(raw string) (string, error) {
	clean := strings.ReplaceAll(raw, ":", "")
	if clean == "" {
		return "", errors.New("empty emoji name after normalization")
	}
	if strings.ContainsAny(clean, " \t") {
		return "", errors.New("emoji name must not contain whitespace")
	}
	return clean, nil
}