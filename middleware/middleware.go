package tokens

import (
	"github.com/gin-gonic/gin"
)

// // RefreshMiddleware
// // Обновление refresh и access токенов
func (th *TokenHandler) RefreshMiddleware(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken := c.Param("refresh")
		guid := c.Param("guid")
		if refreshToken == "" {
			respondWithError(c, 401, "Invalid API token")
			return
		}
		ok, err := th.T.checkDb(refreshToken, guid)
		if err != nil {
			respondWithError(c, 500, "Server error")
			return
		}

		if ok {
			next(c)
		} else {
			respondWithError(c, 401, "Invalid API token")
			return
		}
	}
}
