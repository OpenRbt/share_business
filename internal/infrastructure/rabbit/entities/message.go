package rabbitEntities

type MessageType string

const (
	SessionRequestMessageType      MessageType = "bonus_service/session/request"
	SessionCreatedMessageType      MessageType = "bonus_service/session/created"
	SessionStartMessageType        MessageType = "bonus_service/session/start"
	SessionFinishMessageType       MessageType = "bonus_service/session/finish"
	SessionStateMessageType        MessageType = "bonus_service/session/state"
	SessionUserMessageType         MessageType = "bonus_service/session/user"
	SessionEventMessageType        MessageType = "bonus_service/session/event"
	SessionBonusChargeMessageType  MessageType = "bonus_service/session/bonus/charge"
	SessionBonusConfirmMessageType MessageType = "bonus_service/session/bonus/confirm"
	SessionBonusDiscardMessageType MessageType = "bonus_service/session/bonus/discard"
	SessionBonusRewardMessageType  MessageType = "bonus_service/session/bonus/reward"
	SessionMoneyReportMessageType  MessageType = "bonus_service/session/money-report"
	WashServerDeletionMessageType  MessageType = "bonus_service/wash_server/delete"

	AdminUserMessageType        MessageType = "admin_service/admin_user"
	OrganizationMessageType     MessageType = "admin_service/organization"
	ServerGroupMessageType      MessageType = "admin_service/server_group"
	RequestAdminDataMessageType MessageType = "admin_service/data"
	CreateUserType              MessageType = "admin_service/create_user"
)
