package dbmodels

type AddWashServer struct {
	Name        string `db:"name"`
	Description string `db:"description"`
}
