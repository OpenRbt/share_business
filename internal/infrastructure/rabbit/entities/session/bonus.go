package session

type BonusCharge struct {
	SessionID string `json:"session_id,omitempty"`
	Amount    int64  `json:"amount,omitempty"`
	Post      int64  `json:"post,omitempty"`
}

type BonusChargeConfirm struct {
	SessionID string `json:"session_id,omitempty"`
	Amount    int64  `json:"amount,omitempty"`
}

type BonusChargeDiscard struct {
	SessionID string `json:"session_id,omitempty"`
	Amount    int64  `json:"amount,omitempty"`
}

type BonusReward struct {
	SessionID string `json:"session_id,omitempty"`
	Amount    int    `json:"amount,omitempty"`
	UUID      string `json:"uuid,omitempty"`
}
