package uploadfilebiz

import "mime/multipart"

type uploadFileBiz struct {
	filePath string
	file     *multipart.FileHeader
}

func NewUploadFileBiz(file *multipart.FileHeader, filePath string) *uploadFileBiz {
	return &uploadFileBiz{
		file:     file,
		filePath: filePath,
	}
}
