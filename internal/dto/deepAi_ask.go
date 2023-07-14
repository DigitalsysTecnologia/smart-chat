package dto

type Ask struct {
	Text string `json:"text"`
}

func (a *Ask) ParseFromChatMessageRequest(chatMessageRequest *ChatMessageRequest) {
	a.Text = chatMessageRequest.Question
}
