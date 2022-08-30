package api

import (
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"wash-bonus/internal/api/restapi/models"
	"wash-bonus/internal/app"
	"wash-bonus/internal/def"

	extauthapi "wash-bonus/internal/authentication"

	"github.com/go-openapi/strfmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/phayes/freeport"
	"github.com/powerman/check"
	"github.com/powerman/gotest/testinit"
)

var (
	isolatedEntityID = uuid.New().String()
	profileID        = uuid.New().String()
	sess             = extauthapi.SessionCookieName + "=sess"
	profile          = &extauthapi.Profile{
		ID:               extauthapi.MustParseID(profileID),
		Authn:            true,
		IsolatedEntityID: extauthapi.MustParseID(isolatedEntityID),
	}
	testPermissionID1 = uuid.New().String()
	testPermission1   = &models.Permission{
		ID:   testPermissionID1,
		Name: "aperiam",
	}
	testPermissionID2 = uuid.New().String()
	testPermission2   = &models.Permission{
		ID:   testPermissionID1,
		Name: "nemo",
	}
	testPermissions    = []*models.Permission{testPermission1, testPermission2}
	testAddPermission1 = &models.PermissionAdd{
		Name: "quos",
	}
	testAddPermission2 = &models.PermissionAdd{
		Name: "sapiente",
	}
	testRoleID1 = uuid.New().String()
	testRole1   = &models.Role{
		Active:      false,
		ID:          testRoleID1,
		Name:        "modi",
		Permissions: testPermissions,
	}
	testRoleID2 = uuid.New().String()
	testRole2   = &models.Role{
		Active:      true,
		ID:          testRoleID1,
		Name:        "alias",
		Permissions: testPermissions,
	}
	testRoles    = []*models.Role{testRole1, testRole2}
	testAddRole1 = &models.RoleAdd{
		Active:      false,
		Name:        "aliquid",
		Permissions: []string{testPermissionID1, testPermissionID2},
	}
	testAddRole2 = &models.RoleAdd{
		Active:      true,
		Name:        "ut",
		Permissions: []string{testPermissionID1, testPermissionID2},
	}
	testSessionID1 = uuid.New().String()
	testSession1   = &models.Session{
		Active:       false,
		ClosingAt:    toDateTime("1937-09-19T21:54:22.215Z"),
		CreatedAt:    toDateTime("1985-08-14T02:28:39.697Z"),
		ExpirationAt: toDateTime("2016-04-26T07:50:25.083Z"),
		ID:           testSessionID1,
		UpdateAt:     toDateTime("2003-11-22T13:13:44.768Z"),
		User:         testToken1,
	}
	testSessionID2 = uuid.New().String()
	testSession2   = &models.Session{
		Active:       true,
		ClosingAt:    toDateTime("1935-01-05T15:23:19.581Z"),
		CreatedAt:    toDateTime("1902-09-13T19:47:39.137Z"),
		ExpirationAt: toDateTime("1959-05-01T00:24:24.410Z"),
		ID:           testSessionID1,
		UpdateAt:     toDateTime("1951-05-07T22:54:53.427Z"),
		User:         testToken2,
	}
	testSessions    = []*models.Session{testSession1, testSession2}
	testAddSession1 = &models.SessionAdd{
		Active:       true,
		ClosingAt:    toDateTime("1993-12-02T14:57:09.635Z"),
		ExpirationAt: toDateTime("1941-11-21T14:33:23.468Z"),
		UpdateAt:     toDateTime("1938-01-24T19:49:58.795Z"),
		User:         testToken1.ID,
	}
	testAddSession2 = &models.SessionAdd{
		Active:       false,
		ClosingAt:    toDateTime("1982-01-22T19:04:20.712Z"),
		ExpirationAt: toDateTime("1923-08-21T03:46:47.351Z"),
		UpdateAt:     toDateTime("1990-05-22T23:45:54.490Z"),
		User:         testToken2.ID,
	}
	testTokenID1 = uuid.New().String()
	testToken1   = &models.Token{
		ExpirationAt: toDateTime("2007-07-15T20:00:01.050Z"),
		ID:           testTokenID1,
		Token:        "eveniet",
		Type:         "hic",
	}
	testTokenID2 = uuid.New().String()
	testToken2   = &models.Token{
		ExpirationAt: toDateTime("1946-09-27T09:11:22.090Z"),
		ID:           testTokenID1,
		Token:        "nobis",
		Type:         "ab",
	}
	testTokens    = []*models.Token{testToken1, testToken2}
	testAddToken1 = &models.TokenAdd{
		ExpirationAt: toDateTime("2010-07-01T17:20:49.458Z"),
		Token:        "optio",
		Type:         "eum",
	}
	testAddToken2 = &models.TokenAdd{
		ExpirationAt: toDateTime("1941-08-16T19:04:13.245Z"),
		Token:        "qui",
		Type:         "delectus",
	}
	testUserID1 = uuid.New().String()
	testUser1   = &models.User{
		Active:     true,
		CreatedAt:  toDateTime("1999-02-01T19:26:16.387Z"),
		ID:         testUserID1,
		ModifiedAt: toDateTime("2017-08-18T00:05:29.812Z"),
		Role:       testRole1,
	}
	testUserID2 = uuid.New().String()
	testUser2   = &models.User{
		Active:     true,
		CreatedAt:  toDateTime("1971-02-10T18:28:52.970Z"),
		ID:         testUserID1,
		ModifiedAt: toDateTime("1987-01-27T01:53:29.145Z"),
		Role:       testRole2,
	}
	testUsers    = []*models.User{testUser1, testUser2}
	testAddUser1 = &models.UserAdd{
		Active: false,
		Role:   testRole1.ID,
	}
	testAddUser2 = &models.UserAdd{
		Active: false,
		Role:   testRole2.ID,
	}
	testWashServerID1 = uuid.New().String()
	testWashServer1   = &models.WashServer{
		CreatedAt:    toDateTime("1928-01-22T09:08:46.330Z"),
		ID:           testWashServerID1,
		Key:          "perspiciatis",
		LastUpdateAt: toDateTime("1932-02-08T17:48:20.291Z"),
		ModifiedAt:   toDateTime("1938-07-18T11:08:54.836Z"),
		Name:         "architecto",
	}
	testWashServerID2 = uuid.New().String()
	testWashServer2   = &models.WashServer{
		CreatedAt:    toDateTime("1964-04-20T00:06:13.353Z"),
		ID:           testWashServerID1,
		Key:          "impedit",
		LastUpdateAt: toDateTime("1975-07-01T17:09:49.368Z"),
		ModifiedAt:   toDateTime("1995-04-01T15:38:12.897Z"),
		Name:         "cum",
	}
	testWashServers    = []*models.WashServer{testWashServer1, testWashServer2}
	testAddWashServer1 = &models.WashServerAdd{
		Key:          "possimus",
		LastUpdateAt: toDateTime("2003-05-26T00:16:31.728Z"),
		Name:         "iste",
	}
	testAddWashServer2 = &models.WashServerAdd{
		Key:          "impedit",
		LastUpdateAt: toDateTime("1927-07-27T09:54:35.103Z"),
		Name:         "nihil",
	}
	testWashSessionID1 = uuid.New().String()
	testWashSession1   = &models.WashSession{
		Active:       false,
		ClosingAt:    toDateTime("1927-03-25T13:08:06.200Z"),
		CreatedAt:    toDateTime("1963-10-26T08:40:08.768Z"),
		ExpirationAt: toDateTime("1984-05-12T19:31:47.067Z"),
		ID:           testWashSessionID1,
		UpdateAt:     toDateTime("1947-06-07T10:29:50.719Z"),
		User:         testToken1,
		WashServer:   testWashServer1,
	}
	testWashSessionID2 = uuid.New().String()
	testWashSession2   = &models.WashSession{
		Active:       true,
		ClosingAt:    toDateTime("1917-09-16T06:53:12.344Z"),
		CreatedAt:    toDateTime("1930-07-19T07:10:08.759Z"),
		ExpirationAt: toDateTime("1925-05-01T20:09:28.593Z"),
		ID:           testWashSessionID1,
		UpdateAt:     toDateTime("1908-05-24T07:21:41.633Z"),
		User:         testToken2,
		WashServer:   testWashServer2,
	}
	testWashSessions    = []*models.WashSession{testWashSession1, testWashSession2}
	testAddWashSession1 = &models.WashSessionAdd{
		Active:       true,
		ClosingAt:    toDateTime("1934-11-23T15:55:28.891Z"),
		ExpirationAt: toDateTime("1945-03-09T06:08:08.348Z"),
		UpdateAt:     toDateTime("1989-12-05T05:31:52.006Z"),
		User:         testToken1.ID,
		WashServer:   testWashServer1.ID,
	}
	testAddWashSession2 = &models.WashSessionAdd{
		Active:       true,
		ClosingAt:    toDateTime("2011-11-30T08:43:38.238Z"),
		ExpirationAt: toDateTime("1927-01-12T19:54:58.261Z"),
		UpdateAt:     toDateTime("1966-09-19T06:53:07.594Z"),
		User:         testToken2.ID,
		WashServer:   testWashServer2.ID,
	}

	offset int64 = 0
	limit  int64 = 5

	testList = &models.ListParams{
		Offset: &offset,
		Limit:  limit,
	}
)

