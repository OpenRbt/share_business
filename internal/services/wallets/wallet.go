package wallets

import (
	"washBonus/internal/app"

	"go.uber.org/zap"
)

type walletService struct {
	logger     *zap.SugaredLogger
	walletRepo app.WalletRepo
}

func New(l *zap.SugaredLogger, repo app.WalletRepo) app.WalletService {
	return &walletService{
		logger:     l,
		walletRepo: repo,
	}
}
