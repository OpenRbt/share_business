package bonusreport

import (
	"washbonus/internal/app"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

type repo struct {
	l  *zap.SugaredLogger
	db *dbr.Connection
}

func NewRepo(l *zap.SugaredLogger, db *dbr.Connection) app.BonusReportRepo {
	return &repo{
		l:  l,
		db: db,
	}
}
