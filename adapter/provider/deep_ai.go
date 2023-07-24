package provider

import (
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
	"smart-chat/internal/dto"
	"smart-chat/internal/model"
	"strings"
	"time"
)

type DeepAiProvider struct {
	config *model.Config
	logger *zap.Logger
}

func NewDeepAiProvider(cfg *model.Config, logger *zap.Logger) *DeepAiProvider {
	return &DeepAiProvider{
		config: cfg,
		logger: logger,
	}
}

func (t *DeepAiProvider) CallIA(ctx context.Context, text string) (*dto.Answer, error) {
	requestID := ctx.Value("requestID").(string)

	t.logger.Debug("CallIA",
		zap.String("requestID", requestID),
		zap.String("phase", "Provider"))

	client := &http.Client{}
	answer := dto.Answer{}

	req, err := http.NewRequest("POST", t.config.DeepAi.URL, strings.NewReader("text="+text))
	if err != nil {
		t.logger.Error("Error constructing request",
			zap.String("requestID", requestID),
			zap.Error(err),
			zap.String("phase", "Provider"))
		return nil, err
	}

	req.Header.Set("api-key", t.config.DeepAi.ApiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	t.logger.Debug("Request constructed",
		zap.String("requestID", requestID),
		zap.String("phase", "Provider"))
	zap.Any("request", req)

	answer.QuestionDate = time.Now()
	resp, err := client.Do(req)
	if err != nil {
		t.logger.Error("Error calling IA",
			zap.String("requestID", requestID),
			zap.Error(err),
			zap.String("phase", "Provider"))
		return nil, err
	}
	answer.ResponseDate = time.Now()

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.logger.Error("Error reading response",
			zap.String("requestID", requestID),
			zap.Error(err),
			zap.String("phase", "Provider"))

		return nil, err
	}

	t.logger.Debug("Response read",
		zap.String("requestID", requestID),
		zap.String("phase", "Provider"))

	if err = json.Unmarshal(body, &answer); err != nil {
		t.logger.Error("Error unmarshalling response",
			zap.String("requestID", requestID),
			zap.Error(err),
			zap.String("phase", "Provider"))

		return nil, err
	}

	t.logger.Debug("Response unmarshalled",
		zap.String("requestID", requestID),
		zap.String("phase", "Provider"))
	zap.Any("response", answer)

	return &answer, nil
}
