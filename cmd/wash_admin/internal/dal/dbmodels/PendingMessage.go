package dbmodels

import "time"

type PendingMessage struct {
	ID          int64      `db:"id"`
	MessageType string     `db:"message_type"`
	Payload     []byte     `db:"payload"`
	CreatedAt   time.Time  `db:"created_at"`
	IsSent      bool       `db:"is_sent"`
	SentAt      *time.Time `db:"sent_at"`
}
