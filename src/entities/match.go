//match entities
package entities

import (
	"time"
)

type Match struct {
	MatchID          int64
	MatchDescription string
	CreatedTime      time.Time
	UpdatedTime      time.Time
	IsDeleted        bool
}
