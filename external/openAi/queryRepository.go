package openAi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
)

const (
	openAiUrl = "https://api.openai.com"
)

type ContentResponse struct {
	Content map[string]string `json:"content"`
}

func GenerateContent(title string) (string, error) {
	openAiToken := viper.GetString("OPEN_AI_TOKEN")
	url := "/v1/chat/completions"
	requestBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"response_format": map[string]string{
			"type": "json_object",
		},
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": "Generate article content about " + title + ".",
			},
		},
	}

	requestBodyJSON, err := json.Marshal(requestBody)

	r, err := http.NewRequest("POST", openAiUrl+url, bytes.NewBuffer(requestBodyJSON))

	if err != nil {
		return "", err
	}

	r.Header.Add("Authorization", "Bearer "+openAiToken)

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)

	if err != nil {
		return "", err
	}

	var parsedBody ContentResponse

	err = json.Unmarshal(body, &parsedBody)

	if err != nil {
		return "", err
	}

	fmt.Println(parsedBody)

	return "test", nil
}
