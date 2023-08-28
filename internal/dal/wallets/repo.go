package wallets

import (
	"washBonus/internal/app"

	"github.com/gocraft/dbr/v2"
	"go.uber.org/zap"
)

type walletRepo struct {
	l  *zap.SugaredLogger
	db *dbr.Connection
}

func NewRepo(l *zap.SugaredLogger, db *dbr.Connection) app.WalletRepo {
	return &walletRepo{
		l:  l,
		db: db,
	}
}
