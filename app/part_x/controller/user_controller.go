package controller

import (
	"github.com/gin-gonic/gin"
	"howego/app/part_x/service"
	"howego/model"
	"net/http"
)

type userController struct {
	us service.UserService
}

var uc userController

func UserRouters(e *gin.Engine)  {
	uc = userController{
		us: service.NewUserService(),
	}
	v1 := e.Group("/v1")
	{
		v1.POST("/sign", sign)
		v1.GET("/user/:email", userInfo)
	}
}

func sign(c *gin.Context)  {
	user := model.User{}
	if err := c.ShouldBind(&user); err != nil{
		c.JSON(http.StatusOK, "error param")
		return
	}
	c.JSON(http.StatusOK, uc.us.SignUp(user))
}

func userInfo(c *gin.Context)  {
	email := c.Param("email")
	c.JSON(http.StatusOK, uc.us.GetUserInfo(email))
}