package api

import (
	"testing"

	"wash-bonus/internal/api/restapi/client"
	role "wash-bonus/internal/api/restapi/client/role"

	cl "github.com/go-openapi/runtime/client"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)

func TestGetRole(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := role.NewGetRoleParams()
	params.Body.ID = testRoleID1
	mockApp.EXPECT().GetRole(gomock.Any(), gomock.Any()).Return(appRole(testRole1, true), nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		geted, err := c.Role.GetRole(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
		t.DeepEqual(geted.Payload, apiRole(appRole(testRole1, true)))
	})
}

func TestAddRole(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := role.NewAddRoleParams()
	params.Body = testAddRole1
	mockApp.EXPECT().AddRole(gomock.Any(), gomock.Any()).Return(appRole(testRole1, true), nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		geted, err := c.Role.AddRole(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
		t.DeepEqual(geted.Payload, apiRole(appRole(testRole1, true)))
	})
}

func TestEditRole(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := role.NewEditRoleParams()
	params.Body.Data = testAddRole1
	params.Body.ID = testRole1.ID
	mockApp.EXPECT().EditRole(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		_, err := c.Role.EditRole(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
	})
}

func TestDeleteRole(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := role.NewDeleteRoleParams()
	params.Body.ID = testRole1.ID
	mockApp.EXPECT().DeleteRole(gomock.Any(), gomock.Any()).Return(nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		_, err := c.Role.DeleteRole(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
	})
}

func TestListRole(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := role.NewListRoleParams()
	params.Body = testList
	mockApp.EXPECT().ListRole(gomock.Any(), gomock.Any()).Return(appRoles(testRoles, true), []string{}, nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		list, err := c.Role.ListRole(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
		t.DeepEqual(list.Payload.Items, apiRoles(appRoles(testRoles, true)))
	})
}
