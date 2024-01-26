package rabbitEntities

type SessionState int

const (
	SessionStateStart SessionState = iota
	SessionStateFinish
)

type NewSession struct {
	ID     string `json:"id,omitempty"`
	PostID int64  `json:"post_id,omitempty"`
}

type PostSessions struct {
	NewSessions []string `json:"new_sessions,omitempty"`
	PostID      int64    `json:"post_id,omitempty"`
}

type RequestSessions struct {
	NewSessionsAmount int64 `json:"new_sessions_amount,omitempty"`
	PostID            int64 `json:"post_id,omitempty"`
}

type ChangeSessionState struct {
	SessionID      string                 `json:"session_id,omitempty"`
	State          SessionState           `json:"state,omitempty"`
	AdditionalData map[string]interface{} `json:"additional_data,omitempty"`
}

type AssignUserToSession struct {
	SessionID string `json:"session_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	Post      int64  `json:"post,omitempty"`
}
