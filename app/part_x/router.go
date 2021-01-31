package part_x

import (
	"github.com/gin-gonic/gin"
	"howego/app/part_x/service"
	"howego/model"
	"net/http"
)

func Routers(e *gin.Engine)  {
	v1 := e.Group("/v1")
	{
		v1.POST("/sign", sign)
		v1.GET("/user/:email", userInfo)
	}
}

func sign(c *gin.Context)  {
	us := service.NewUserService()
	user := model.User{}
	if err := c.ShouldBind(&user); err != nil{
		c.JSON(http.StatusOK, "error param")
		return
	}
	c.JSON(http.StatusOK, us.SignUp(user))
}

func userInfo(c *gin.Context)  {
	us := service.NewUserService()
	email := c.Param("email")
	c.JSON(http.StatusOK, us.GetUserInfo(email))
}