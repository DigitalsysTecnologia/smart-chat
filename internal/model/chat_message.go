package model

import (
	"smart-chat/internal/dto"
	"time"
)

type ChatMessage struct {
	ID           int64     `json:"id"`
	UserID       string    `json:"user_id"`
	Question     string    `json:"question"`
	ResponseID   string    `json:"response_id"`
	Response     string    `json:"response"`
	QuestionDate time.Time `json:"question_date"`
	ResponseDate time.Time `json:"response_date"`
	ChatID       uint64    `json:"chat_id"`
	Chat         Chat      `json:"chat"`
}

func (c *ChatMessage) ParseFromChatMessageRequestAndAnswer(chatMessageRequest *dto.ChatMessageRequest, answer *dto.Answer) {
	c.UserID = chatMessageRequest.UserID
	c.Question = chatMessageRequest.Question
	c.ResponseID = answer.ID
	c.Response = answer.Output
	c.QuestionDate = time.Now()
	c.ResponseDate = time.Now()
	c.ChatID = chatMessageRequest.ChatID

}
