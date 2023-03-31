package vo

type RoutingKey string

const (
	WashAdminRoutingKey             RoutingKey = "wash_admin"
	WashAdminServesEventsRoutingKey            = "wash_admin_servers"
	WashBonusRoutingKey                        = "wash_bonus"
)
