package session

type UserAssign struct {
	SessionID string `json:"session_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
}
