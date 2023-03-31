package vo

type MessageType string

const (
	SessionRequestMessageType      MessageType = "bonus_service/session/request"
	SessionCreatedMessageType                  = "bonus_service/session/created"
	SessionStartMessageType                    = "bonus_service/session/start"
	SessionFinishMessageType                   = "bonus_service/session/finish"
	SessionStateMessageType                    = "bonus_service/session/state"
	SessionUserMessageType                     = "bonus_service/session/user"
	SessionEventMessageType                    = "bonus_service/session/event"
	SessionBonusChargeMessageType              = "bonus_service/session/bonus/charge"
	SessionBonusConfirmMessageType             = "bonus_service/session/bonus/confirm"
	SessionBonusDiscardMessageType             = "bonus_service/session/bonus/discard"

	AdminServerRegisteredMessageType MessageType = "admin_service/server/registered"
	AdminServerUpdatedMessageType                = "admin_service/server/updated"
)
