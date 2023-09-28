package session

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
