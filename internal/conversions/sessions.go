package conversions

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
	"washBonus/internal/infrastructure/rabbit/entity/session"
	"washBonus/openapi/models"

	"github.com/shopspring/decimal"

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
		WashServer: entity.WashServer{ID: db.WashServer.UUID},
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

func SessionUserAssign(sessionID uuid.UUID, userID string, post int64) session.UserAssign {
	return session.UserAssign{
		SessionID: sessionID.String(),
		UserID:    userID,
		Post:      post,
	}
}

func SessionBonusCharge(sessionID uuid.UUID, amount decimal.Decimal, post int64) session.BonusCharge {
	return session.BonusCharge{
		SessionID: sessionID.String(),
		Amount:    amount.IntPart(),
		Post:      post,
	}
}
