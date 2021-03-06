package utils

// 生成token, 解析token
import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/hero1s/ginweb/i18n"
	"github.com/hero1s/ginweb/pkg/log"
	"net/http"
	"time"
)

var tokenSecret string
var (
	TokenExpired          = errors.New("token is expired")
	TokenUsedBeforeIssued = errors.New("token used before issued")
	TokenNotValidYet      = errors.New("token is not valid yet")
	TokenInvalid          = errors.New("token invalid")
)

var TokenExporeTime = 500 //10小时过期

func SetTokenSecretKey(secretKey string, exporeHour int) {
	tokenSecret = secretKey
	TokenExporeTime = exporeHour
}

type CustomClaims struct {
	jwt.StandardClaims
	Uid      uint64 `json:"uid"`
	RoleId   uint64 `json:"role_id"`
	UserData string `json:"user_data"`
}

// generate token
func GenerateToken(claims CustomClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	// 1小时过期
	claims.ExpiresAt = time.Now().Add(time.Duration(TokenExporeTime) * time.Hour).Unix()
	claims.IssuedAt = time.Now().Unix()
	token.Claims = claims
	return token.SignedString([]byte(tokenSecret))
}

// decode token
func DecodeTokenByStr(tokenStr string) (*CustomClaims, error) {
	claims := &CustomClaims{}
	parser := &jwt.Parser{}
	token, err := parser.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		log.Debug("解析token失败:%v", err.Error())
		return nil, wrapError(err)
	}
	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, wrapError(TokenInvalid)
}

// 函数在base_func.go
func DecodeToken(r *http.Request) (*CustomClaims, error) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(tokenSecret), nil
	}, request.WithClaims(&CustomClaims{}))
	if err != nil {
		log.Debug("解析token失败:%v", err.Error())
		return nil, wrapError(err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, wrapError(TokenInvalid)
}

// refresh token
func RefreshToken(r *http.Request) (string, error) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(tokenSecret), nil
	}, request.WithClaims(&CustomClaims{}))
	if err != nil {
		log.Debug("解析token失败:%v", err.Error())
		return "", wrapError(err)
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.ExpiresAt = time.Now().Add(time.Duration(TokenExporeTime) * time.Hour).Unix()
		return GenerateToken(*claims)
	}
	return "", wrapError(TokenInvalid)
}

// Validates time based claims "exp, iat, nbf".
// There is no accounting for clock skew.
// As well, if any of the above claims are not in the token, it will still
// be considered a valid claim.
func (c CustomClaims) Valid() error {
	vErr := new(jwt.ValidationError)
	now := jwt.TimeFunc().Unix()

	// The claims below are optional, by default, so if they are set to the
	// default value in Go, let's not fail the verification for them.
	if c.VerifyExpiresAt(now, false) == false {
		//delta := time.Unix(now, 0).Sub(time.Unix(c.ExpiresAt, 0))
		// vErr.Inner = fmt.Errorf("token is expired by %v", delta)
		vErr.Inner = TokenExpired
		vErr.Errors |= jwt.ValidationErrorExpired
	}

	if c.VerifyIssuedAt(now, false) == false {
		vErr.Inner = TokenUsedBeforeIssued
		vErr.Errors |= jwt.ValidationErrorIssuedAt
	}

	if c.VerifyNotBefore(now, false) == false {
		vErr.Inner = TokenNotValidYet
		vErr.Errors |= jwt.ValidationErrorNotValidYet
	}

	if vErr.Errors == 0 {
		return nil
	}
	return vErr
}

func wrapError(err error) error {
	if err == nil {
		return nil
	}
	switch err.Error() {
	case TokenInvalid.Error(), TokenNotValidYet.Error(), TokenUsedBeforeIssued.Error():
		return i18n.TokenInvalid
	case TokenExpired.Error():
		return i18n.TokenExpired
	}
	return i18n.Unauthorized
}
