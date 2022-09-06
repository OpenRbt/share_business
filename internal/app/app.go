// Code generated by mtgroup-generator.
package app

// Make sure not to overwrite this file after you generated it because all your edits would be lost!

// App provides application features service.
type App interface {
            GetPermission(prof Profile, id string) (*Permission, error)
            AddPermission(prof Profile, m *Permission) (*Permission, error)
            EditPermission(prof Profile, id string, m *Permission) error
            DeletePermission(prof Profile, id string) error
            ListPermission(prof Profile, params *ListParams) ([]*Permission, []string, error)
	
            GetRole(prof Profile, id string) (*Role, error)
            AddRole(prof Profile, m *Role) (*Role, error)
            EditRole(prof Profile, id string, m *Role) error
            DeleteRole(prof Profile, id string) error
            ListRole(prof Profile, params *ListParams) ([]*Role, []string, error)
	
    AddPermissionsRole(id string, isolatedEntityID string, itemsID []string, items []*Permission) error
    DeletePermissionsRole(id string, isolatedEntityID string, items []string) error
            GetUser(prof Profile, id string) (*User, error)
            AddUser(prof Profile, m *User) (*User, error)
            EditUser(prof Profile, id string, m *User) error
            DeleteUser(prof Profile, id string) error
            ListUser(prof Profile, params *ListParams) ([]*User, []string, error)
	
            GetWashServer(prof Profile, id string) (*WashServer, error)
            AddWashServer(prof Profile, m *WashServer) (*WashServer, error)
            EditWashServer(prof Profile, id string, m *WashServer) error
            DeleteWashServer(prof Profile, id string) error
            ListWashServer(prof Profile, params *ListParams) ([]*WashServer, []string, error)
	

    
        AddTestData(prof Profile) error
    
}

// Repo interface for data repository
type Repo interface {
            GetPermission(id string, isolatedEntityID string) (*Permission, error)
            AddPermission(profileID string, isolatedEntityID string, m *Permission) (*Permission, error)
            EditPermission(id string, isolatedEntityID string, m *Permission) error
            DeletePermission(id string, profileID string, isolatedEntityID string) error
            ListPermission(isolatedEntityID string, params *ListParams) ([]*Permission, []string, error)
	
            GetRole(id string, isolatedEntityID string) (*Role, error)
            AddRole(profileID string, isolatedEntityID string, m *Role) (*Role, error)
            EditRole(id string, isolatedEntityID string, m *Role) error
            DeleteRole(id string, profileID string, isolatedEntityID string) error
            ListRole(isolatedEntityID string, params *ListParams) ([]*Role, []string, error)
	
    AddPermissionsRole(id string, isolatedEntityID string, itemsID []string, items []*Permission) error
    DeletePermissionsRole(id string, isolatedEntityID string, items []string) error
            GetUser(id string, isolatedEntityID string) (*User, error)
            AddUser(profileID string, isolatedEntityID string, m *User) (*User, error)
            EditUser(id string, isolatedEntityID string, m *User) error
            DeleteUser(id string, profileID string, isolatedEntityID string) error
            ListUser(isolatedEntityID string, params *ListParams) ([]*User, []string, error)
	
            GetWashServer(id string, isolatedEntityID string) (*WashServer, error)
            AddWashServer(profileID string, isolatedEntityID string, m *WashServer) (*WashServer, error)
            EditWashServer(id string, isolatedEntityID string, m *WashServer) error
            DeleteWashServer(id string, profileID string, isolatedEntityID string) error
            ListWashServer(isolatedEntityID string, params *ListParams) ([]*WashServer, []string, error)
	

    

    AddTestData(profileID, isolatedEntityID string) error
}


    type ListParams struct {
        Offset       int64
        Limit        int64
        FilterGroups []*FilterGroup
        SortBy       string
        OrderBy      string
    }

    type FilterGroup struct {
        Key         string
        LogicFilter bool
        Filters     []*Filter
    }

    type Filter struct {
        Value      string
        Operator   string
        IgnoreCase bool
    }


type app struct {
	repo     Repo
	rulesSet RulesSet
}

func New(r Repo, rs RulesSet) (App) {
	return &app{
		repo: r,
        rulesSet: rs,
	}
}
    func (a *app) AddTestData(prof Profile) error {
        if !prof.Authz.Admin {
            return ErrAccessDenied
        }
        return a.repo.AddTestData(prof.ID, prof.IsolatedEntityID)
    }
