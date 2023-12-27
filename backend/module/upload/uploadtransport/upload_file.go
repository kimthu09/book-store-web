package uploadtransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/component/uploadprovider"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"path/filepath"
)

type ResUploadFile struct {
	Data string `json:"data"`
}

// UploadFile
// @BasePath /v1
// @Security BearerAuth
// @Summary Upload file
// @Tags upload
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Upload file"
// @Param folderName formData string false "Folder name (default: images)"
// @Response 200 {object} ResUploadFile "url"
// @Router /upload [post]
func UploadFile(appCtx appctx.AppContext, serverStaticPath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		toLocal := c.DefaultPostForm("toLocal", "true")
		if toLocal != "true" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Not support upload to cloud yet",
			})
			return
		}

		folderName := c.DefaultPostForm("folderName", "images")

		if (folderName != "images") && (folderName != "avatars") {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Folder name must be images or avatars",
			})
			return
		}
		gen := generator.NewShortIdGenerator()
		id, err := gen.GenerateId()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		fileName := id + filepath.Ext(fileHeader.Filename)

		// Open the uploaded file
		file, err := fileHeader.Open()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		staticProvider := uploadprovider.NewStaticUploadProvider(appCtx.GetStaticPath())
		res, err := staticProvider.UploadImage(data, folderName, fileName)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		url := res.Url

		if res.CloudName == "local" || res.CloudName == "" || res.Url == "" {
			url = staticProvider.GetStaticUrl(appCtx, serverStaticPath, res)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(url))
	}
}
