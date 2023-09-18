package controllers

import (
	"context"
	"washbonus/internal/app"
	"washbonus/internal/entities"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
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

func (ctrl *walletController) Get(ctx context.Context, auth app.Auth, pagination entities.Pagination) ([]entities.Wallet, error) {
	wallets, err := ctrl.walletSvc.Get(ctx, auth.User.ID, pagination)
	if err != nil {
		return nil, err
	}

	return ctrl.addPendingBalancesToWallets(ctx, auth.User.ID, wallets)
}

func (ctrl *walletController) GetByOrganizationId(ctx context.Context, auth app.Auth, organizationID uuid.UUID) (entities.Wallet, error) {
	wallet, err := ctrl.walletSvc.GetOrCreate(ctx, auth.User.ID, organizationID)
	if err != nil {
		return entities.Wallet{}, err
	}

	return ctrl.addPendingBalanceToWallet(ctx, auth.User.ID, wallet)
}

func (ctrl *walletController) addPendingBalancesToWallets(ctx context.Context, userID string, wallets []entities.Wallet) ([]entities.Wallet, error) {
	pendingBalances, err := ctrl.sessionSvc.GetUserPendingBalances(ctx, userID)
	if err != nil {
		return nil, err
	}

	orgBalances := make(map[string]decimal.Decimal, len(pendingBalances))

	for _, balance := range pendingBalances {
		orgBalances[balance.OrganizationID.String()] = balance.PendingBalance
	}

	for i := range wallets {
		wallets[i].PendingBalance = orgBalances[wallets[i].Organization.ID.String()]
	}
	return wallets, nil
}

func (ctrl *walletController) addPendingBalanceToWallet(ctx context.Context, userID string, wallet entities.Wallet) (entities.Wallet, error) {
	pendingBalance, err := ctrl.sessionSvc.GetUserPendingBalanceByOrganization(ctx, userID, wallet.Organization.ID)
	if err != nil {
		return wallet, err
	}
	wallet.PendingBalance = pendingBalance
	return wallet, nil
}
