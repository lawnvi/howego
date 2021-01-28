package part_x

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routers(e *gin.Engine)  {
	v1 := e.Group("/v1")
	{
		v1.GET("/test", test)
	}
}

func test(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"message": "hello world"})
}