package rabbitEntities

type AdminUser struct {
	ID             string  `json:"id"`
	OrganizationID *string `json:"organizationId,omitempty"`
	Email          string  `json:"email"`
	Name           string  `json:"name"`
	Role           string  `json:"role"`
	Version        int     `json:"version"`
}

type CreateUser struct {
	ID         string `json:"id"`
	ServiceKey string `json:"service_key"`
	Exchange   string `json:"exchange"`
}
