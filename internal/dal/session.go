// Code generated by mtgroup-generator.
package dal

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"wash-bonus/internal/app"
)

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

type Session struct {
	ID           uuid.UUID      `db:"id"`
	Active       sql.NullBool   `db:"active"`
	ClosingAt    *time.Time     `db:"closing_at"`
	CreatedAt    *time.Time     `db:"created_at"`
	ExpirationAt *time.Time     `db:"expiration_at"`
	UpdateAt     *time.Time     `db:"update_at"`
	UserID       sql.NullString `db:"user_id"`
	User         Token
}

var SessionProps = map[string]columnProps{
	"active": {
		sqlName:  "active",
		typeName: "bool",
	},
	"closingAt": {
		sqlName:  "closing_at",
		typeName: "date-time",
	},
	"createdAt": {
		sqlName:  "created_at",
		typeName: "date-time",
	},
	"expirationAt": {
		sqlName:  "expiration_at",
		typeName: "date-time",
	},
	"id": {
		sqlName:  "id",
		typeName: "uuid",
	},
	"updateAt": {
		sqlName:  "update_at",
		typeName: "date-time",
	},
	"user": {
		sqlName:  "user_id",
		typeName: "uuid",
	},
}

func (a *Repo) GetSession(id string, isolatedEntityID string) (*app.Session, error) {
	return a.getSession(id, isolatedEntityID)
}

func (a *Repo) AddSession(profileID string, isolatedEntityID string, m *app.Session) (*app.Session, error) {
	id, err := a.addSession(profileID, isolatedEntityID, m)
	if err != nil {
		return nil, err
	}
	return a.getSession(id, isolatedEntityID)
}

func (a *Repo) EditSession(id string, isolatedEntityID string, m *app.Session) error {
	if err := a.editSession(id, isolatedEntityID, m); err != nil {
		return err
	}

	return nil
}

func (a *Repo) DeleteSession(id string, profileID string, isolatedEntityID string) error {
	t := time.Now()
	res, err := a.db.NamedExec(sqlDeleteSession, argDeleteSession{
		ID:               id,
		DeletedAt:        &t,
		DeletedBy:        profileID,
		IsolatedEntityID: isolatedEntityID,
	})
	if err != nil {
		return err
	}
	if count, _ := res.RowsAffected(); count == 0 {
		return app.ErrNotFound
	}

	return nil
}

func (a *Repo) ListSession(isolatedEntityID string, params *app.ListParams) ([]*app.Session, []string, error) {
	ms := []Session{}
	warnings := []string{}

	var orderQuery string
	switch params.SortBy {
	case "":
	default:
		warnings = append(warnings, fmt.Sprintf("Sorting by '%s' is not avaliable or '%s' is not a valid sort key", params.SortBy, params.SortBy))
	}

	if orderQuery != "" {
		switch params.OrderBy {
		case "ASC", "":
			orderQuery += " ASC"
		case "DESC":
			orderQuery += " DESC"
		}
	}

	bf := newBuilderFilter(params.FilterGroups, SessionProps)

	sqlFilters, namedVars, warningsFromPrepared := bf.preparedSQLFilters()
	warnings = append(warnings, warningsFromPrepared...)

	namedVars["isolated_entity_id"] = isolatedEntityID

	var offset, limit string
	var err error

	nestedFilterGroups := bf.nestedFilterGroups()

	externalPagination := false
	if len(nestedFilterGroups) != 0 {
		externalPagination = true
	}
	if !externalPagination {
		offset = " OFFSET :offset"
		namedVars["offset"] = params.Offset
		if params.Limit != 0 {
			limit = " LIMIT :limit"
			namedVars["limit"] = params.Limit
		}
	}

	err = a.db.NamedSelect(&ms, sqlListSession+sqlFilters+orderQuery+offset+limit, namedVars)
	if err != nil {
		return nil, nil, err
	}

	result := []Session{}
	for i := range ms {
		if err := ms[i].LazyLoading(isolatedEntityID, a); err != nil {
			return nil, nil, err
		}

		ok := true
		for j, filterGroup := range nestedFilterGroups {
			for _, filter := range filterGroup.Filters {
				var validFilter error
				ok, validFilter = ms[i].NestedFilter(filterGroup.Key, filter)
				if validFilter != nil {
					warnings = append(warnings, fmt.Sprintf("Filter key: '%s'. Error: %s", filterGroup.Key, validFilter.Error()))
					nestedFilterGroups = append(nestedFilterGroups[:j], nestedFilterGroups[j+1:]...)
					j--
				}
				if (!ok && filterGroup.LogicFilter) || (ok && !filterGroup.LogicFilter) {
					break
				}
			}
		}
		if ok {
			result = append(result, ms[i])
		}
	}

	if externalPagination {
		start, end := pagination(int(params.Offset), int(params.Limit), len(result))
		result = result[start:end]
	}

	return appSessions(result), warnings, nil
}

func (m *Session) LazyLoading(isolatedEntityID string, a *Repo) (err error) {
	if err = a.db.NamedGet(&m.User, sqlGetUserForSessionLazyLoading, argGetToken{
		ID:               m.UserID,
		IsolatedEntityID: isolatedEntityID,
	}); err != nil && err != sql.ErrNoRows {
		return
	}
	return nil
}

