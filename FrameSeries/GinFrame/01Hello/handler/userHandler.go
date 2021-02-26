package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserSave(c *gin.Context) {
	username := c.Param("name")
	c.String(http.StatusOK, username+"已保存")
}
