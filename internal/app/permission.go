// Code generated by mtgroup-generator.
package app

import ()

// Make sure not to overwrite this file after you generated it because all your edits would be lost!
type Permission struct {
	ID   string
	Name string
}

func (a *app) GetPermission(prof Profile, id string) (*Permission, error) {
	if !a.rulesSet.GetPermissionAccessManager(prof) {
		return nil, ErrAccessDenied
	}
	return a.repo.GetPermission(id, prof.IsolatedEntityID)
}
func (a *app) AddPermission(prof Profile, m *Permission) (*Permission, error) {
	if !a.rulesSet.AddPermissionAccessManager(prof) {
		return nil, ErrAccessDenied
	}
	return a.repo.AddPermission(prof.ID, prof.IsolatedEntityID, m)
}
func (a *app) EditPermission(prof Profile, id string, m *Permission) error {
	if !a.rulesSet.EditPermissionAccessManager(prof) {
		return ErrAccessDenied
	}
	return a.repo.EditPermission(id, prof.IsolatedEntityID, m)
}
func (a *app) DeletePermission(prof Profile, id string) error {
	if !a.rulesSet.DeletePermissionAccessManager(prof) {
		return ErrAccessDenied
	}
	return a.repo.DeletePermission(id, prof.ID, prof.IsolatedEntityID)
}
func (a *app) ListPermission(prof Profile, params *ListParams) ([]*Permission, []string, error) {
	if !a.rulesSet.ListPermissionAccessManager(prof) {
		return nil, nil, ErrAccessDenied
	}
	return a.repo.ListPermission(prof.IsolatedEntityID, params)
}
