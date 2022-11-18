package dal

import (
	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

type Storage struct {
	db *dbr.Connection
	l  *zap.SugaredLogger
}

func New(db *dbr.Connection, logger *zap.SugaredLogger) *Storage {
	return &Storage{
		db: db,
		l:  logger,
	}
}
