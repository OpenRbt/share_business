package dbmodels

type RegisterWashServer struct {
	Title       string `db:"title"`
	Description string `db:"description"`
	ServiceKey  string `db:"service_key"`
	CreatedBy   string `db:"created_by"`
}
