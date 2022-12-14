// Code generated by mtgroup-generator.
package app

import (
	"time"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!
type User struct {
	Active     bool
	CreatedAt  *time.Time
	ID         string
	ModifiedAt *time.Time
	Role       *Role
}

func (a *app) GetUser(prof Profile, id string) (*User, error) {
	if !a.rulesSet.GetUserAccessManager(prof) {
		return nil, ErrAccessDenied
	}
	return a.repo.GetUser(id, prof.IsolatedEntityID)
}
func (a *app) AddUser(prof Profile, m *User) (*User, error) {
	if !a.rulesSet.AddUserAccessManager(prof) {
		return nil, ErrAccessDenied
	}
	return a.repo.AddUser(prof.ID, prof.IsolatedEntityID, m)
}
func (a *app) EditUser(prof Profile, id string, m *User) error {
	if !a.rulesSet.EditUserAccessManager(prof) {
		return ErrAccessDenied
	}
	return a.repo.EditUser(id, prof.IsolatedEntityID, m)
}
func (a *app) DeleteUser(prof Profile, id string) error {
	if !a.rulesSet.DeleteUserAccessManager(prof) {
		return ErrAccessDenied
	}
	return a.repo.DeleteUser(id, prof.ID, prof.IsolatedEntityID)
}
func (a *app) ListUser(prof Profile, params *ListParams) ([]*User, []string, error) {
	if !a.rulesSet.ListUserAccessManager(prof) {
		return nil, nil, ErrAccessDenied
	}
	return a.repo.ListUser(prof.IsolatedEntityID, params)
}
