package connection

import (
	cache2 "github.com/patrickmn/go-cache"
	"go.uber.org/zap"
	"time"
)

type cache struct {
	cache *cache2.Cache
	l     *zap.SugaredLogger
}

func New(l *zap.SugaredLogger) *cache {
	c := cache2.New(15*time.Minute, 1*time.Hour)
	return &cache{
		l:     l,
		cache: c,
	}
}
