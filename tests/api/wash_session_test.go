package api

import (
	"testing"

	"wash-bonus/internal/api/restapi/client"
	washSession "wash-bonus/internal/api/restapi/client/wash_session"

	cl "github.com/go-openapi/runtime/client"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)

func TestGetWashSession(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := washSession.NewGetWashSessionParams()
	params.Body.ID = testWashSessionID1
	mockApp.EXPECT().GetWashSession(gomock.Any(), gomock.Any()).Return(appWashSession(testWashSession1, true), nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		geted, err := c.WashSession.GetWashSession(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
		t.DeepEqual(geted.Payload, apiWashSession(appWashSession(testWashSession1, true)))
	})
}

func TestAddWashSession(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := washSession.NewAddWashSessionParams()
	params.Body = testAddWashSession1
	mockApp.EXPECT().AddWashSession(gomock.Any(), gomock.Any()).Return(appWashSession(testWashSession1, true), nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		geted, err := c.WashSession.AddWashSession(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
		t.DeepEqual(geted.Payload, apiWashSession(appWashSession(testWashSession1, true)))
	})
}

func TestEditWashSession(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := washSession.NewEditWashSessionParams()
	params.Body.Data = testAddWashSession1
	params.Body.ID = testWashSession1.ID
	mockApp.EXPECT().EditWashSession(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		_, err := c.WashSession.EditWashSession(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
	})
}

func TestDeleteWashSession(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := washSession.NewDeleteWashSessionParams()
	params.Body.ID = testWashSession1.ID
	mockApp.EXPECT().DeleteWashSession(gomock.Any(), gomock.Any()).Return(nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		_, err := c.WashSession.DeleteWashSession(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
	})
}

func TestListWashSession(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := washSession.NewListWashSessionParams()
	params.Body = testList
	mockApp.EXPECT().ListWashSession(gomock.Any(), gomock.Any()).Return(appWashSessions(testWashSessions, true), []string{}, nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		list, err := c.WashSession.ListWashSession(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
		t.DeepEqual(list.Payload.Items, apiWashSessions(appWashSessions(testWashSessions, true)))
	})
}
