package conversions

import (
	"wash_admin/internal/app"
	"wash_admin/internal/dal/dbmodels"
)

func WashServerFromDB(dbWashServer dbmodels.WashServer) app.WashServer {
	return app.WashServer{
		ID:          dbWashServer.ID.UUID,
		Title:       dbWashServer.Title,
		Description: dbWashServer.Description,
		Owner:       dbWashServer.Owner.UUID,
		ServiceKey:  dbWashServer.ServiceKey,
	}
}
