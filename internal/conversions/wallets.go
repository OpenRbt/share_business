package conversions

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
	"washBonus/openapi/models"

	"github.com/go-openapi/strfmt"
)

func WalletFromDB(model dbmodels.Wallet) entity.Wallet {
	return entity.Wallet{
		ID:             model.ID,
		UserID:         model.UserID,
		OrganizationID: model.OrganizationID,
		Balance:        model.Balance,
	}
}

func WalletToRest(e entity.Wallet) *models.Wallet {
	return &models.Wallet{
		ID:             strfmt.UUID(e.ID.String()),
		UserID:         e.UserID,
		OrganizationID: strfmt.UUID(e.OrganizationID.String()),
		Balance:        e.Balance.IntPart(),
		PendingBalance: e.PendingBalance.IntPart(),
	}
}

func WalletsFromDB(wallets []dbmodels.Wallet) []entity.Wallet {
	res := make([]entity.Wallet, len(wallets))

	for i, value := range wallets {
		res[i] = WalletFromDB(value)
	}

	return res
}

func WalletsToRest(wallets []entity.Wallet) []*models.Wallet {
	res := make([]*models.Wallet, len(wallets))

	for i, value := range wallets {
		res[i] = WalletToRest(value)
	}

	return res
}
