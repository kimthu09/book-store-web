package uploadprovider

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
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

func (provider *staticUploadProvider) UploadImage(data []byte, folderName string, filename string) (common.Image, error) {
	image := common.Image{
		CloudName: "local",
		Url:       "/" + folderName + "/" + filename,
	}

	// create file
	fullPath := filepath.Join(provider.staticPath, folderName, filename)
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

func (provider *staticUploadProvider) GetStaticUrl(appCtx appctx.AppContext, serverStaticPath string, image common.Image) string {
	if image.CloudName != "local" && image.CloudName != "" {
		return image.Url
	}

	url := appCtx.GetServerHost() + serverStaticPath + image.Url
	return url
}
