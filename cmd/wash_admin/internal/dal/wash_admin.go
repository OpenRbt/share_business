package dal

import (
	"context"
	"errors"
	"wash_admin/internal/conversions"
	"wash_admin/internal/dal/dbmodels"
	"wash_admin/internal/entity"

	"github.com/gocraft/dbr/v2"
)

func (s *Storage) GetWashAdmin(ctx context.Context, identity string) (entity.WashAdmin, error) {
	var dbWashAdmin dbmodels.WashAdmin

	err := s.db.NewSession(nil).
		Select("*").
		From("wash_admins").
		Where("identity = ?", identity).
		LoadOneContext(ctx, &dbWashAdmin)

	switch {
	case err == nil:
		return conversions.WashAdminFromDB(dbWashAdmin), err
	case errors.Is(err, dbr.ErrNotFound):
		return entity.WashAdmin{}, entity.ErrNotFound
	default:
		return entity.WashAdmin{}, err
	}
}
