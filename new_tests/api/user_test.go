package api

import (
	"testing"

	"wash-bonus/internal/api/restapi/client"
	user "wash-bonus/internal/api/restapi/client/user"

	cl "github.com/go-openapi/runtime/client"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)

func TestGetUser(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := user.NewGetUserParams()
	params.Body.ID = testUserID1
	mockApp.EXPECT().GetUser(gomock.Any(), gomock.Any()).Return(appUser(testUser1, true), nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		geted, err := c.User.GetUser(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
		t.DeepEqual(geted.Payload, apiUser(appUser(testUser1, true)))
	})
}

func TestAddUser(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := user.NewAddUserParams()
	params.Body = testAddUser1
	mockApp.EXPECT().AddUser(gomock.Any(), gomock.Any()).Return(appUser(testUser1, true), nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		geted, err := c.User.AddUser(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
		t.DeepEqual(geted.Payload, apiUser(appUser(testUser1, true)))
	})
}

func TestEditUser(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := user.NewEditUserParams()
	params.Body.Data = testAddUser1
	params.Body.ID = testUser1.ID
	mockApp.EXPECT().EditUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		_, err := c.User.EditUser(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
	})
}

func TestDeleteUser(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := user.NewDeleteUserParams()
	params.Body.ID = testUser1.ID
	mockApp.EXPECT().DeleteUser(gomock.Any(), gomock.Any()).Return(nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		_, err := c.User.DeleteUser(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
	})
}

func TestListUser(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := user.NewListUserParams()
	params.Body = testList
	mockApp.EXPECT().ListUser(gomock.Any(), gomock.Any()).Return(appUsers(testUsers, true), []string{}, nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		list, err := c.User.ListUser(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
		t.DeepEqual(list.Payload.Items, apiUsers(appUsers(testUsers, true)))
	})
}
