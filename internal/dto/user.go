package dto

import (
	"database/sql"
	"github.com/go-openapi/strfmt"
	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/app/entity/vo"
	"wash-bonus/internal/dal/dbmodel"
	models2 "wash-bonus/internal/transport/rest/restapi/models"
)

func UsersFromDB(uu []dbmodel.User) []entity.User {
	res := make([]entity.User, len(uu))

	for i, u := range uu {
		res[i] = UserFromDB(u)
	}

	return res
}

func UserFromDB(u dbmodel.User) entity.User {
	identityID := ""
	if u.IdentityID.Valid {
		identityID = u.IdentityID.String
	}

	return entity.User{
		Active:     u.Active,
		CreatedAt:  u.CreatedAt,
		ID:         u.ID,
		IdentityID: identityID,
		ModifiedAt: u.ModifiedAt,
	}
}

func UserToDB(u entity.User) dbmodel.User {
	dbIdentityID := sql.NullString{}

	if u.IdentityID != "" {
		dbIdentityID.String = u.IdentityID
		dbIdentityID.Valid = true
	}

	return dbmodel.User{
		Active:     u.Active,
		CreatedAt:  u.CreatedAt,
		ID:         u.ID,
		IdentityID: dbIdentityID,
		ModifiedAt: u.ModifiedAt,
	}
}

func UserToRest(u entity.User) models2.User {
	return models2.User{
		Active:     u.Active,
		CreatedAt:  (*strfmt.DateTime)(u.CreatedAt),
		FirebaseID: u.IdentityID,
		ID:         u.ID.String(),
		ModifiedAt: (*strfmt.DateTime)(u.ModifiedAt),
	}
}

func UsersToRest(uu []entity.User) []*models2.User {
	apiUsers := make([]*models2.User, len(uu))

	for i, u := range uu {
		apiUser := UserToRest(u)
		apiUsers[i] = &apiUser
	}
	return apiUsers
}

func UserFromRestAdd(u models2.UserAdd) entity.User {
	return entity.User{
		Active: u.Active,
	}
}

func UserFromRestUpdate(u models2.UserUpdate) vo.UserUpdate {
	return vo.UserUpdate{
		Active: u.Active,
	}
}
