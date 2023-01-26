package models

import "time"

type (
	SessionCreation struct {
		NewSessions []string `json:"new_sessions,omitempty"`
		PostID      int64    `json:"post_id,omitempty"`
	}

	SessionRequest struct {
		NewSessionsAmount int64 `json:"new_sessions_amount,omitempty"`
		PostID            int64 `json:"post_id,omitempty"`
	}

	SessionStateChange struct {
		SessionID      string                 `json:"session_id,omitempty"`
		State          int                    `json:"state,omitempty"`
		AdditionalData map[string]interface{} `json:"additional_data,omitempty"`
	}

	SessionUserAssign struct {
		SessionID string `json:"session_id,omitempty"`
		UserID    string `json:"user_id,omitempty"`
	}

	SessionBonusCharge struct {
		SessionID string `json:"session_id,omitempty"`
		Amount    int64  `json:"amount,omitempty"`
	}

	SessionEventLog struct {
		SessionID string                 `json:"session_id,omitempty"`
		Date      time.Time              `json:"date"`
		EventType int64                  `json:"event_type,omitempty"`
		Payload   map[string]interface{} `json:"payload,omitempty"`
	}

	ServerRegistered struct {
		ID          string `json:"id,omitempty"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	}

	ServerUpdate struct {
		ID          string  `json:"id,omitempty"`
		Title       *string `json:"title,omitempty"`
		Description *string `json:"description,omitempty"`
		Deleted     *bool   `json:"deleted,omitempty"`
	}
)
