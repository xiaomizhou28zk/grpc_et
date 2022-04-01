package logic

import (
	lru "github.com/hashicorp/golang-lru"
)

var UserCache *lru.Cache

func InitCache() error {
	var err error
	UserCache, err = lru.New(2000000)
	if err != nil {
		return err
	}
	return nil
}
