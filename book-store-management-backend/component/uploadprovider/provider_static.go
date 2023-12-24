package uploadprovider

import (
	"book-store-management-backend/common"
	"os"
	"path/filepath"
)

type staticUploadProvider struct {
	staticPath string
}

func NewStaticUploadProvider(staticPath string) *staticUploadProvider {
	return &staticUploadProvider{
		staticPath: staticPath,
	}
}

func (provider *staticUploadProvider) UploadImage(data []byte, filename string) (common.Image, error) {
	image := common.Image{
		CloudName: "local",
		Url:       "/" + filename,
	}

	// create file
	fullPath := filepath.Join(provider.staticPath, filename)
	file, err := os.Create(fullPath)
	if err != nil {
		return common.Image{}, err // Return the error if file creation fails
	}
	defer file.Close()

	// write data to file
	if _, err := file.Write(data); err != nil {
		return common.Image{}, err
	}

	return image, nil
}
