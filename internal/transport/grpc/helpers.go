package grpc

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (s *Service) createConnectionIfMotExist(key string) error {
	s.connectionsMutex.RLock()
	if _, ok := s.connections[key]; ok {
		return nil
	}
	s.connectionsMutex.RUnlock()

	server, err := s.repo.GetWashServerByKey(key)
	if err != nil {
		return err
	}

	newConnection := Connection{
		Valid:    true,
		Server:   *server,
		Sessions: make(map[string]*Session, 12),
	}

	s.connectionsMutex.Lock()
	s.connections[key] = &newConnection
	s.connectionsMutex.Unlock()

	return nil
}

func (s *Service) getConnection(key string) (*Connection, error) {
	s.connectionsMutex.RLock()
	val, ok := s.connections[key]
	s.connectionsMutex.RUnlock()

	if !ok {
		return nil, ErrNotFound
	}

	return val, nil
}

func (c *Connection) createSessionIfNotExist(id string) error {
	uid, err := uuid.FromString(id)
	if err != nil {
		return ErrBadFormat
	}

	c.sessionsMutex.RLock()
	_, ok := c.Sessions[id]
	c.sessionsMutex.RUnlock()
	if ok {
		return ErrSessionAlreadyExists
	}

	c.sessionsMutex.Lock()
	c.Sessions[id] = &Session{
		ID:            uid,
		PostID:        0,
		User:          nil,
		Amount:        decimal.Decimal{},
		ConsumeAmount: decimal.Decimal{},
		Processed:     false,
	}
	c.sessionsMutex.Unlock()

	return nil
}

func (c *Connection) getSession(id string) (*Session, error) {
	_, err := uuid.FromString(id)
	if err != nil {
		return nil, ErrBadFormat
	}

	c.sessionsMutex.RLock()
	val, ok := c.Sessions[id]
	c.sessionsMutex.RUnlock()
	if !ok {
		return nil, ErrNotFound
	}

	return val, nil
}

func (c *Connection) updateSession(id string, enteredAmount string) error {
	_, err := uuid.FromString(id)
	if err != nil {
		return ErrBadFormat
	}
	amount, err := decimal.NewFromString(enteredAmount)
	if err != nil {
		return ErrBadAmount
	}

	c.sessionsMutex.RLock()
	val, ok := c.Sessions[id]
	c.sessionsMutex.RUnlock()
	if !ok {
		return ErrNotFound
	}

	c.sessionsMutex.Lock()
	val.Amount = amount
	c.sessionsMutex.Unlock()

	return nil
}

func (c *Connection) processSession(id string) error {
	_, err := uuid.FromString(id)
	if err != nil {
		return ErrBadFormat
	}

	c.sessionsMutex.RLock()
	val, ok := c.Sessions[id]
	c.sessionsMutex.RUnlock()
	if !ok {
		return ErrNotFound
	}

	c.sessionsMutex.Lock()
	val.Processed = true
	c.sessionsMutex.Unlock()

	return nil
}

func (c *Connection) finishSession(id string) error {
	c.sessionsMutex.Lock()
	defer c.sessionsMutex.Unlock()

	if _, ok := c.Sessions[id]; ok {
		delete(c.Sessions, id)
	}

	return nil
}
