package app

import (
	"github.com/OpenRbt/share_business/wash_rabbit/entity/admin"
	uuid "github.com/satori/go.uuid"
)

type WashServer struct {
	ID          uuid.UUID
	Title       string
	Description string
	ServiceKey  string
	Owner       uuid.UUID
}

type RegisterWashServer struct {
	Title       string
	Description string
}

type UpdateWashServer struct {
	ID          uuid.UUID
	Title       *string
	Description *string
}

func WashServerToRabbit(e WashServer) admin.ServerRegistered {
	return admin.ServerRegistered{
		ID:          e.ID.String(),
		Title:       e.Title,
		Description: e.Description,
	}
}

func WashServerUpdateToRabbit(e UpdateWashServer, deleted bool) admin.ServerUpdate {
	var del *bool
	if deleted {
		t := true
		del = &t
	}

	return admin.ServerUpdate{
		ID:          e.ID.String(),
		Title:       e.Title,
		Description: e.Description,
		Deleted:     del,
	}
}
