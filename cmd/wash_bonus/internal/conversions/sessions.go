package conversions

import (
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
	models2 "wash_bonus/internal/infrastructure/rabbit/models"
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

func SessionToRabbit(e entity.Session) models2.NewSession {
	return models2.NewSession{
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

func SessionUserAssign(sessionID uuid.UUID, user entity.User) models2.SessionUserAssign {
	return models2.SessionUserAssign{
		SessionID: sessionID.String(),
		UserID:    user.ID,
	}
}
