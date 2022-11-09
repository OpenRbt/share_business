package dto

import (
	"database/sql"
	"github.com/go-openapi/strfmt"
	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/app/entity/vo"
	"wash-bonus/internal/dal/dbmodel"
	"wash-bonus/openapi/models"
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

func UserToRest(u entity.User) models.User {
	return models.User{
		Active:     u.Active,
		CreatedAt:  (*strfmt.DateTime)(u.CreatedAt),
		FirebaseID: u.IdentityID,
		ID:         u.ID.String(),
		ModifiedAt: (*strfmt.DateTime)(u.ModifiedAt),
	}
}

func UsersToRest(uu []entity.User) []*models.User {
	apiUsers := make([]*models.User, len(uu))

	for i, u := range uu {
		apiUser := UserToRest(u)
		apiUsers[i] = &apiUser
	}
	return apiUsers
}

func UserFromRestAdd(u models.UserAdd) entity.User {
	return entity.User{
		Active: u.Active,
	}
}

func UserFromRestUpdate(u models.UserUpdate) vo.UserUpdate {
	return vo.UserUpdate{
		Active: u.Active,
	}
}
