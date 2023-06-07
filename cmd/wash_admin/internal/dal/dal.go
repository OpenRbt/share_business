package dal

import (
	"errors"
	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

var (
	ErrBadPayload = errors.New("bad message payload")
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
