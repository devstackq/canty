package entities

type Advertisement struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	URL     string `json:"url"`
}

type AdvertisementRepository interface {
	Create(ad *Advertisement) error
	Read(id string) (*Advertisement, error)
	Update(ad *Advertisement) error
	Delete(id string) error
}
