package conversions

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
	"washbonus/openapi/bonus/models"

	"github.com/go-openapi/strfmt"
)

func WalletFromDB(model dbmodels.Wallet) entities.Wallet {
	return entities.Wallet{
		ID:     model.ID,
		UserID: model.UserID,
		Organization: entities.OrganizationForWallet{
			ID:   model.OrganizationID,
			Name: model.OrganizationName,
		},
		Balance: model.Balance,
	}
}

func WalletToRest(e entities.Wallet) *models.Wallet {
	id := e.ID.String()
	organizationID := e.Organization.ID.String()
	balance := e.Balance.IntPart()
	pendingBalance := e.PendingBalance.IntPart()

	organization := models.Organization{
		ID:   (*strfmt.UUID)(&organizationID),
		Name: &e.Organization.Name,
	}

	return &models.Wallet{
		ID:             (*strfmt.UUID)(&id),
		UserID:         &e.UserID,
		Organization:   &organization,
		Balance:        &balance,
		PendingBalance: &pendingBalance,
	}
}

func WalletsFromDB(wallets []dbmodels.Wallet) []entities.Wallet {
	res := make([]entities.Wallet, len(wallets))

	for i, value := range wallets {
		res[i] = WalletFromDB(value)
	}

	return res
}

func WalletsToRest(wallets []entities.Wallet) []*models.Wallet {
	res := make([]*models.Wallet, len(wallets))

	for i, value := range wallets {
		res[i] = WalletToRest(value)
	}

	return res
}
