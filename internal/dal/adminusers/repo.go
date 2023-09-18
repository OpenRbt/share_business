package adminusers

import (
	"washbonus/internal/app"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

type adminUserRepo struct {
	l  *zap.SugaredLogger
	db *dbr.Connection
}

func NewRepo(l *zap.SugaredLogger, db *dbr.Connection) app.AdminRepo {
	return &adminUserRepo{
		l:  l,
		db: db,
	}
}
