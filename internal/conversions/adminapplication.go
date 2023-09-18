package conversions

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
	"washbonus/openapi/admin/models"

	"github.com/go-openapi/strfmt"
	uuid "github.com/satori/go.uuid"
)

func AdminApplicationFromDb(d dbmodels.AdminApplication) entities.AdminApplication {
	return entities.AdminApplication{
		ID: d.ID,
		User: entities.FirebaseUser{
			ID:    d.AdminUserID,
			Name:  d.Name,
			Email: d.Email,
		},
		Status: ApplicationStatusToApp(string(d.Status)),
	}
}

func AdminApplicationsFromDb(d []dbmodels.AdminApplication) []entities.AdminApplication {
	apps := make([]entities.AdminApplication, len(d))

	for idx, app := range d {
		apps[idx] = AdminApplicationFromDb(app)
	}

	return apps
}

func AdminApplicationToRest(e entities.AdminApplication) models.AdminApplication {
	id := e.ID.String()

	return models.AdminApplication{
		ID:     (*strfmt.UUID)(&id),
		Status: (*models.ApplicationStatusEnum)(&e.Status),
		User: &models.FirebaseUser{
			ID:    e.User.ID,
			Name:  e.User.Name,
			Email: strfmt.Email(e.User.Email),
		},
	}
}

func AdminApplicationsToRest(apps []entities.AdminApplication) []*models.AdminApplication {
	models := make([]*models.AdminApplication, len(apps))

	for idx, app := range apps {
		rest := AdminApplicationToRest(app)
		models[idx] = &rest
	}

	return models
}

func AdminApplicationCreationToDB(app entities.AdminApplicationCreation) dbmodels.AdminApplicationCreation {
	return dbmodels.AdminApplicationCreation{
		AdminUserID: app.User.ID,
		Name:        app.User.Name,
		Email:       app.User.Email,
	}
}

func AdminApplicationCreationFromRest(m models.AdminApplicationCreation) entities.AdminApplicationCreation {
	return entities.AdminApplicationCreation{
		User: entities.FirebaseUser{
			ID:    m.User.ID,
			Name:  m.User.Name,
			Email: m.User.Email.String(),
		},
	}
}

func AdminApplicationReviewToDB(app entities.AdminApplicationReview) dbmodels.AdminApplicationReview {
	return dbmodels.AdminApplicationReview{
		Status:         ApplicationStatusToDB(app.Status),
		OrganizationID: app.OrganizationID,
	}
}

func AdminApplicationReviewFromRest(m models.AdminApplicationReview) entities.AdminApplicationReview {
	var status entities.ApplicationStatus
	switch m.Status {
	case "accept":
		status = entities.Accepted
	case "reject":
		status = entities.Rejected
	}

	var orgID *uuid.UUID
	if m.OrganizationID != nil {
		parsedOrgID := uuid.FromStringOrNil(m.OrganizationID.String())
		orgID = &parsedOrgID
	}

	return entities.AdminApplicationReview{
		Status:         status,
		OrganizationID: orgID,
	}
}

func AdminApplicationFilterToDB(filter entities.AdminApplicationFilter) dbmodels.AdminApplicationFilter {
	mod := dbmodels.AdminApplicationFilter{
		Pagination: PaginationToDB(filter.Pagination),
	}

	if filter.Status != nil {
		stat := ApplicationStatusToDB(*filter.Status)
		mod.Status = &stat
	}

	return mod
}

func AdminApplicationFilterFromRest(pagination entities.Pagination, status *string) entities.AdminApplicationFilter {
	mod := entities.AdminApplicationFilter{
		Pagination: pagination,
	}

	if status != nil {
		stat := ApplicationStatusToApp(*status)
		mod.Status = &stat
	}

	return mod
}

func ApplicationStatusToApp(status string) entities.ApplicationStatus {
	switch status {
	case string(dbmodels.Accepted):
		return entities.Accepted
	case string(dbmodels.Rejected):
		return entities.Rejected
	case string(dbmodels.Pending):
		return entities.Pending
	default:
		panic("Unknown db application status, dbStatus - " + status)
	}
}

func ApplicationStatusToDB(status entities.ApplicationStatus) dbmodels.ApplicationStatus {
	switch status {
	case entities.Accepted:
		return dbmodels.Accepted
	case entities.Rejected:
		return dbmodels.Rejected
	case entities.Pending:
		return dbmodels.Pending
	default:
		panic("Unknown app application status, appStatus - " + status)
	}
}
