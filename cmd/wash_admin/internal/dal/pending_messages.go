package dal

import (
	"time"
	"wash_admin/internal/app"
	"wash_admin/internal/conversions"
	"wash_admin/internal/dal/dbmodels"
)

func (s *Storage) SavePendingMessage(msg app.PendingMessage) error {
	createdAt := time.Now().UTC()

	bytes, ok := msg.Payload.([]byte)
	if !ok {
		return ErrBadPayload
	}
	_, err := s.db.NewSession(nil).
		InsertInto("rabbit_send_log").
		Columns("message_type", "payload", "created_at").
		Values(msg.MessageType, bytes, createdAt).
		Exec()

	return err
}

func (s *Storage) GetPendingMessages(lastMessageID int64) ([]app.PendingMessage, error) {
	var dbMessages []dbmodels.PendingMessage

	count, err := s.db.NewSession(nil).
		Select("*").
		From("rabbit_send_log").
		Where("sent = false AND id > ?", lastMessageID).
		Load(&dbMessages)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, app.ErrNotFound
	}

	return conversions.PendingMessagesFromDb(dbMessages), nil
}

func (s *Storage) MarkMessageAsSent(messageID int64) error {
	sentAt := time.Now().UTC()
	_, err := s.db.NewSession(nil).
		Update("rabbit_send_log").
		Set("sent", true).
		Set("sent_at", sentAt).
		Where("id = ?", messageID).
		Exec()

	return err
}
