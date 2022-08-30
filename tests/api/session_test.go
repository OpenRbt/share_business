package api

import (
	"testing"

	"wash-bonus/internal/api/restapi/client"
	session "wash-bonus/internal/api/restapi/client/session"

	cl "github.com/go-openapi/runtime/client"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)

func TestGetSession(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := session.NewGetSessionParams()
	params.Body.ID = testSessionID1
	mockApp.EXPECT().GetSession(gomock.Any(), gomock.Any()).Return(appSession(testSession1, true), nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		geted, err := c.Session.GetSession(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
		t.DeepEqual(geted.Payload, apiSession(appSession(testSession1, true)))
	})
}

func TestAddSession(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := session.NewAddSessionParams()
	params.Body = testAddSession1
	mockApp.EXPECT().AddSession(gomock.Any(), gomock.Any()).Return(appSession(testSession1, true), nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		geted, err := c.Session.AddSession(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
		t.DeepEqual(geted.Payload, apiSession(appSession(testSession1, true)))
	})
}

func TestEditSession(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := session.NewEditSessionParams()
	params.Body.Data = testAddSession1
	params.Body.ID = testSession1.ID
	mockApp.EXPECT().EditSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		_, err := c.Session.EditSession(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
	})
}

func TestDeleteSession(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := session.NewDeleteSessionParams()
	params.Body.ID = testSession1.ID
	mockApp.EXPECT().DeleteSession(gomock.Any(), gomock.Any()).Return(nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		_, err := c.Session.DeleteSession(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
	})
}

func TestListSession(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := session.NewListSessionParams()
	params.Body = testList
	mockApp.EXPECT().ListSession(gomock.Any(), gomock.Any()).Return(appSessions(testSessions, true), []string{}, nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		list, err := c.Session.ListSession(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
		t.DeepEqual(list.Payload.Items, apiSessions(appSessions(testSessions, true)))
	})
}
