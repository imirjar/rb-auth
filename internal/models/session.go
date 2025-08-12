package models

import "time"

type Session struct {
	ID      int
	User    User
	Created time.Time
	Updated time.Time
	// more data
}
