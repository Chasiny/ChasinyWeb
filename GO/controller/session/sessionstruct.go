package session

import "sync"

type UserSession struct {
	Name string
	Level int64
	Sid string
	CreateTime int64
}

type SessionManager struct {
	Key string
	SessionTime int64
	Count int64
	SessionLock sync.Locker
}

var sessionMap map[string]UserSession