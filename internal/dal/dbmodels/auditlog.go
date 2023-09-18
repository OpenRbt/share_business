package dbmodels

import "encoding/json"

type (
	ResourceType string

	AuditLogCreation struct {
		Resource ResourceType `db:"resource"`
		EntityID string       `db:"entity_id"`
		Action   string       `db:"action"`
		UserID   string       `db:"user_performing_action"`
		Data     *[]byte      `db:"data"`
	}
)

const (
	AdminApplicationsResource ResourceType = "admin_applications"
	AdminUsersResource        ResourceType = "admin_users"
	OrganizationsResource     ResourceType = "organizations"
	ServerGroupsResource      ResourceType = "server_groups"
	WashServersResource       ResourceType = "wash_servers"
)

func BuildAuditLog(res ResourceType, entID, action, userID string, data interface{}) (AuditLogCreation, error) {
	log := AuditLogCreation{
		Resource: res,
		EntityID: entID,
		Action:   action,
		UserID:   userID,
	}

	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			return AuditLogCreation{}, err
		}

		log.Data = &jsonData
	}

	return log, nil
}
