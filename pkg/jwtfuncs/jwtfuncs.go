package jwtfuncs

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type AccessTokenClaims struct {
	ID       int
	Username string
	IsAdmin  bool
	Exp      int64
}

var (
	ErrNoSecretInEnv    = errors.New("secret key not found in .env file")
	ErrTokenExpired     = errors.New("token is expired")
	ErrValidateToken    = errors.New("could not validate token")
	ErrSignatureInvalid = errors.New("signature is invalid")
)

// Create a new access token.
func CreateAccessToken(claims *AccessTokenClaims) (accessToken string, err error) {
	// TODO: Pass secret key as an argument with config.Config
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		err = ErrNoSecretInEnv
		return
	}

	at := jwt.New(jwt.SigningMethodHS256)
	atClaims := at.Claims.(jwt.MapClaims)
	atClaims["id"] = claims.ID
	atClaims["username"] = claims.Username
	atClaims["is_admin"] = claims.IsAdmin
	atClaims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	accessToken, err = at.SignedString([]byte(secretKey))
	return
}

// Extract claims from any token.
func extractTokenClaims(token string) (claimsMap jwt.MapClaims, err error) {
	// TODO: Pass secret key as an argument with config.Config
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		err = ErrNoSecretInEnv
		return
	}
	_, err = jwt.ParseWithClaims(token, &claimsMap, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		var validationError *jwt.ValidationError
		switch {
		case errors.As(err, &validationError):
			switch validationError.Errors {
			case jwt.ValidationErrorExpired:
				err = ErrTokenExpired
			case jwt.ValidationErrorSignatureInvalid:
				err = ErrSignatureInvalid
			default:
				err = ErrValidateToken
			}
		default:
			return
		}
	}
	return
}

// Extract claims from access token.
func ExtractAccessTokenClaims(accessToken string) (cls *AccessTokenClaims, isExpired bool, err error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		err = ErrNoSecretInEnv
		return
	}

	var claimsMap jwt.MapClaims
	claimsMap, err = extractTokenClaims(accessToken)
	if err != nil && !errors.Is(err, ErrTokenExpired) {
		return
	} else if errors.Is(err, ErrTokenExpired) {
		isExpired = true
		err = nil
	}
	cls = &AccessTokenClaims{}
	var ok bool
	var floatID, floatExp float64

	floatID, ok = claimsMap["id"].(float64)
	if !ok {
		return
	}
	cls.Username, ok = claimsMap["username"].(string)
	if !ok {
		return
	}
	cls.IsAdmin, ok = claimsMap["is_blocked"].(bool)
	if !ok {
		return
	}
	floatExp, ok = claimsMap["exp"].(float64)
	if !ok {
		return
	}
	cls.ID = int(floatID)
	cls.Exp = int64(floatExp)
	return
}
