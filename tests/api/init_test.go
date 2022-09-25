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

	extauthapi "github.com/mtgroupit/mt-mock-extauthapi"

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
	testUserID1 = uuid.New().String()
	testUser1   = &models.User{
		Active:     true,
		CreatedAt:  toDateTime("1999-02-01T19:26:16.387Z"),
		ID:         testUserID1,
		ModifiedAt: toDateTime("2017-08-18T00:05:29.812Z"),
	}
	testUserID2 = uuid.New().String()
	testUser2   = &models.User{
		Active:     true,
		CreatedAt:  toDateTime("1971-02-10T18:28:52.970Z"),
		ID:         testUserID1,
		ModifiedAt: toDateTime("1987-01-27T01:53:29.145Z"),
	}
	testUsers    = []*models.User{testUser1, testUser2}
	testAddUser1 = &models.UserAdd{
		Active: false,
	}
	testAddUser2 = &models.UserAdd{
		Active: false,
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
