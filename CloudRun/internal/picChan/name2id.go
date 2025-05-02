package picChan

import (
	"io"
	"fmt"
	"net/http"
	"encoding/json"
)

func Name2Id(name, name2IdUrl, token string) (string, error) {
	req, _ := http.NewRequest("GET", name2IdUrl + name, nil)
	req.Header.Set("Authorization", "Bearer " + token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return "", err
	}

	id, ok := jsonData["id"].(string)
	if !ok {
		return "", fmt.Errorf("emoji id not found in response")
	}
	return id, nil
}

