package model

import "time"

type MixinMessage struct {
	Action         string    `gorm:"type:varchar(50); not null" json:"action"`
	UserId         int       `gorm:"type:int;not null;index" json:"user_id"`
	ConversationId int       `gorm:"type:int;not null;" json:"conversation_id"`
	Category       string    `gorm:"type:varchar(50); not null" json:"category"`
	State          string    `gorm:"type:varchar(20); not null" json:"state"`
	MessageId      int       `gorm:"type:int;not null;comment:UUID;uniqueIndex" json:"message_id"`
	Content        string    `gorm:"type:varchar(50);comment:decrepted data;" json:"content"`
	Raw            string    `gorm:"type:varchar(50);not null;" json:"raw"`
	ProcessedAt    time.Time `gorm:"type:datetime(3);" json:"processed_at"`
	CreatedAt      time.Time `gorm:"type:datetime(3); not null;" json:"created_at"`
	UpdatedAt      time.Time `gorm:"type:datetime(3);not null;" json:"updated_at"`
}
