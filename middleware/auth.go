package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("session_token")
		if err != nil {
			if strings.Contains(ctx.GetHeader("Content-Type"), "application/json") {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
			} else {
				ctx.Redirect(303, "/login")
			}
			return
		}
		tokenString := cookie.Value
		claims:= &model.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(model.JwtKey), nil
		})
		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(400, gin.H{"error": "Unauthorized"})
		}
		ctx.Set("email", claims.Email)
		ctx.Next()
		// TODO: answer here
	})
}
