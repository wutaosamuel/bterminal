package utils

import (
	"sync"
	"time"

	memcache "github.com/patrickmn/go-cache"
	uuid "github.com/satori/go.uuid"
)

// CookieUtils keep date in memory
type CookieUtils struct {
	*sync.RWMutex // Read & Write locker

	session *memcache.Cache   // keep cookie in a period time
	token   map[string]string // avoiding multiple submit form action
}

/////////////////// Setter & Getter //////////////////

// NewCookie create new in memory cookie storage
// save Cookie by default time
func NewCookie(defaultExpiration time.Duration) *CookieUtils {
	return &CookieUtils{
		session: memcache.New(defaultExpiration, -1),
		token: make(map[string]string)}
}

// SetSession generate client uuid and store in memory
func (c *CookieUtils) SetSession() {
	sessionID := uuid.Must(uuid.NewV4()).String()
	c.session.SetDefault(sessionID, "1")
}

// SetToken set token, use once only
func (c *CookieUtils) SetToken() {
	token := uuid.Must(uuid.NewV4()).String()
	c.Lock()
	c.token[token] = "1"
	c.Unlock()
}

/////////////////// Main //////////////////

// IsSession check sessionID
func (c *CookieUtils) IsSession(sessionID string) bool {
	_, found := c.session.Get(sessionID)
	return found
}

// IsToken check if is token then delete it
func (c *CookieUtils) IsToken(token string) bool {
	c.Lock()
	_, found := c.token[token]
	if found {
		delete(c.token, token)
	}
	c.Unlock()
	return found
}
