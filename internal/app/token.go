// Code generated by mtgroup-generator.
package app

import (
	"time"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!
type Token struct {
	ExpirationAt *time.Time
	ID           string
	Token        string
	Type         string
}

func (a *app) GetToken(prof Profile, id string) (*Token, error) {
	if !a.rulesSet.GetTokenAccessManager(prof) {
		return nil, ErrAccessDenied
	}
	return a.repo.GetToken(id, prof.IsolatedEntityID)
}
func (a *app) AddToken(prof Profile, m *Token) (*Token, error) {
	if !a.rulesSet.AddTokenAccessManager(prof) {
		return nil, ErrAccessDenied
	}
	return a.repo.AddToken(prof.ID, prof.IsolatedEntityID, m)
}
func (a *app) DeleteToken(prof Profile, id string) error {
	if !a.rulesSet.DeleteTokenAccessManager(prof) {
		return ErrAccessDenied
	}
	return a.repo.DeleteToken(id, prof.ID, prof.IsolatedEntityID)
}
