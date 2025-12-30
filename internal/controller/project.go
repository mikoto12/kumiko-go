package controller

import (
	"kumiko/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetProject(c *gin.Context) {
	utils.Success(c, "test")
}
