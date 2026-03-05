package session

import (
	"context"
	"sync"
	"time"
)

// Session represents a client session
type Session struct {
	ID         string
	ClientInfo map[string]interface{}
	State      string // "initialized", "active", "closed"
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Metadata   map[string]interface{}
}

// Manager manages sessions
type Manager struct {
	sessions map[string]*Session
	mu      sync.RWMutex
}

// NewManager creates a new session manager
func NewManager() *Manager {
	return &Manager{
		sessions: make(map[string]*Session),
	}
}

// Create creates a new session
func (m *Manager) Create(id string, clientInfo map[string]interface{}) *Session {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	session := &Session{
		ID:         id,
		ClientInfo: clientInfo,
		State:      "initialized",
		CreatedAt:  now,
		UpdatedAt:  now,
		Metadata:   make(map[string]interface{}),
	}
	m.sessions[id] = session
	return session
}

// Get retrieves a session by ID
func (m *Manager) Get(id string) (*Session, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	session, ok := m.sessions[id]
	return session, ok
}

// Update updates session state
func (m *Manager) Update(id string, state string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if session, ok := m.sessions[id]; ok {
		session.State = state
		session.UpdatedAt = time.Now()
		return true
	}
	return false
}

// Delete removes a session
func (m *Manager) Delete(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, ok := m.sessions[id]; ok {
		delete(m.sessions, id)
		return true
	}
	return false
}

// List returns all sessions
func (m *Manager) List() []*Session {
	m.mu.RLock()
	defer m.mu.RUnlock()

	sessions := make([]*Session, 0, len(m.sessions))
	for _, s := range m.sessions {
		sessions = append(sessions, s)
	}
	return sessions
}

// Cleanup removes expired sessions
func (m *Manager) Cleanup(maxAge time.Duration) int {
	m.mu.Lock()
	defer m.mu.Unlock()

	now := time.Now()
	deleted := 0

	for id, s := range m.sessions {
		if now.Sub(s.UpdatedAt) > maxAge {
			delete(m.sessions, id)
			deleted++
		}
	}
	return deleted
}

// Count returns the number of sessions
func (m *Manager) Count() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.sessions)
}

// ContextKey for session ID
type ContextKey string

const SessionIDKey ContextKey = "session_id"

// NewContext creates a context with session ID
func NewContext(ctx context.Context, sessionID string) context.Context {
	return context.WithValue(ctx, SessionIDKey, sessionID)
}

// FromContext gets session ID from context
func FromContext(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(SessionIDKey).(string)
	return id, ok
}
