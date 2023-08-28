package organizations

import (
	"washBonus/internal/app"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

type serverGroupRepo struct {
	l  *zap.SugaredLogger
	db *dbr.Connection
}

func NewRepo(l *zap.SugaredLogger, db *dbr.Connection) app.ServerGroupRepo {
	return &serverGroupRepo{
		l:  l,
		db: db,
	}
}
