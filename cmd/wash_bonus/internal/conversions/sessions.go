package conversions

import (
	"github.com/OpenRbt/share_business/wash_rabbit/entity/session"
	"github.com/shopspring/decimal"
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
	"wash_bonus/openapi/models"

	uuid "github.com/satori/go.uuid"
)

func SessionFromDB(db dbmodels.Session) entity.Session {
	var user *entity.User

	if db.User != nil {
		user = &entity.User{
			ID: *db.User,
		}
	}

	return entity.Session{
		ID:         db.ID.UUID,
		User:       user,
		Post:       db.PostID,
		WashServer: entity.WashServer{Id: db.WashServer.UUID},
		Finished:   db.Finished,
	}
}

func SessionToRabbit(e entity.Session) session.NewSession {
	return session.NewSession{
		ID:     e.ID.String(),
		PostID: e.Post,
	}
}

func SessionToRest(e entity.Session) *models.Session {
	return &models.Session{
		PostBalance: 0, //TODO: add post balance field
		PostID:      e.Post,
		WashServer:  WashServerToRest(e.WashServer),
	}
}

func SessionUserAssign(sessionID uuid.UUID, userID string) session.UserAssign {
	return session.UserAssign{
		SessionID: sessionID.String(),
		UserID:    userID,
	}
}

func SessionBonusCharge(sessionID uuid.UUID, amount decimal.Decimal) session.BonusCharge {
	return session.BonusCharge{
		SessionID: sessionID.String(),
		Amount:    amount.IntPart(),
	}
}
