package db

import (
	"github.com/go-bongo/bongo"
)

func Config() *bongo.Config {
	config := &bongo.Config{
		ConnectionString: "localhost",
		Database:         "test",
	}
	return config
}
