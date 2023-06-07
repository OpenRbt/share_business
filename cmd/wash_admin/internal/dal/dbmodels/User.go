package dbmodels

type User struct {
	ID   string `db:"id"`
	Role string `db:"role"`
}

const (
	AdminRole    string = "admin"
	UserRole     string = "user"
	EngineerRole string = "engineer"
)

type UpdateUser struct {
	ID   string `db:"id"`
	Role string `db:"role"`
}
