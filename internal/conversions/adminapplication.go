package conversions

import (
	"fmt"
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

func adminRoleFromRest(role models.AdminUserRole) (*entities.Role, error) {
	if role == models.AdminUserRoleNoAccess {
		return nil, entities.ErrInvalidRole
	}
	r, err := RoleSelectionRest(role)
	return &r, err
}

func adminStatusFromRest(status models.ApplicationStatusEnum, role *models.AdminUserRole, orgID *strfmt.UUID) (entities.ApplicationStatus, error) {
	if status != models.ApplicationStatusEnumAccepted && status != models.ApplicationStatusEnumRejected {
		return "", fmt.Errorf("Unknown status - %s: %w", status, entities.ErrBadRequest)
	}

	if status == models.ApplicationStatusEnumRejected {
		return entities.Rejected, nil
	}

	if orgID == nil && (role == nil || *role != models.AdminUserRoleSystemManager) {
		return "", entities.ErrOrganizationIDRequired
	}

	return entities.Accepted, nil
}

func AdminApplicationReviewToDB(app entities.AdminApplicationReview) dbmodels.AdminApplicationReview {
	if app.Role == nil {
		return dbmodels.AdminApplicationReview{
			Status:         ApplicationStatusToDB(app.Status),
			OrganizationID: app.OrganizationID,
		}
	}

	r := RoleSelectionDB(*app.Role)
	return dbmodels.AdminApplicationReview{
		Status:         ApplicationStatusToDB(app.Status),
		OrganizationID: app.OrganizationID,
		Role:           &r,
	}
}

func AdminApplicationReviewFromRest(m models.AdminApplicationReview) (entities.AdminApplicationReview, error) {
	status, err := adminStatusFromRest(m.Status, m.Role, m.OrganizationID)
	if err != nil {
		return entities.AdminApplicationReview{}, err
	}

	var orgID *uuid.UUID
	if m.OrganizationID != nil {
		parsedOrgID, err := uuid.FromString(m.OrganizationID.String())
		if err != nil {
			return entities.AdminApplicationReview{}, fmt.Errorf("Wrong Organization ID: %w", entities.ErrBadRequest)
		}

		orgID = &parsedOrgID
	}

	var role *entities.Role
	if m.Role != nil {
		role, err = adminRoleFromRest(*m.Role)
		if err != nil {
			return entities.AdminApplicationReview{}, err
		}
	}

	return entities.AdminApplicationReview{
		Status:         status,
		OrganizationID: orgID,
		Role:           role,
	}, nil
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
