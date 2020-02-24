package utils

import (
	"sync"
	"time"

	memcache "github.com/patrickmn/go-cache"
	uuid "github.com/satori/go.uuid"
)

// CookieFunction define cookie
type CookieFunction interface {
	NewCookie()
	SetSession()
	SetToken()
	IsSession()
	IsToken()
}

// CookieUtils keep date in memory
type CookieUtils struct {
	*sync.RWMutex // Read & Write locker

	session *memcache.Cache // keep cookie in a period time
	token   *memcache.Cache // avoiding multiple submit form action
}

/////////////////// Setter & Getter //////////////////

// NewCookie create new in memory cookie storage
// save Cookie by default time
func NewCookie(defaultExpiration time.Duration) *CookieUtils {
	return &CookieUtils{
		session: memcache.New(defaultExpiration, -1),
		token:   memcache.New(defaultExpiration, 3*time.Hour)}
}

// SetSession generate client uuid and store in memory
// return session for setting cookie
func (c *CookieUtils) SetSession() string {
	sessionID := uuid.Must(uuid.NewV4()).String()
	c.session.SetDefault(sessionID, "1")
	return sessionID
}

// SetToken set token, use once only
// return token for setting cookie
func (c *CookieUtils) SetToken() string {
	token := uuid.Must(uuid.NewV4()).String()
	c.token.SetDefault(token, "2")
	return token
}

/////////////////// Main //////////////////

// IsSession check sessionID
func (c *CookieUtils) IsSession(sessionID string) bool {
	_, found := c.session.Get(sessionID)
	return found
}

// IsToken check if is token then delete it
func (c *CookieUtils) IsToken(token string) bool {
	_, found := c.token.Get(token)
	if found {
		c.token.Delete(token)
	}
	return found
}
