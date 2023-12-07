package conversions

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
	rabbitEntities "washbonus/internal/infrastructure/rabbit/entities"
	"washbonus/openapi/bonus/models"

	"github.com/go-openapi/strfmt"
	"github.com/shopspring/decimal"

	uuid "github.com/satori/go.uuid"
)

func SessionFromDB(db dbmodels.Session) entities.Session {
	var user *entities.User

	if db.User != nil {
		user = &entities.User{
			ID: *db.User,
		}
	}

	return entities.Session{
		ID:         db.ID.UUID,
		User:       user,
		Post:       db.PostID,
		WashServer: entities.WashServer{ID: db.WashServer.UUID},
		Finished:   db.Finished,
	}
}

func SessionToRabbit(e entities.Session) rabbitEntities.NewSession {
	return rabbitEntities.NewSession{
		ID:     e.ID.String(),
		PostID: e.Post,
	}
}

func SessionToRest(e entities.Session) *models.Session {
	washServer := models.WashServer{
		ID:             e.WashServer.ID.String(),
		Name:           e.WashServer.Title,
		Description:    e.WashServer.Description,
		OrganizationID: strfmt.UUID(e.WashServer.OrganizationID.String()),
		GroupID:        strfmt.UUID(e.WashServer.GroupID.String()),
	}
	return &models.Session{
		PostID:      e.Post,
		PostBalance: 0, //TODO: add post balance field
		WashServer:  &washServer,
	}
}

func SessionUserAssign(sessionID uuid.UUID, userID string, post int64) rabbitEntities.AssignUserToSession {
	return rabbitEntities.AssignUserToSession{
		SessionID: sessionID.String(),
		UserID:    userID,
		Post:      post,
	}
}

func SessionBonusCharge(sessionID uuid.UUID, amount decimal.Decimal, post int64) rabbitEntities.BonusCharge {
	return rabbitEntities.BonusCharge{
		SessionID: sessionID.String(),
		Amount:    amount.IntPart(),
		Post:      post,
	}
}
