package api

import (
	"testing"

	"wash-bonus/internal/api/restapi/client"
	washServer "wash-bonus/internal/api/restapi/client/wash_server"
	
	cl "github.com/go-openapi/runtime/client"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)


func TestGetWashServer(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := washServer.NewGetWashServerParams()
			params.Body.ID = testWashServerID1
		mockApp.EXPECT().GetWashServer(gomock.Any(), gomock.Any()).Return(appWashServer(testWashServer1), nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			geted, err := c.WashServer.GetWashServer(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(geted.Payload, apiWashServer(appWashServer(testWashServer1)))
		})
}

func TestAddWashServer(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := washServer.NewAddWashServerParams()	
		params.Body = testAddWashServer1
		mockApp.EXPECT().AddWashServer(gomock.Any(), gomock.Any()).Return(appWashServer(testWashServer1), nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			geted, err := c.WashServer.AddWashServer(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(geted.Payload, apiWashServer(appWashServer(testWashServer1)))
		})
}

func TestEditWashServer(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := washServer.NewEditWashServerParams()
			params.Body.Data = testAddWashServer1
				params.Body.ID = testWashServer1.ID
		mockApp.EXPECT().EditWashServer(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			_, err := c.WashServer.EditWashServer(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
		})
}

func TestDeleteWashServer(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := washServer.NewDeleteWashServerParams()
			params.Body.ID = testWashServer1.ID
		mockApp.EXPECT().DeleteWashServer(gomock.Any(), gomock.Any()).Return(nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			_, err := c.WashServer.DeleteWashServer(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
		})
}

func TestListWashServer(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
		mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := washServer.NewListWashServerParams()
		params.Body = testList
		mockApp.EXPECT().ListWashServer(gomock.Any(), gomock.Any()).Return(appWashServers(testWashServers), []string{}, nil)

		t.Run("", func(tt *testing.T) {
			t := check.T(tt)
			list, err := c.WashServer.ListWashServer(params, cl.APIKeyAuth("Authorization", "header", sess))
			t.Nil(err)
			t.DeepEqual(list.Payload.Items, apiWashServers(appWashServers(testWashServers)))
		})
}
