package connection

import (
	cache2 "github.com/patrickmn/go-cache"
	"wash_bonus/internal/entity"
)

func (c *cache) Set(apiKey string, connection entity.WashServerConnection) {
	c.cache.Set(apiKey, &connection, cache2.DefaultExpiration)
}

func (c *cache) Get(apiKey string) *entity.WashServerConnection {
	connection, ok := c.cache.Get(apiKey)
	if !ok {
		return nil
	}
	return connection.(*entity.WashServerConnection)
}

func (c *cache) Refresh(apiKey string, connection entity.WashServerConnection) error {
	return c.cache.Replace(apiKey, &connection, cache2.DefaultExpiration)
}
