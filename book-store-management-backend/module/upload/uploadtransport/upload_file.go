package uploadtransport

import (
	"book-store-management-backend/common"
	"book-store-management-backend/component/appctx"
	"book-store-management-backend/component/generator"
	"book-store-management-backend/component/uploadprovider"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UploadFile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		gen := generator.NewShortIdGenerator()
		upP := uploadprovider.NewFirebaseStorageUploadProvider("uit-bookstore-app.appspot.com", "/data/private/firebase-auth-key.json")
		picId, err := gen.GenerateId()
		if err != nil {
			panic(err)
		}
		res, err := upP.UploadImage(c.Request.Context(), dataBytes, fmt.Sprintf("%s/%s-%s", folder, picId, fileHeader.Filename))
		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(res))

	}
}
