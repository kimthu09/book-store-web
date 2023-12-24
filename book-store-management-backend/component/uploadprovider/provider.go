package uploadprovider

import "book-store-management-backend/common"

type UploadStaticProvider interface {
	UploadImage(data []byte, path string) (common.Image, error)
}
