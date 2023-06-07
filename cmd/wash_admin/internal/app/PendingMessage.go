package app

import (
	rabbit_vo "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	"time"
)

type MessageID int64
type PendingMessage struct {
	ID          MessageID
	MessageType rabbit_vo.MessageType
	Payload     interface{}
	CreatedAt   time.Time
	IsSent      bool
	SentAt      *time.Time
}
