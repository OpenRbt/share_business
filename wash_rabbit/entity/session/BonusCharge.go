package session

type BonusCharge struct {
	SessionID string `json:"session_id,omitempty"`
	Amount    int64  `json:"amount,omitempty"`
}
