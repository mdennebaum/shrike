package cache

import (
	"time"
)

type Item struct {
	Value   interface{}
	Expires *time.Time
}

func NewItem(value interface{}, expiration *time.Time) {
	i := Item{value, expiration}
	return &i
}

func (i *Item) IsExpired() bool {
	return (time.Now() >= i.Expires)
}
