package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从Header获取Authorization: Bearer <token>
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "缺少或无效的Token"})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 从配置读取secret
		secret := []byte(viper.GetString("jwt.secret"))
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 校验签名算法
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return secret, nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Token无效或已过期"})
			return
		}
		// 校验通过，继续后续处理
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := claims["user_id"]
			username := claims["username"]
			c.Set("user_id", userID)
			c.Set("username", username)
		}
		c.Next()
	}
}
