// Code generated by mtgroup-generator.
package dal

import (
	"database/sql"
	"time"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

const (
	sqlGetWashServer = `
	SELECT
		id,
		created_at,
		key,
		last_update_at,
		modified_at,
		name
	FROM
		wash_servers
	WHERE
		id=:id AND
		isolated_entity_id=:isolated_entity_id AND
		NOT deleted
	`

	sqlGetMyWashServerID = `
	SELECT
		id
	FROM
		wash_servers
	WHERE
		created_by=:created_by AND
		isolated_entity_id=:isolated_entity_id AND
		NOT deleted AND
		bound
	`

	sqlAddWashServer = `
	INSERT INTO wash_servers(
		created_at,
		key,
		last_update_at,
		modified_at,
		name,
		id,
		created_by,
		isolated_entity_id
	) VALUES (
		:created_at,
		:key,
		:last_update_at,
		:modified_at,
		:name,
		:id,
		:created_by,
		:isolated_entity_id
	)
	RETURNING
		id
	`

	sqlBindWashServerToProfile = `
	UPDATE
		wash_servers
	SET
		bound=true
	WHERE
		id=:id AND
		created_by=:created_by AND
		isolated_entity_id=:isolated_entity_id AND
		NOT deleted AND
		NOT bound
	`

	sqlDeleteWashServer = `
	UPDATE
		wash_servers
	SET
		deleted=true,
		deleted_at=:deleted_at,
		deleted_by=:deleted_by
	WHERE
		id=:id AND
		isolated_entity_id=:isolated_entity_id AND
		NOT deleted
	`

	sqlEditWashServer = `
	UPDATE
		wash_servers
	SET
		key=:key,
		last_update_at=:last_update_at,
		modified_at=:modified_at,
		name=:name
	WHERE
		id=:id AND
		isolated_entity_id=:isolated_entity_id AND
		NOT deleted
	`
	sqlSetModifiedParamsWashServer = `
	UPDATE
		wash_servers
	SET
		modified_at=:modified_at
		
	WHERE
		id=:id AND
		isolated_entity_id=:isolated_entity_id
	`

	sqlListWashServer = `
	SELECT
		id,
		created_at,
		key,
		last_update_at,
		modified_at,
		name
	FROM
		wash_servers
	WHERE
		isolated_entity_id=:isolated_entity_id AND
		NOT deleted
	`
	sqlListWashServerCount = `
	SELECT
		COUNT(*)
	FROM
		wash_servers
	WHERE
		isolated_entity_id=:isolated_entity_id AND
		NOT deleted
	`
)

type (
	argGetWashServer struct {
		ID               sql.NullString `db:"id"`
		IsolatedEntityID string         `db:"isolated_entity_id"`
	}

	argGetMyWashServerID struct {
		CreatedBy        string `db:"created_by"`
		IsolatedEntityID string `db:"isolated_entity_id"`
	}
	argAddWashServer struct {
		ID               string     `db:"id"`
		CreatedAt        *time.Time `db:"created_at"`
		Key              string     `db:"key"`
		LastUpdateAt     *time.Time `db:"last_update_at"`
		ModifiedAt       *time.Time `db:"modified_at"`
		Name             string     `db:"name"`
		CreatedBy        string     `db:"created_by"`
		IsolatedEntityID string     `db:"isolated_entity_id"`
	}
	argBindWashServerToProfile struct {
		ID               string `db:"id"`
		CreatedBy        string `db:"created_by"`
		IsolatedEntityID string `db:"isolated_entity_id"`
	}
	argEditWashServer struct {
		ID               string     `db:"id"`
		CreatedAt        *time.Time `db:"created_at"`
		Key              string     `db:"key"`
		LastUpdateAt     *time.Time `db:"last_update_at"`
		ModifiedAt       *time.Time `db:"modified_at"`
		Name             string     `db:"name"`
		IsolatedEntityID string     `db:"isolated_entity_id"`
	}
	argDeleteWashServer struct {
		ID               string     `db:"id"`
		DeletedAt        *time.Time `db:"deleted_at"`
		DeletedBy        string     `db:"deleted_by"`
		IsolatedEntityID string     `db:"isolated_entity_id"`
	}
	argSetModifiedParamsWashServer struct {
		ID               string     `db:"id"`
		ModifiedAt       *time.Time `db:"modified_at"`
		IsolatedEntityID string     `db:"isolated_entity_id"`
	}
)
