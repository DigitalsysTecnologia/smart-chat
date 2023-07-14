package dto

type ChatMessageRequest struct {
	UserID   string `json:"user_id"`
	ChatID   uint64 `json:"chat_id"`
	Question string `json:"question"`
}
