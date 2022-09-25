package entity

type IdentityProfile struct {
	UID    string
	Email  string
	Claims map[string]bool
}