func (a *Repo) getSession(id string, isolatedEntityID string) (*app.Session, error) {
	var m Session
	if err := a.db.NamedGet(&m, sqlGetSession, argGetSession{
		ID:               newNullUUID(id),
		IsolatedEntityID: isolatedEntityID,
	}); err != nil {
		if err == sql.ErrNoRows {
			return nil, app.ErrNotFound
		}
		return nil, err
	}
	if err := m.LazyLoading(isolatedEntityID, a); err != nil {
		return nil, err
	}
	return appSession(m), nil
}

func (a *Repo) addSession(profileID string, isolatedEntityID string, m *app.Session) (string, error) {
	SessionID := uuid.New().String()
	t := time.Now()
	m.CreatedAt = &t
	var userID interface{}
	if m.User != nil {
		userID = m.User.ID
	}
	if err := a.db.NamedGet(&SessionID, sqlAddSession, argAddSession{
		ID:               SessionID,
		Active:           m.Active,
		ClosingAt:        m.ClosingAt,
		CreatedAt:        m.CreatedAt,
		ExpirationAt:     m.ExpirationAt,
		UpdateAt:         m.UpdateAt,
		UserID:           userID,
		CreatedBy:        profileID,
		IsolatedEntityID: isolatedEntityID,
	}); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return "", app.ErrDuplicateID
		}
		return "", err
	}
	return SessionID, nil
}

func (a *Repo) getMySessionID(profileID, isolatedEntityID string) (id string, err error) {
	if err = a.db.NamedGet(&id, sqlGetMySessionID, argGetMySessionID{
		CreatedBy:        profileID,
		IsolatedEntityID: isolatedEntityID,
	}); err != nil {
		if err == sql.ErrNoRows {
			return "", app.ErrNotFound
		}
		return
	}
	return
}

func (a *Repo) bindToProfileSession(id, profileID, isolatedEntityID string) error {
	res, err := a.db.NamedExec(sqlBindSessionToProfile, argBindSessionToProfile{
		ID:               id,
		CreatedBy:        profileID,
		IsolatedEntityID: isolatedEntityID,
	})
	if err != nil {
		return err
	}

	if count, _ := res.RowsAffected(); count == 0 {
		return app.ErrNotFound
	}
	return nil
}

func (a *Repo) editSession(id string, isolatedEntityID string, m *app.Session) error {
	var userID interface{}
	if m.User != nil {
		userID = m.User.ID
	}

	res, err := a.db.NamedExec(sqlEditSession, argEditSession{
		ID:               id,
		Active:           m.Active,
		ClosingAt:        m.ClosingAt,
		CreatedAt:        m.CreatedAt,
		ExpirationAt:     m.ExpirationAt,
		UpdateAt:         m.UpdateAt,
		UserID:           userID,
		IsolatedEntityID: isolatedEntityID,
	})
	if err != nil {
		return err
	}

	if count, _ := res.RowsAffected(); count == 0 {
		return app.ErrNotFound
	}

	return nil
}
func (m *Session) NestedFilter(key string, filter *app.Filter) (ok bool, err error) {
	if strings.Contains(key, ".") {
		splitedFilter := strings.SplitN(key, ".", 2)
		key = splitedFilter[1]
		switch splitedFilter[0] {
		case "user":
			ok, err = m.User.Filter(key, filter)
		default:
			ok, err = true, errNotExistFilterKey
		}
	} else {
		ok, err = m.Filter(key, filter)
	}
	return
}

func (m *Session) Filter(key string, filter *app.Filter) (ok bool, err error) {
	columnType := SessionProps[key].typeName
	if err = validateOperator(filter.Operator, columnType); err != nil {
		return true, err
	}
	if err = vaidateIgnoreCase(filter.IgnoreCase, columnType); err != nil {
		return true, err
	}
	if err := validateValue(filter.Value, columnType); err != nil {
		return true, err
	}
	switch key {
	case "id":
		ok = compareUUID(filter.Operator, m.ID, filter.Value)
	case "active":
		ok = compareBool(filter.Operator, m.Active.Bool, filter.Value)
	case "closingAt":
		ok = compareTime(filter.Operator, *m.ClosingAt, filter.Value)
	case "createdAt":
		ok = compareTime(filter.Operator, *m.CreatedAt, filter.Value)
	case "expirationAt":
		ok = compareTime(filter.Operator, *m.ExpirationAt, filter.Value)
	case "updateAt":
		ok = compareTime(filter.Operator, *m.UpdateAt, filter.Value)
	default:
		ok, err = true, errNotExistFilterKey
	}
	return
}

func appSession(m Session) *app.Session {
	if m.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return nil
	}
	return &app.Session{
		ID:           m.ID.String(),
		Active:       m.Active.Bool,
		ClosingAt:    m.ClosingAt,
		CreatedAt:    m.CreatedAt,
		ExpirationAt: m.ExpirationAt,
		UpdateAt:     m.UpdateAt,
		User:         appToken(m.User),
	}
}

func appSessions(ms []Session) []*app.Session {
	ams := []*app.Session{}
	for _, m := range ms {
		ams = append(ams, appSession(m))
	}

	return ams
}
