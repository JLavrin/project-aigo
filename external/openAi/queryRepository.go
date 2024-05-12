package openAi

import (
	"fmt"
	"github.com/JLavrin/project-aigo/util/fetch"
	"github.com/spf13/viper"
	"net/http"
)

const (
	openAiUrl = "https://api.openai.com"
)

type ContentResponse struct {
	Content map[string]string `json:"content"`
}

func GenerateContent(title *string) (string, error) {
	openAiToken := viper.GetString("OPEN_AI_TOKEN")
	fmt.Println("openApiToken", openAiToken)
	url := "/v1/chat/completions"
	requestBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": "Generate content for title: " + *title + ". It should be about 500 words, plain text",
			},
		},
	}

	res, err := fetch.Post(&fetch.Options{
		Method: "POST",
		Url:    openAiUrl + url,
		Headers: http.Header{
			"Authorization": []string{"Bearer " + openAiToken},
			"Content-Type":  []string{"application/json"},
		},
		Body: requestBody,
	})
	if err != nil {
		return "", err
	}

	fmt.Println("res", res.(map[string]interface{})["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string))

	return "", nil
}
