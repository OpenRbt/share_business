//go:build integration
// +build integration

package dal

import (
	"context"

	"github.com/google/uuid"
	"github.com/powerman/gotest/testinit"
	"github.com/powerman/pqx"
	"github.com/powerman/structlog"

	"wash-bonus/internal/app"
	"wash-bonus/internal/def"
)

var testRepo *Repo
var ctx = context.Background()

var (
	isolatedEntityID = uuid.New().String()
	profID1          = uuid.New().String()
	profID2          = uuid.New().String()
	listParams       = &app.ListParams{
		Offset: 0,
		Limit:  5,
	}
)

func init() { testinit.Setup(2, setupIntegration) }

func setupIntegration() {
	const dbSuffix = "dal"
	const migrationDir = "../../migration"

	dbCfg := pqx.Config{
		DBName:                          def.DBName,
		User:                            def.DBUser,
		Pass:                            def.DBPass,
		Host:                            def.DBHost,
		Port:                            def.DBPort,
		ConnectTimeout:                  3 * def.TestSecond,
		StatementTimeout:                3 * def.TestSecond,
		LockTimeout:                     3 * def.TestSecond,
		IdleInTransactionSessionTimeout: 3 * def.TestSecond,
		SSLMode:                         pqx.SSLDisable,
	}
	_, cleanup, err := pqx.EnsureTempDB(structlog.New(), dbSuffix, dbCfg)
	if err != nil {
		testinit.Fatal(err)
	}
	testinit.Teardown(cleanup)
	dbCfg.DBName += "_" + dbSuffix

	testRepo, err = New(ctx, dbCfg, migrationDir, false)
	if err != nil {
		testinit.Fatal(err)
	}
	testinit.Teardown(testRepo.Close)
}

func (a *Repo) truncate() error {
	_, err := a.db.Exec("TRUNCATE permissions, roles, users, tokens, sessions, wash_servers, wash_sessions RESTART IDENTITY CASCADE")
	return err
}
