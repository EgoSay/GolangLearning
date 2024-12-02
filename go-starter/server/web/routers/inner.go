package routers

import (
	"github.com/gin-gonic/gin"
	"go-starter/web/apis"
)

type inner struct{}

func NewInner() *inner {
	return &inner{}
}

func (this *inner) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

func (this *inner) Register(e *gin.Engine) {
	//e.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, "pong")
	//})

	pubV1 := e.Group("/inner/v1")

	//用户权限校验
	pubV1.Use(this.authMiddleware())

	//挖矿相关
	innerApi := new(apis.InnerApi)
	s := pubV1.Group("/inner")
	{
		s.GET("/xxx", innerApi.Inner)
		s.POST("/xxxx", innerApi.Inner)
	}

}
