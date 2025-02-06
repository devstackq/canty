package entities

type VideoRepository interface {
	Create(video *Video) error
	Read(id string) (*Video, error)
	Update(video *Video) error
	Delete(id string) error
}
type Video struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	URL         string   `json:"url"`
	Content     []byte   `json:"content"`    // Контент видео как []byte
	FilePath    string   `json:"file_path"`  // Путь к локальному видеофайлу
	Format      string   `json:"format"`     // Формат видео (например, MP4, WebM)
	Duration    int      `json:"duration"`   // Длительность видео в секундах
	Resolution  string   `json:"resolution"` // Разрешение видео (например, 1920x1080)
	Size        int64    `json:"size"`       // Размер видеофайла в байтах
	Tags        []string `json:"tags"`       // Теги, связанные с видео
}
