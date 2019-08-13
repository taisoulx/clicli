package api

import (
	"clicli/service"

	"github.com/gin-gonic/gin"
)

// CreateVideo 视频投稿
func CreateVideo(c *gin.Context) {
	var service service.CreateVideoService

	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(201, ErrorResponse(err))
	}
}

//ShowVideo 视频详情接口
func ShowVideo(c *gin.Context) {
	var service service.ShowVideoService
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// ListVideo 视频列表接口
func ListVideo(c *gin.Context) {
	var service service.ListVideoService

	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(201, ErrorResponse(err))
	}
}

// UpdateVideo 视频更新接口
func UpdateVideo(c *gin.Context) {
	var service service.UpdateVideoService

	if err := c.ShouldBind(&service); err == nil {
		res := service.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		c.JSON(201, ErrorResponse(err))
	}
}

// DeleteVideo 视频删除接口
func DeleteVideo(c *gin.Context) {
	var service service.DeleteVideoService
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}
