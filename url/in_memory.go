package url

import (
	"fmt"
	"math/rand"
	"time"
)

type inMemory struct {
	db map[string]*Url
}

func (i *inMemory) Create(url string) (*Url, error) {
	newUrl := Url{
		ShortID:     randStringRunes(6),
		OriginalUrl: url,
		CreatedAt:   time.Now(),
	}
	i.db[newUrl.ShortID] = &newUrl
	return &newUrl, nil
}

func (i *inMemory) FindByShortID(shortID string) (*Url, error) {
	val, ok := i.db[shortID]
	if !ok {
		return nil, fmt.Errorf("couldn't find link")
	}
	return val, nil
}

func NewInMemoryRepository() Repository {
	return &inMemory{db: map[string]*Url{}}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
