package rabbitEntities

type AdminUser struct {
	ID             string  `json:"id"`
	OrganizationID *string `json:"organizationId,omitempty"`
	Email          string  `json:"email"`
	Name           string  `json:"name"`
	Role           string  `json:"role"`
	Version        int     `json:"version"`
}

type UserCreation struct {
	ID            string `json:"id"`
	ServiceKey    string `json:"service_key"`
	Exchange      string `json:"exchange,omitempty"`
	ReadExchange  string `json:"read_exchange,omitempty"`
	WriteExchange string `json:"write_exchange,omitempty"`
}
