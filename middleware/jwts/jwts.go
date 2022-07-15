package jwts

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateTokenByLong 生成token
func GenerateTokenByLong(method jwt.SigningMethod, id int64, secret string, expired int64) (string, error) {
	expiresAt := jwtExpired(expired)
	claims := jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: expiresAt},
		ID:        strconv.FormatInt(id, 10),
	}

	token := jwt.NewWithClaims(method, claims)

	return token.SignedString([]byte(secret))
}

// GenerateTokenByString 生成token
func GenerateTokenByString(method jwt.SigningMethod, subject string, secret string, expired int64) (string, error) {
	expiresAt := jwtExpired(expired)
	claims := jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: expiresAt},
		Subject:   subject,
	}

	token := jwt.NewWithClaims(method, claims)

	return token.SignedString([]byte(secret))
}

// GenerateToken 生成token
func GenerateToken(method jwt.SigningMethod, subject string, id int64, secret string, expired int64) (string, error) {
	expiresAt := jwtExpired(expired)
	claims := jwt.RegisteredClaims{
		ExpiresAt: &jwt.NumericDate{Time: expiresAt},
		Subject:   subject,
		ID:        strconv.FormatInt(id, 10),
	}

	token := jwt.NewWithClaims(method, claims)

	return token.SignedString([]byte(secret))
}

// expireAt 过期时间
func jwtExpired(expire int64) time.Time {
	expireDate := time.Now()
	if expire == 0 {
		expireDate = expireDate.Add(time.Hour * 24)
	} else {
		expireDate = expireDate.Add(time.Duration(expire) * time.Second)
	}

	return expireDate
}

// ParseToken 解析token
func ParseToken(tokenStr string, secret string) (*jwt.RegisteredClaims, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			if e.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("token expired")
			} else if e.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				return nil, errors.New("token signature invalid")
			} else if e.Errors&jwt.ValidationErrorUnverifiable != 0 {
				return nil, errors.New("token unverifiable")
			} else if e.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token malformed")
			} else {
				return nil, errors.New("token invalid")
			}
		}
	} else {
		if claim, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
			return claim, nil
		}
	}

	return nil, err
}
