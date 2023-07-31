package model

type Chat struct {
	ID        uint64 `json:"id" gorm:"primaryKey"`
	UserID    string `json:"user_id" gorm:"column:user_id"`
	CreatedAt string `json:"created_at" gorm:"column:created_at;type:timestamp"`
	UpdatedAt string `json:"updated_at" gorm:"column:updated_at;type:timestamp"`
}

func (Chat) TableName() string {
	return "chat"
}
