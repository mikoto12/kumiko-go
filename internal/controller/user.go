package controller

import (
	"kumiko/internal/model"
	"kumiko/internal/service"
	"kumiko/internal/utils"
	"kumiko/pkg/elasticsearch"
	"kumiko/pkg/logger"
	"kumiko/pkg/rabbitmq"

	"github.com/gin-gonic/gin"
)

// GetUserList godoc
// @Summary 获取用户列表
// @Description 获取所有用户信息
// @Tags 用户
// @Success 200 {array} model.User
// @Failure 400 {object} map[string]interface{}
// @Router /user [get]
func GetUserList(c *gin.Context) {
	users, err := service.GetUserList()
	if err != nil {
		utils.Fail(c, "User not found")
		return
	}
	utils.Success(c, users)
}

// GetUser godoc
// @Summary 获取用户信息
// @Description 通过ID获取用户
// @Tags 用户
// @Param id path int true "用户ID"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]interface{}
// @Router /user/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := service.GetUserByID(id)
	if err != nil {
		utils.Fail(c, "User not found")
		return
	}
	utils.Success(c, user)
}

// CreateUser godoc
// @Summary 创建用户
// @Description 创建新用户
// @Tags 用户
// @Accept application/x-www-form-urlencoded
// @Param name formData string true "用户名"
// @Param email formData string true "邮箱"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]interface{}
// @Router /user [post]
func CreateUser(c *gin.Context) {
	user := model.User{
		Name:  c.PostForm("name"),
		Email: c.PostForm("email"),
	}
	service.CreateUser(&user)
	utils.Success(c, user)
}

// Test godoc
// @Summary 测试接口
// @Description 测试用
// @Tags 测试
// @Success 200 {object} map[string]interface{}
// @Router /test [get]
func Test(c *gin.Context) {
	userId, exists := c.Get("user_id")
	if exists {
		utils.Success(c, userId)
	}
	err := rabbitmq.Publish("test_queue", []byte("hello rabbitmq"))
	if err != nil {
		// 处理错误
	}
	// 新增
	err1 := elasticsearch.IndexDoc("user", "1", map[string]interface{}{"name": "Alice", "age": 18})
	logger.StdError("新增报错：", err1)
	// 删除
	err2 := elasticsearch.DeleteDoc("user", "1")
	logger.StdError("删除报错：", err2)
	// 查询
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": "Alice",
			},
		},
	}
	results, err3 := elasticsearch.SearchDoc("user", query)
	logger.StdError("查询数据：", results)
	logger.StdError("查询报错：", err3)

	utils.Success(c, "test")
}
