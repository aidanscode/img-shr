package flags

import (
	"errors"
	"flag"
)

type AppFlags struct {
	ImagesPath string
	DBPath string
}

func GetAppFlags() (*AppFlags, error) {
	imagesPath := flag.String("images", "./images", "Path to directory where images will be stored")
	dbPath := flag.String("db", "./imgshr.sqlite", "Path to SQLite db")
	flag.Parse()

	if *imagesPath == "" {
		return nil, errors.New("images path flag (optional) cannot be empty")
	}
	if *dbPath == "" {
		return nil, errors.New("db path flag (optional) cannot be empty")
	}

	return &AppFlags{ImagesPath: *imagesPath, DBPath: *dbPath}, nil
}
