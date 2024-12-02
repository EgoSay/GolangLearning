package routers

import (
	"github.com/gin-gonic/gin"
	"go-starter/web/apis"
)

type admin struct{}

func NewAdmin() *admin {
	return &admin{}
}

func (this *admin) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}

func (this *admin) Register(r *gin.Engine) {
	r.Use(this.authMiddleware())
	pubV1 := r.Group("/admin/v1")

	//存币2.0
	adminCtrl := &apis.AdminApi{}
	test1Pub := pubV1.Group("/test1")
	test2Pub := pubV1.Group("/test2")
	{
		test1Pub.POST("/xxxx", adminCtrl.Admin)
		test2Pub.GET("/xxxx", adminCtrl.Admin)
	}

	return
}

func (this *admin) before() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
