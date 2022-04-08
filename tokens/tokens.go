package tokens

import (
	"encoding/base64"
	"time"

	"testMEDOS/users"

	"github.com/dgrijalva/jwt-go"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Tokens struct {
	db        *mongo.Client
	SecretKey string
}

func NewToken(db *mongo.Client, secretKey string) *Tokens {
	return &Tokens{
		db:        db,
		SecretKey: secretKey,
	}
}

type AuthToken struct {
	RefreshToken string `json:"refresh_token"`
	Token        string `json:"access_token"`
}

type AuthTokenClaim struct {
	*jwt.StandardClaims
	User users.User
}

// generateTokens ...
// Генерация refresh и access токенов
func (t *Tokens) generateTokens(guid string) (AuthToken, users.User, error) {
	var u users.User
	u.Guid = guid
	expiresAt := time.Now().Add(time.Minute * 1).Unix()

	token := jwt.New(jwt.SigningMethodHS512) // SHA512

	token.Claims = &AuthTokenClaim{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		u,
	}

	tokenString, err := token.SignedString([]byte(t.SecretKey))

	if err != nil {
		return AuthToken{}, u, err
	}

	refreshToken := uuid.New()

	refreshTokenBase64 := base64.StdEncoding.EncodeToString([]byte(refreshToken.String()))

	return AuthToken{
		Token:        tokenString,
		RefreshToken: refreshTokenBase64,
	}, u, nil
}
