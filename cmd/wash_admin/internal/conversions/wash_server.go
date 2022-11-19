package conversions

import (
	"wash_admin/internal/entity"
	"wash_admin/openapi/models"
)

func WashServerToRest(wash_server entity.WashServer) models.WashServer {
	return models.WashServer{
		ID: wash_server.ID.String(),
		Name: wash_server.Name,
		Description: wash_server.Description,
		APIKey: wash_server.APIKey,
	}
}