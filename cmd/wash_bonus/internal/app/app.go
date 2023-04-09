package app

type Auth struct {
	UID          string
	Disabled     bool
	UserMetadata *AuthUserMeta
}

type AuthUserMeta struct {
	CreationTimestamp  int64
	LastLogInTimestamp int64

	LastRefreshTimestamp int64
}
