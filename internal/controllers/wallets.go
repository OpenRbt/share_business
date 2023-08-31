package controllers

import (
	"context"
	"washBonus/internal/app"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type walletController struct {
	logger     *zap.SugaredLogger
	walletSvc  app.WalletService
	sessionSvc app.SessionService
}

func NewWalletController(l *zap.SugaredLogger, walletSvc app.WalletService, sessionSvc app.SessionService) app.WalletController {
	return &walletController{
		logger:     l,
		walletSvc:  walletSvc,
		sessionSvc: sessionSvc,
	}
}

func (ctrl *walletController) Get(ctx context.Context, auth app.Auth, pagination entity.Pagination) ([]entity.Wallet, error) {
	wallets, err := ctrl.walletSvc.Get(ctx, auth.User.ID, pagination)
	if err != nil {
		return nil, err
	}

	return ctrl.addPendingBalancesToWallets(ctx, wallets)
}

func (ctrl *walletController) GetById(ctx context.Context, auth app.Auth, walletID uuid.UUID) (entity.Wallet, error) {
	wallet, err := ctrl.walletSvc.GetById(ctx, walletID)
	if err != nil {
		return entity.Wallet{}, err
	}

	if !app.IsAdmin(auth.User) && wallet.UserID != auth.User.ID {
		return entity.Wallet{}, entity.ErrAccessDenied
	}

	return ctrl.addPendingBalanceToWallet(ctx, auth.User.ID, wallet)
}

func (ctrl *walletController) GetByOrganizationId(ctx context.Context, auth app.Auth, organizationID uuid.UUID) (entity.Wallet, error) {
	wallet, err := ctrl.walletSvc.GetOrCreate(ctx, auth.User.ID, organizationID)
	if err != nil {
		return entity.Wallet{}, err
	}

	return ctrl.addPendingBalanceToWallet(ctx, auth.User.ID, wallet)
}

func (ctrl *walletController) addPendingBalancesToWallets(ctx context.Context, wallets []entity.Wallet) ([]entity.Wallet, error) {
	for i, wallet := range wallets {
		var err error
		wallets[i], err = ctrl.addPendingBalanceToWallet(ctx, wallet.UserID, wallet)
		if err != nil {
			return nil, err
		}
	}
	return wallets, nil
}

func (ctrl *walletController) addPendingBalanceToWallet(ctx context.Context, userID string, wallet entity.Wallet) (entity.Wallet, error) {
	pendingBalance, err := ctrl.sessionSvc.GetUserOrganizationPendingBalance(ctx, userID, wallet.OrganizationID)
	if err != nil {
		return wallet, err
	}
	wallet.PendingBalance = pendingBalance
	return wallet, nil
}
