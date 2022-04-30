package global

import (
	"github.com/wader/gormstore/v2"
)

const SessionKey = "session"

var SessionStore *gormstore.Store
