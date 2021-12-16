package api

import (
	"sync"

	"github.com/google/uuid"
	"github.com/yanando/StudentLink/datamanager"
)

type SessionManager struct {
	sessionsLock sync.Mutex
	sessions     map[string]*datamanager.User
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		sessionsLock: sync.Mutex{},
		sessions:     make(map[string]*datamanager.User),
	}
}

func (sManager *SessionManager) GetUserBySessionID(id string) (*datamanager.User, bool) {
	sManager.sessionsLock.Lock()
	user, contained := sManager.sessions[id]
	sManager.sessionsLock.Unlock()
	return user, contained
}

func (sManager *SessionManager) CreateSession(user *datamanager.User) string {
	sessionId := uuid.New().String()
	sManager.sessionsLock.Lock()
	sManager.sessions[sessionId] = user
	sManager.sessionsLock.Unlock()
	return sessionId
}
