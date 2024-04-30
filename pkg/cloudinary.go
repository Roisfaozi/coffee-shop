package pkg

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"os"
)

func Cloudinary(file interface{}) (string, error) {
	name := os.Getenv("CD_NAME")
	key := os.Getenv("CD_KEY")
	secret := os.Getenv("CD_SECRET")

	cld, _ := cloudinary.NewFromParams(name, key, secret)
	result, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder:         "Coffee_shop",
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})
	if err != nil {
		return "", err
	}
	return result.URL, nil
}
