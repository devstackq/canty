package storage

import (
	"io/ioutil"

	"canty/internal/core/entities"
)

type FileStorage struct {
	StoragePath string
}

func NewFileStorage(storagePath string) *FileStorage {
	return &FileStorage{StoragePath: storagePath}
}

func (fs *FileStorage) SaveVideo(video *entities.Video) error {
	filePath := fs.StoragePath + "/" + video.ID + "." + video.Format
	err := ioutil.WriteFile(filePath, video.Content, 0644)
	if err != nil {
		return err
	}
	video.FilePath = filePath
	return nil
}
