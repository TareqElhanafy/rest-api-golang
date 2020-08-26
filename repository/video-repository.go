package repository

import (
	"RestApi/entity"

	//sql driver wuth gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//VideoRepository interface
type VideoRepository interface {
	Store(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
}

func (db *database) Store(video entity.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Find(&videos)
	return videos
}
