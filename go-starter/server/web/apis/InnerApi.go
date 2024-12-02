package apis

import (
	"github.com/gin-gonic/gin"
	"go-starter/web/services"
	"net/http"
)

type InnerApi struct {
	adminService services.AdminService
}

func (this *InnerApi) Inner(c *gin.Context) {
	// TODO
	c.JSON(http.StatusOK, gin.H{
		"errno":  0,
		"errmsg": "",
		"data":   "data",
	})
}
