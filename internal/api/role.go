// Code generated by mtgroup-generator.
package api

import (
	"errors"
	"wash-bonus/internal/def"

	"github.com/go-openapi/swag"

	"wash-bonus/internal/api/restapi/models"
	role "wash-bonus/internal/api/restapi/restapi/operations/role"
	"wash-bonus/internal/app"

	extauthapi "wash-bonus/internal/authentication"

	"github.com/go-openapi/runtime/middleware"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!
func (svc *service) GetRole(params role.GetRoleParams, profile interface{}) middleware.Responder {
	prof := profile.(*extauthapi.Profile)
	c, err := svc.app.GetRole(toAppProfile(prof), params.Body.ID)
	switch {
	default:
		log.PrintErr("GetRole server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return role.NewGetRoleDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("GetRole client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return role.NewGetRoleDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("GetRole client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return role.NewGetRoleDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("GetRole ok", "id", params.Body.ID)
		return role.NewGetRoleOK().WithPayload(apiRole(c))
	}
}
func (svc *service) AddRole(params role.AddRoleParams, profile interface{}) middleware.Responder {
	prof := profile.(*extauthapi.Profile)
	c, err := svc.app.AddRole(toAppProfile(prof), appRoleAdd(params.Body))
	switch {
	default:
		log.PrintErr("AddRole server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return role.NewAddRoleDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("AddRole client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return role.NewAddRoleDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("AddRole ok")
		return role.NewAddRoleCreated().WithPayload(apiRole(c))
	}
}
func (svc *service) EditRole(params role.EditRoleParams, profile interface{}) middleware.Responder {
	prof := profile.(*extauthapi.Profile)
	err := svc.app.EditRole(toAppProfile(prof), params.Body.ID, appRoleAdd(params.Body.Data))
	switch {
	default:
		log.PrintErr("EditRole server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return role.NewEditRoleDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("EditRole client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return role.NewEditRoleDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("EditRole client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return role.NewEditRoleDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("EditRole ok")
		return role.NewEditRoleOK()
	}
}
func (svc *service) DeleteRole(params role.DeleteRoleParams, profile interface{}) middleware.Responder {
	prof := profile.(*extauthapi.Profile)
	err := svc.app.DeleteRole(toAppProfile(prof), params.Body.ID)
	switch {
	default:
		log.PrintErr("DeleteRole server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return role.NewDeleteRoleDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("DeleteRole client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return role.NewDeleteRoleDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("DeleteRole client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return role.NewDeleteRoleDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("DeleteRole ok", "id", params.Body.ID)
		return role.NewDeleteRoleNoContent()
	}
}
func (svc *service) ListRole(params role.ListRoleParams, profile interface{}) middleware.Responder {
	prof := profile.(*extauthapi.Profile)
	c, warnings, err := svc.app.ListRole(toAppProfile(prof), appListParams(params.Body))
	switch {
	default:
		log.PrintErr("ListRole server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return role.NewListRoleDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrAccessDenied):
		log.Info("ListRole client error", def.LogHTTPStatus, codeForbidden.status, "code", codeForbidden.extra, "err", err)
		return role.NewListRoleDefault(codeForbidden.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeForbidden.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("ListRole ok")
		return role.NewListRoleOK().WithPayload(&role.ListRoleOKBody{
			Items:    apiRoles(c),
			Warnings: warnings,
		})
	}
}

func (svc *service) AddPermissionsRole(params role.AddPermissionsRoleParams, profile interface{}) middleware.Responder {
	prof := profile.(*extauthapi.Profile)
	err := svc.app.AddPermissionsRole(params.Body.ID, prof.IsolatedEntityID.String(), params.Body.ItemsID, appPermissionsAdd(params.Body.Items))
	switch {
	default:
		log.PrintErr("AddPermissionsRole server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return role.NewAddPermissionsRoleDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("AddPermissionsRole client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return role.NewAddPermissionsRoleDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case errors.Is(err, app.ErrNotFoundArrayItem):
		log.Info("AddPermissionsRole client error", def.LogHTTPStatus, codeNotFoundArrayItem.status, "code", codeNotFoundArrayItem.extra, "err", err)
		return role.NewAddPermissionsRoleDefault(codeNotFoundArrayItem.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFoundArrayItem.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("AddPermissionsRole ok")
		return role.NewAddPermissionsRoleOK()
	}
}

func (svc *service) DeletePermissionsRole(params role.DeletePermissionsRoleParams, profile interface{}) middleware.Responder {
	prof := profile.(*extauthapi.Profile)
	err := svc.app.DeletePermissionsRole(params.Body.ID, prof.IsolatedEntityID.String(), params.Body.Items)
	switch {
	default:
		log.PrintErr("DeletePermissionsRole server error", def.LogHTTPStatus, codeInternal.status, "code", codeInternal.extra, "err", err)
		return role.NewDeletePermissionsRoleDefault(codeInternal.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeInternal.extra),
			Message: swag.String("internal error"),
		})
	case errors.Is(err, app.ErrNotFound):
		log.Info("DeletePermissionsRole client error", def.LogHTTPStatus, codeNotFound.status, "code", codeNotFound.extra, "err", err)
		return role.NewDeletePermissionsRoleDefault(codeNotFound.status).WithPayload(&models.Error{
			Code:    swag.Int32(codeNotFound.extra),
			Message: swag.String(err.Error()),
		})
	case err == nil:
		log.Info("DeletePermissionsRole ok")
		return role.NewDeletePermissionsRoleOK()
	}
}

func apiRole(a *app.Role) *models.Role {
	if a == nil {
		return nil
	}
	return &models.Role{
		ID:          a.ID,
		Active:      a.Active,
		Name:        a.Name,
		Permissions: apiPermissions(a.Permissions),
	}
}

func apiRoles(apps []*app.Role) []*models.Role {
	apis := []*models.Role{}
	for i := range apps {
		apis = append(apis, apiRole(apps[i]))
	}
	return apis
}

func appRole(a *models.Role, withStructs bool) *app.Role {
	if a == nil {
		return nil
	}
	role := &app.Role{}
	if withStructs {
		role.Permissions = appPermissions(a.Permissions)
	}
	role.ID = a.ID
	role.Active = a.Active
	role.Name = a.Name

	return role
}

func appRoles(apis []*models.Role, withStructs bool) []*app.Role {
	apps := []*app.Role{}
	for i := range apis {
		apps = append(apps, appRole(apis[i], withStructs))
	}
	return apps
}

func appRoleAdd(a *models.RoleAdd) *app.Role {
	if a == nil {
		return nil
	}
	role := &app.Role{}
	role.Active = a.Active
	role.Name = a.Name
	if len(a.Permissions) > 0 {
		for _, id := range a.Permissions {
			role.Permissions = append(role.Permissions, &app.Permission{
				ID: id,
			})
		}
	}

	return role
}

func appRolesAdd(apis []*models.RoleAdd) []*app.Role {
	apps := []*app.Role{}
	for i := range apis {
		apps = append(apps, appRoleAdd(apis[i]))
	}
	return apps
}
