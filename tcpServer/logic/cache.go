package logic

import (
	lru "github.com/hashicorp/golang-lru"
)

var UserCache *lru.Cache

func InitCache() error {
	var err error
	UserCache, err = lru.New(200)
	if err != nil {
		return err
	}
	return nil
}
