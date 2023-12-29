package rabbitEntities

type ServerGroup struct {
	ID             string `json:"id"`
	OrganizationID string `json:"organizationId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	UTCOffset      int32  `json:"utcOffset"`
	IsDefault      bool   `json:"isDefault"`
	Deleted        bool   `json:"deleted"`
	Version        int    `json:"version"`
}
