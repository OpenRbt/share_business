package api

import (
	"fmt"
	"net/http"
	"testing"
		"time"

	"wash-bonus/internal/api/restapi/models"
	"wash-bonus/internal/app"
	"wash-bonus/internal/def"
	extauthapi "github.com/mtgroupit/mt-mock-extauthapi"
		"github.com/go-openapi/strfmt"
	"github.com/golang/mock/gomock"
	"github.com/powerman/check"
	"github.com/powerman/gotest/testinit"
	"github.com/google/uuid"
	"github.com/phayes/freeport"
)

var (
	isolatedEntityID = uuid.New().String()
	profileID           = uuid.New().String()
	sess             = extauthapi.SessionCookieName + "=sess"
	profile          = &extauthapi.Profile{
		ID:    extauthapi.MustParseID(profileID),
		Authn: true,
		IsolatedEntityID: extauthapi.MustParseID(isolatedEntityID),
	}
	testPermissionID1  = uuid.New().String()
	testPermission1 = &models.Permission{
		ID: testPermissionID1,
		Name: "ullam",
	}
		testPermissionID2  = uuid.New().String()
		testPermission2 = &models.Permission{
			ID: testPermissionID1,
			Name: "ea",
		}
		testPermissions = []*models.Permission{testPermission1, testPermission2}
	testAddPermission1 = &models.PermissionAdd{
		Name: "eos",
	}
	testAddPermission2 = &models.PermissionAdd{
		Name: "suscipit",
	}
	testRoleID1  = uuid.New().String()
	testRole1 = &models.Role{
		Active: false,
		ID: testRoleID1,
		Name: "illo",
		Permissions: testPermissions,
	}
		testRoleID2  = uuid.New().String()
		testRole2 = &models.Role{
			Active: true,
			ID: testRoleID1,
			Name: "aut",
			Permissions: testPermissions,
		}
		testRoles = []*models.Role{testRole1, testRole2}
	testAddRole1 = &models.RoleAdd{
		Active: false,
		Name: "tempora",
		Permissions: []string{testPermissionID1, testPermissionID2},
	}
	testAddRole2 = &models.RoleAdd{
		Active: true,
		Name: "unde",
		Permissions: []string{testPermissionID1, testPermissionID2},
	}
	testUserID1  = uuid.New().String()
	testUser1 = &models.User{
		Active: true,
		CreatedAt: toDateTime("1936-08-27T03:11:30.941Z"),
		FirebaseId: "amet",
		ID: testUserID1,
		ModifiedAt: toDateTime("1905-08-20T07:59:17.606Z"),
		Role: testRole1,
	}
		testUserID2  = uuid.New().String()
		testUser2 = &models.User{
			Active: true,
			CreatedAt: toDateTime("1924-05-22T12:55:15.083Z"),
			FirebaseId: "nobis",
			ID: testUserID1,
			ModifiedAt: toDateTime("1984-04-11T07:36:18.998Z"),
			Role: testRole2,
		}
		testUsers = []*models.User{testUser1, testUser2}
	testAddUser1 = &models.UserAdd{
		Active: false,
		FirebaseId: "et",
		Role: testRole1.ID,
	}
	testAddUser2 = &models.UserAdd{
		Active: true,
		FirebaseId: "molestiae",
		Role: testRole2.ID,
	}
	testWashServerID1  = uuid.New().String()
	testWashServer1 = &models.WashServer{
		CreatedAt: toDateTime("1913-09-07T09:28:44.012Z"),
		ID: testWashServerID1,
		Key: "maiores",
		LastUpdateAt: toDateTime("1942-07-24T19:44:43.383Z"),
		ModifiedAt: toDateTime("1926-02-21T19:19:31.361Z"),
		Name: "soluta",
	}
		testWashServerID2  = uuid.New().String()
		testWashServer2 = &models.WashServer{
			CreatedAt: toDateTime("1953-01-22T15:18:49.602Z"),
			ID: testWashServerID1,
			Key: "ut",
			LastUpdateAt: toDateTime("1947-07-01T00:59:35.103Z"),
			ModifiedAt: toDateTime("1933-10-04T12:19:04.065Z"),
			Name: "quae",
		}
		testWashServers = []*models.WashServer{testWashServer1, testWashServer2}
	testAddWashServer1 = &models.WashServerAdd{
		Key: "quia",
		LastUpdateAt: toDateTime("2008-07-07T21:56:05.503Z"),
		Name: "dolores",
	}
	testAddWashServer2 = &models.WashServerAdd{
		Key: "ratione",
		LastUpdateAt: toDateTime("1949-12-03T00:54:30.526Z"),
		Name: "quidem",
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

	port , err := freeport.GetFreePort()
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

	func toDateTime(date interface{}) (*strfmt.DateTime) {
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

	func toDate(date interface{}) (*strfmt.Date) {
		if date == nil {
			return nil
		}
		var dt strfmt.Date
		dt.Scan(date)
		return &dt
	}

