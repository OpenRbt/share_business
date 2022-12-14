// Code generated by mtgroup-generator.
package api

import (
	"time"

	"wash-bonus/internal/api/restapi/models"
	"wash-bonus/internal/app"

	"github.com/go-openapi/strfmt"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func apiSession(a *app.Session) *models.Session {
	if a == nil {
		return nil
	}
	return &models.Session{
		ID:           a.ID,
		Active:       a.Active,
		ClosingAt:    (*strfmt.DateTime)(a.ClosingAt),
		CreatedAt:    (*strfmt.DateTime)(a.CreatedAt),
		ExpirationAt: (*strfmt.DateTime)(a.ExpirationAt),
		UpdateAt:     (*strfmt.DateTime)(a.UpdateAt),
		User:         apiToken(a.User),
	}
}

func apiSessions(apps []*app.Session) []*models.Session {
	apis := []*models.Session{}
	for i := range apps {
		apis = append(apis, apiSession(apps[i]))
	}
	return apis
}

func appSession(a *models.Session, withStructs bool) *app.Session {
	if a == nil {
		return nil
	}
	session := &app.Session{}
	if withStructs {
		session.User = appToken(a.User)
	}
	session.ID = a.ID
	session.Active = a.Active
	session.ClosingAt = (*time.Time)(a.ClosingAt)
	session.CreatedAt = (*time.Time)(a.CreatedAt)
	session.ExpirationAt = (*time.Time)(a.ExpirationAt)
	session.UpdateAt = (*time.Time)(a.UpdateAt)

	return session
}

func appSessions(apis []*models.Session, withStructs bool) []*app.Session {
	apps := []*app.Session{}
	for i := range apis {
		apps = append(apps, appSession(apis[i], withStructs))
	}
	return apps
}

func appSessionAdd(a *models.SessionAdd) *app.Session {
	if a == nil {
		return nil
	}
	session := &app.Session{}
	session.Active = a.Active
	session.ClosingAt = (*time.Time)(a.ClosingAt)
	session.ExpirationAt = (*time.Time)(a.ExpirationAt)
	session.UpdateAt = (*time.Time)(a.UpdateAt)
	if a.User != "" {
		session.User = &app.Token{ID: a.User}
	}

	return session
}

func appSessionsAdd(apis []*models.SessionAdd) []*app.Session {
	apps := []*app.Session{}
	for i := range apis {
		apps = append(apps, appSessionAdd(apis[i]))
	}
	return apps
}
