package service

import (
	"go-crud/model"
	"go-crud/serializer"
)

// UpdateVideoService 更新视频的服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}

// Update 更新视频
func (service *UpdateVideoService) Update(id string) serializer.Response {
	video := model.Video{}
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "未找到对应内容",
			Error:  err.Error(),
		}
	}
	video.Title = service.Title
	video.Info = service.Info
	err = model.DB.Save(&video).Error
	if err != nil {
		return serializer.Response{
			Status: 50002,
			Msg:    "视频更新失败",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
