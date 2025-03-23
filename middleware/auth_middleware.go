package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/michaelyusak/go-helper/appconstant"
	"github.com/michaelyusak/go-helper/dto"
	"github.com/michaelyusak/kredit-plus-xyz/utils"
)

func AuthMiddleware(jwtHelper utils.JWTHelper) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get(appconstant.Authorization)
		t := strings.Split(authHeader, " ")

		if len(t) != 2 || t[0] != appconstant.Bearer {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Message: appconstant.MsgUnauthorized})
			return
		}

		authToken := t[1]

		_, err := jwtHelper.ParseAndVerify(authToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, dto.ErrorResponse{Message: appconstant.MsgUnauthorized})
			return
		}

		c.Next()
	}
}
