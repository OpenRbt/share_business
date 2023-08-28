package wallets

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gocraft/dbr/v2"
	"github.com/gocraft/dbr/v2/dialect"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestWalletRepo_GetById(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	conn := &dbr.Connection{
		DB:            db,
		EventReceiver: &dbr.NullEventReceiver{},
		Dialect:       dialect.PostgreSQL,
	}

	repo := NewRepo(zap.NewNop().Sugar(), conn)

	testUUID := uuid.NewV4()

	rows := sqlmock.NewRows([]string{"id", "user_id", "organization_id", "is_default", "balance"}).
		AddRow(testUUID, "testUserID", "testOrgID", true, "100.5")

	mock.ExpectQuery("SELECT * FROM wallets WHERE (id = ?)").
		WithArgs(testUUID).
		WillReturnRows(rows)

	wallet, err := repo.GetById(ctx, testUUID)

	assert.NoError(t, err)
	assert.Equal(t, testUUID, wallet.ID)
	assert.Equal(t, "testUserID", wallet.UserID)

	assert.NoError(t, mock.ExpectationsWereMet())
}
