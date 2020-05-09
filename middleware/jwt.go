package middleware

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hero1s/ginweb/i18n"
	"github.com/hero1s/ginweb/pkg/utils"
	"net/http"
)

// JWT is jwt middleware
func JWT(c *gin.Context) {
	m, err := utils.DecodeToken(c.Request)
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
		default:
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": i18n.Unauthorized.Error(),
			"msg":  i18n.GetErrorMsg(i18n.Unauthorized, 1),
			"data": nil,
		})
		c.Abort()
		return
	}
	info, _ := json.Marshal(m)
	c.Header("user_info", string(info))
	c.Next()
}
