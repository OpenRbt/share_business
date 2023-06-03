package conversions

import (
	"github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	"wash_admin/internal/app"
	"wash_admin/internal/dal/dbmodels"
)

func PendingMessagesFromDb(dbMsgs []dbmodels.PendingMessage) []app.PendingMessage {
	res := make([]app.PendingMessage, len(dbMsgs))
	for i, db := range dbMsgs {
		res[i] = PendingMessageFromDB(db)
	}
	return res
}

func PendingMessageFromDB(db dbmodels.PendingMessage) app.PendingMessage {
	return app.PendingMessage{
		ID:          app.MessageID(db.ID),
		MessageType: vo.MessageType(db.MessageType),
		Payload:     db.Payload,
		CreatedAt:   db.CreatedAt,
		IsSent:      db.IsSent,
		SentAt:      db.SentAt,
	}
}
