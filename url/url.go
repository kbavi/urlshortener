package url

import (
	"time"
)

type Url struct {
	ShortID     string    `json:"short_id"`
	OriginalUrl string    `json:"original_url"`
	CreatedAt   time.Time `json:"created_at"`
}

type Repository interface {
	Create(string) (*Url, error)
	FindByShortID(shortID string) (*Url, error)
}

type Service interface {
	Create(string) (*Url, error)
	FindByShortID(shortID string) (*Url, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Create(url string) (*Url, error) {
	return s.repo.Create(url)
}

func (s *service) FindByShortID(shortID string) (*Url, error) {
	return s.repo.FindByShortID(shortID)
}
