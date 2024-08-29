package url

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type UrlDbEntity struct {
	gorm.Model
	ShortID     string
	OriginalUrl string
	CreatedAt   time.Time
}

func (u *UrlDbEntity) TableName() string {
	return "urls"
}

type sqliteRepo struct {
	db *gorm.DB
}

func NewSqliteRepo(db *gorm.DB) Repository {
	return &sqliteRepo{db: db}
}

func (r *sqliteRepo) Create(url string) (*Url, error) {
	newUrl := UrlDbEntity{ShortID: randStringRunes(6), OriginalUrl: url, CreatedAt: time.Now()}
	result := r.db.Create(&newUrl)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Url{
		ShortID:     newUrl.ShortID,
		OriginalUrl: newUrl.OriginalUrl,
		CreatedAt:   newUrl.CreatedAt,
	}, nil
}

func (r *sqliteRepo) FindByShortID(shortID string) (*Url, error) {
	var url UrlDbEntity
	r.db.Where("short_id = ?", shortID).First(&url)
	if url.ShortID == "" {
		return nil, fmt.Errorf("couldn't find link")
	}
	return &Url{
		ShortID:     url.ShortID,
		OriginalUrl: url.OriginalUrl,
		CreatedAt:   url.CreatedAt,
	}, nil
}
