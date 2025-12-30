package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

// Login godoc
// @Summary 用户登录
// @Description 用户登录，返回JWT Token
// @Tags 认证
// @Accept application/x-www-form-urlencoded
// @Param username formData string true "用户名"
// @Param password formData string true "密码"
// @Success 200 {object} map[string]interface{} "返回token"
// @Failure 400 {object} map[string]interface{}
// @Router /login [post]
func Login(c *gin.Context) {

	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Print(password)
	userID := 123

	secret := []byte(viper.GetString("jwt.secret"))
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "生成Token失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  tokenString,
	})
}
