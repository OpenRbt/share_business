package rabbitEntities

type Organization struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	UTCOffset   int32  `json:"utcOffset"`
	IsDefault   bool   `json:"isDefault"`
	Deleted     bool   `json:"deleted"`
	Version     int    `json:"version"`
}
