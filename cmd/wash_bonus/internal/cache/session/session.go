package session

import (
	cache2 "github.com/patrickmn/go-cache"
	uuid "github.com/satori/go.uuid"
	"wash_bonus/internal/entity"
)

func (c *cache) GetSession(sessionID uuid.UUID) *entity.Session {
	session, ok := c.cache.Get(sessionID.String())
	if !ok {
		return nil
	}
	return session.(*entity.Session)
}

func (c *cache) SetSession(session entity.Session) {
	c.cache.Set(session.ID.String(), &session, cache2.DefaultExpiration)
}

func (c *cache) RefreshSession(session entity.Session) error {
	return c.cache.Replace(session.ID.String(), &session, cache2.DefaultExpiration)
}
