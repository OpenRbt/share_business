package rabbitEntities

type Organization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	IsDefault   bool   `json:"isDefault"`
	Deleted     bool   `json:"deleted"`
	Version     int    `json:"version"`
}
