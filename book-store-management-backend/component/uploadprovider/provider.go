package uploadprovider

import (
	"book-store-management-backend/common"
	"context"
)

type UploadProvider interface {
	UploadImage(ctx context.Context, data []byte, dst string) (*common.Image, error)
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
