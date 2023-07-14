package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"smart-chat/internal/dto"
	"smart-chat/internal/model"
	"strings"
)

type DeepAiProvider struct {
	config *model.Config
}

func NewDeepAiProvider(cfg *model.Config) *DeepAiProvider {
	return &DeepAiProvider{
		config: cfg,
	}
}

func (t *DeepAiProvider) GetConnection(ask *dto.Ask) (*dto.Answer, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", t.config.DeepAi.URL, strings.NewReader("text="+ask.Text))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	req.Header.Set("api-key", t.config.DeepAi.ApiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	asnwer := dto.Answer{}

	if err = json.Unmarshal(body, &asnwer); err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	return &asnwer, nil
}
