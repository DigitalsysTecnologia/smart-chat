package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"smart-chat/internal/dto"
	"smart-chat/internal/model"
	"strings"
	"time"
)

type DeepAiProvider struct {
	config *model.Config
	logger *SystemLogger
}

func NewDeepAiProvider(cfg *model.Config, logger *SystemLogger) *DeepAiProvider {
	return &DeepAiProvider{
		config: cfg,
		logger: logger,
	}
}

func (t *DeepAiProvider) CallIA(ctx context.Context, text string) (*dto.Answer, error) {
	requestID := ctx.Value("requestID").(string)

	t.logger.NewLog("CallIA", "requestID", requestID,
		"text", text).
		Debug().
		Phase("Provider").
		Exe()

	client := &http.Client{}
	answer := dto.Answer{}

	req, err := http.NewRequest("POST", t.config.DeepAi.URL, strings.NewReader("text="+text))
	if err != nil {
		t.logger.NewLog("Error creating request", "requestID", requestID).
			Error().
			Description("Error creating request: " + err.Error()).
			Phase("Provider").
			Exe()
		return nil, err
	}

	req.Header.Set("api-key", t.config.DeepAi.ApiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	t.logger.NewLog("constructed request", "requestID", requestID,
		"request", req).
		Debug().
		Phase("Provider").
		Exe()

	answer.QuestionDate = time.Now()
	resp, err := client.Do(req)
	if err != nil {
		t.logger.NewLog("Error calling IA", "requestID", requestID).
			Error().
			Description("Error calling IA: " + err.Error()).
			Phase("Provider").
			Exe()
		return nil, err
	}
	answer.ResponseDate = time.Now()

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.logger.NewLog("Error reading response", "requestID", requestID).
			Error().
			Description("Error reading response: " + err.Error()).
			Phase("Provider").
			Exe()
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	t.logger.NewLog("Response received", "requestID", requestID,
		"response", string(body)).
		Debug().
		Phase("Provider").
		Exe()

	if err = json.Unmarshal(body, &answer); err != nil {
		t.logger.NewLog("Error unmarshalling response", "requestID", requestID).
			Error().
			Description("Error unmarshalling response: " + err.Error()).
			Phase("Provider").
			Exe()
		fmt.Println("Error unmarshalling response:", err)
		return nil, err
	}

	t.logger.NewLog("Response unmarshalled", "requestID", requestID,
		"answer", answer).
		Debug().
		Phase("Provider").
		Exe()

	return &answer, nil
}
