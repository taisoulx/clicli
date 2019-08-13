package service

import (
	"clicli/model"
	"clicli/serializer"
)

// ShowVideoService 视频详情的服务
type ShowVideoService struct {
}

// Show 展示视频
func (service *ShowVideoService) Show(id string) serializer.Response {
	video := model.Video{}
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "未找到对应内容",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
