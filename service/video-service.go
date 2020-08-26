package service

import (
	"RestApi/entity"
	"RestApi/repository"
)

// VideoService interface descripes the save,find all functions
type VideoService interface {
	Store(entity.Video) entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
}

type videoService struct {
	repository repository.VideoRepository
}

// NewVideoService insatance of the VideoService interface
func NewVideoService(repo repository.VideoRepository) VideoService {
	return &videoService{
		repository: repo,
	}
}

func (service *videoService) Store(video entity.Video) entity.Video {
	service.repository.Store(video)
	return video
}
func (service *videoService) Update(video entity.Video) {
	service.repository.Update(video)
}
func (service *videoService) Delete(video entity.Video) {
	service.repository.Delete(video)
}

func (service *videoService) FindAll() []entity.Video {
	return service.repository.FindAll()
}
