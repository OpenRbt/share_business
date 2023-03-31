package session

import (
	"github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
)

type StateChange struct {
	SessionID      string                 `json:"session_id,omitempty"`
	State          vo.SessionState        `json:"state,omitempty"`
	AdditionalData map[string]interface{} `json:"additional_data,omitempty"`
}
