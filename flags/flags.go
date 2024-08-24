package flags

import (
	"errors"
	"flag"
)

type AppFlags struct {
	ImagesPath string
	DBPath string
	MaxUploadSizeBytes int64
}

func GetAppFlags() (*AppFlags, error) {
	imagesPath := flag.String("images", "./images", "Path to directory where images will be stored")
	dbPath := flag.String("db", "./imgshr.sqlite", "Path to SQLite db")
	maxUploadSizeBytes := flag.Int64("uploadsize", 10000000, "Max upload size in bytes")
	flag.Parse()

	if *imagesPath == "" {
		return nil, errors.New("images path flag (optional) cannot be empty")
	}
	if *dbPath == "" {
		return nil, errors.New("db path flag (optional) cannot be empty")
	}
	if *maxUploadSizeBytes == 0 {
		return nil, errors.New("max file upload size (optional) cannot be empty")
	}

	return &AppFlags{ImagesPath: *imagesPath, DBPath: *dbPath, MaxUploadSizeBytes: *maxUploadSizeBytes}, nil
}
