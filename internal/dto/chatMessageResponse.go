package dto

type ChatMessageResponse struct {
	Answer string `json:"answer"`
}

func (c *ChatMessageResponse) ParseFromChatMessageResponse(answer *Answer) {
	c.Answer = answer.Output
}
