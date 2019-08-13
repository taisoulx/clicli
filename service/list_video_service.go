package service

import (
	"clicli/model"
	"clicli/serializer"
)

// ShowVideoService 视频列表的服务
type ListVideoService struct {
}

// List视频列表
func (service *ListVideoService) List() serializer.Response {
	videos := []model.Video{}
	err := model.DB.Find(&videos).Error
	if err != nil {
		return serializer.Response{
			Status: 5000,
			Msg:    "数据库查询错误",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}
}
