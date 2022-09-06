package api

import (
	"testing"

	"wash-bonus/internal/api/restapi/client"
	permission "wash-bonus/internal/api/restapi/client/permission"
	
	cl "github.com/go-openapi/runtime/client"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)


func TestGetPermission(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := permission.NewGetPermissionParams()
			params.Body.ID = testPermissionID1
		mockApp.EXPECT().GetPermission(gomock.Any(), gomock.Any()).Return(appPermission(testPermission1), nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			geted, err := c.Permission.GetPermission(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(geted.Payload, apiPermission(appPermission(testPermission1)))
		})
}

func TestAddPermission(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := permission.NewAddPermissionParams()	
		params.Body = testAddPermission1
		mockApp.EXPECT().AddPermission(gomock.Any(), gomock.Any()).Return(appPermission(testPermission1), nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			geted, err := c.Permission.AddPermission(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(geted.Payload, apiPermission(appPermission(testPermission1)))
		})
}

func TestEditPermission(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := permission.NewEditPermissionParams()
			params.Body.Data = testAddPermission1
				params.Body.ID = testPermission1.ID
		mockApp.EXPECT().EditPermission(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			_, err := c.Permission.EditPermission(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
		})
}

func TestDeletePermission(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := permission.NewDeletePermissionParams()
			params.Body.ID = testPermission1.ID
		mockApp.EXPECT().DeletePermission(gomock.Any(), gomock.Any()).Return(nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			_, err := c.Permission.DeletePermission(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
		})
}

func TestListPermission(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := permission.NewListPermissionParams()
		params.Body = testList
		mockApp.EXPECT().ListPermission(gomock.Any(), gomock.Any()).Return(appPermissions(testPermissions), []string{}, nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			list, err := c.Permission.ListPermission(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(list.Payload.Items, apiPermissions(appPermissions(testPermissions)))
		})
}
