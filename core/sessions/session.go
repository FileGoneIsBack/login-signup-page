package sessions

import (
	"login/core/database"
	"time"
)

var (
	Sessions = map[string]Session{}
)

// Session is used to store the user & expiry
type Session struct {
	*database.User
	Expiry time.Time
}

func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}

func Count() int {
    return len(Sessions)
}
