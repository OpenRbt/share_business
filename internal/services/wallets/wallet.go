package wallets

import (
	"washBonus/internal/app"

	"go.uber.org/zap"
)

type walletService struct {
	logger     *zap.SugaredLogger
	walletRepo app.WalletRepo
	orgRepo    app.OrganizationRepo
}

func New(l *zap.SugaredLogger, repo app.WalletRepo, orgRepo app.OrganizationRepo) app.WalletService {
	return &walletService{
		logger:     l,
		walletRepo: repo,
		orgRepo:    orgRepo,
	}
}
