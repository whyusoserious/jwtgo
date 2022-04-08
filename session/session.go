package tokens

import (
	"testMEDOS/users"
)

type Session struct {
	RefreshToken string
	User         users.User
}
