package rabbitEntities

type ServerGroup struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	OrganizationID string `json:"organizationId"`
	IsDefault      bool   `json:"isDefault"`
	Deleted        bool   `json:"deleted"`
	Version        int    `json:"version"`
	CostPerDay     int64  `json:"costPerDay"`
}
