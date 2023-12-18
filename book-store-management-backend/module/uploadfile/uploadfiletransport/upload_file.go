package uploadfiletransport

import (
	"book-store-management-backend/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func UploadFile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Specify the form field name
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Specify the directory where files should be saved
		staticDir := "./static"

		// Create the full path for the new file
		filename := filepath.Join(staticDir, filepath.Base(file.Filename))

		// Save the file
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "File uploaded successfully.",
			"path":    filename,
		})
	}
}
