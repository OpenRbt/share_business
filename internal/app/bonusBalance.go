package app

type BonusBalance struct {
	ID      string
	UserId  string
	Balance float64
}

func (a *app) GetBonusBalance(id string) (*BonusBalance, error) {
	return a.repo.GetBonusBalance(id)
}

func (a *app) AddBonusBalance(balance float64, userId string) (*BonusBalance, error) {
	return a.repo.AddBonusBalance(balance, userId)
}

func (a *app) EditBonusBalance(Id string, balance float64) error {
	return a.repo.EditBonusBalance(Id, balance)
}

func (a *app) DeleteBonusBalance(id string, userId string) error {
	return a.repo.DeleteBonusBalance(id, userId)
}
