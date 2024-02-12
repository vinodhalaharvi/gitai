package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vinodhalaharvi/gitai/openai"
	"io"
	"net/http"
	"os"
)

func CallOpenAI(systemPrompt, userPrompt string) ([]byte, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable not set")
	}

	// Prepare the OpenAIRequest
	body := openai.OpenAIRequest{
		Model: "gpt-4-0125-preview",
		Messages: []openai.OpenAIRoleContent{
			openai.OpenAIRoleContent{
				Role:    "system",
				Content: systemPrompt,
			},
			openai.OpenAIRoleContent{
				Role:    "user",
				Content: userPrompt,
			},
		},
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Make the request to OpenAI
	url := "https://api.openai.com/v1/chat/completions"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the response
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBytes, nil
}
