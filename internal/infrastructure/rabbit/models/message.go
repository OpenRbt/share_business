package models

type SessionCreation struct {
	NewSessions []string `json:"new_sessions,omitempty"`
	PostID      int64    `json:"post_id,omitempty"`
}

type SessionState int

const (
	SessionStateStart SessionState = iota
	SessionStateFinish
)
