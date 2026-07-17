package provider

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// defaultListSessionTTL is how long an unused provider List continuation token remains valid.
const defaultListSessionTTL = 5 * time.Minute

// listSession stores provider-owned state for one List continuation token.
type listSession struct {
	cloudControlNextToken string
	bufferedIDs           []string
	returned              int64
	limit                 int64
	resourceToken         string
	queryHash             string
	lastAccess            time.Time
	inUse                 bool
}

// nextTokenPtr returns the CloudControl continuation token for the next ListResources request.
func (s *listSession) nextTokenPtr() *string {
	if s.cloudControlNextToken == "" {
		return nil
	}
	return &s.cloudControlNextToken
}

// remainingLimit returns the caller limit still available across this List session.
func (s *listSession) remainingLimit() *int64 {
	if s.limit <= 0 {
		return nil
	}
	remaining := s.limit - s.returned
	return &remaining
}

// listSessionStore owns provider continuation tokens for in-flight List sessions.
type listSessionStore struct {
	mu       sync.Mutex
	sessions map[string]*listSession
	ttl      time.Duration
	now      func() time.Time
}

// newListSessionStore creates a List continuation session store with the supplied token TTL.
func newListSessionStore(ttl time.Duration) *listSessionStore {
	return &listSessionStore{
		sessions: map[string]*listSession{},
		ttl:      ttl,
		now:      time.Now,
	}
}

// put stores a List session behind a newly generated provider continuation token.
func (s *listSessionStore) put(session *listSession) (string, error) {
	token, err := newListContinuationToken()
	if err != nil {
		return "", err
	}

	s.mu.Lock()
	defer s.mu.Unlock()
	now := s.now()
	s.cleanupExpiredLocked(now)
	session.lastAccess = now
	session.inUse = false
	s.sessions[token] = session
	return token, nil
}

// get locks a List session for one continuation request and returns a release callback.
func (s *listSessionStore) get(
	token, resourceToken, queryHash string, limit int64,
) (*listSession, func(), error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := s.now()
	s.cleanupExpiredLocked(now)
	session, ok := s.sessions[token]
	if !ok {
		return nil, nil, status.Error(codes.InvalidArgument, "invalid or expired continuation_token")
	}
	if session.inUse {
		return nil, nil, status.Error(codes.FailedPrecondition, "continuation_token is already in use")
	}
	if session.resourceToken != resourceToken {
		return nil, nil, status.Error(
			codes.InvalidArgument,
			"continuation_token was created for a different resource type",
		)
	}
	if session.queryHash != queryHash {
		return nil, nil, status.Error(codes.InvalidArgument, "continuation_token was created for a different query")
	}
	if session.limit != limit {
		return nil, nil, status.Error(codes.InvalidArgument, "continuation_token was created for a different limit")
	}
	session.inUse = true
	session.lastAccess = now

	release := func() {
		s.mu.Lock()
		defer s.mu.Unlock()
		if current, ok := s.sessions[token]; ok && current == session {
			current.inUse = false
			current.lastAccess = s.now()
		}
	}
	return session, release, nil
}

// remove deletes a provider continuation token and its List session.
func (s *listSessionStore) remove(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessions, token)
}

// cleanupExpiredLocked removes unused expired sessions while the store mutex is held.
func (s *listSessionStore) cleanupExpiredLocked(now time.Time) {
	if s.ttl <= 0 {
		return
	}
	for token, session := range s.sessions {
		if !session.inUse && now.Sub(session.lastAccess) > s.ttl {
			delete(s.sessions, token)
		}
	}
}

// newListContinuationToken returns an opaque provider continuation token.
func newListContinuationToken() (string, error) {
	var buf [16]byte
	if _, err := rand.Read(buf[:]); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf[:]), nil
}
