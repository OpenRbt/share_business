package app

type AccessRule func(user WashUser) bool

func checkAccess(user WashUser, rule AccessRule) bool {
	return rule(user)
}
func allRules(accessRules ...AccessRule) AccessRule {
	return func(user WashUser) bool {
		var ok bool

		for _, rule := range accessRules {
			ok = rule(user)
			if !ok {
				return false
			}
		}

		return true
	}
}

func anyRule(accessRule ...AccessRule) AccessRule {
	return func(user WashUser) bool {
		var ok bool

		for _, rule := range accessRule {
			ok = rule(user)
			if ok {
				return ok
			}
		}

		return false
	}
}

func roleAdmin(user WashUser) bool {
	return user.Role == AdminRole
}

func roleEngineer(user WashUser) bool {
	return user.Role == EngineerRole
}

func userIsOwner(server WashServer) AccessRule {
	return func(user WashUser) bool {
		return server.Owner == user.ID
	}
}