func TestMain(m *testing.M) { testinit.Main(m) }

func testNewServer(t *check.C) (string, func(), *app.MockApp, *MockAuthSvc) {
	t.Helper()
	ctrl := gomock.NewController(t)

	mockApp := app.NewMockApp(ctrl)
	mockExtAuthSvc := NewMockAuthSvc(ctrl)

	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}

	server, err := NewServer(mockApp, mockExtAuthSvc, Config{
		Host:     "localhost",
		Port:     port,
		BasePath: def.APIBasePath,
	})
	t.Nil(err, "NewServer")
	t.Nil(server.Listen(), "server.Listen")
	errc := make(chan error, 1)
	go func() { errc <- server.Serve() }()

	shutdown := func() {
		t.Helper()
		t.Nil(server.Shutdown(), "server.Shutdown")
		t.Nil(<-errc, "server.Serve")
		ctrl.Finish()
	}

	url := fmt.Sprintf("localhost:%d", server.Port)

	return url, shutdown, mockApp, mockExtAuthSvc
}

type matchCookie string // Implements gomock.Matcher.

func (m matchCookie) String() string { return string(m) }
func (m matchCookie) Matches(x interface{}) bool {
	for _, c := range (&http.Request{Header: map[string][]string{"Cookie": {x.(string)}}}).Cookies() {
		if c.String() == string(m) {
			return true
		}
	}
	return false
}
func fromDateTime(dt strfmt.DateTime) time.Time {
	return time.Time(dt)
}

func toDateTime(date interface{}) *strfmt.DateTime {
	if date == nil {
		return nil
	}
	var dt strfmt.DateTime
	dt.Scan(date)
	return &dt
}

func fromDate(d strfmt.Date) time.Time {
	return time.Time(d)
}

func toDate(date interface{}) *strfmt.Date {
	if date == nil {
		return nil
	}
	var dt strfmt.Date
	dt.Scan(date)
	return &dt
}
