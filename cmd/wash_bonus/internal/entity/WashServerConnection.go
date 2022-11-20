package entity

type WashServerConnection struct {
	WashServer WashServer

	Sessions map[string]*Session
}
