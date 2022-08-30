package api

import (
	"testing"

	"wash-bonus/internal/api/restapi/client"
	token "wash-bonus/internal/api/restapi/client/token"

	cl "github.com/go-openapi/runtime/client"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
)

func TestGetToken(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := token.NewGetTokenParams()
	params.Body.ID = testTokenID1
	mockApp.EXPECT().GetToken(gomock.Any(), gomock.Any()).Return(appToken(testToken1), nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		geted, err := c.Token.GetToken(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
		t.DeepEqual(geted.Payload, apiToken(appToken(testToken1)))
	})
}

func TestAddToken(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := token.NewAddTokenParams()
	params.Body = testAddToken1
	mockApp.EXPECT().AddToken(gomock.Any(), gomock.Any()).Return(appToken(testToken1), nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		geted, err := c.Token.AddToken(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
		t.DeepEqual(geted.Payload, apiToken(appToken(testToken1)))
	})
}

func TestDeleteToken(tt *testing.T) {
	t := check.T(tt)
	// t.Parallel()
	tsURL, shutdown, mockApp, mockExtAuthSvc := testNewServer(t)
	defer shutdown()
	c := client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host: tsURL,
	})
	mockExtAuthSvc.EXPECT().GetUserProfile(gomock.Any(), sess).Return(profile, nil)

	params := token.NewDeleteTokenParams()
	params.Body.ID = testToken1.ID
	mockApp.EXPECT().DeleteToken(gomock.Any(), gomock.Any()).Return(nil)

	t.Run("", func(tt *testing.T) {
		t := check.T(tt)
		_, err := c.Token.DeleteToken(params, cl.APIKeyAuth("Authorization", "header", sess))
		t.Nil(err)
	})
}
