package common

import (
	"book-store-management-backend/component/appctx"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Id        int    `json:"id" gorm:"column:id;"`
	Url       string `json:"url" gorm:"column:url;"`
	CloudName string `json:"cloudName,omitempty" gorm:"-"`
}

func (Image) TableName() string { return TableImage }

func (j *Image) Fulfill(domain string) {
	j.Url = fmt.Sprintf("%s/%s", domain, j.Url)
}

func (j *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal data from DB:", value))
	}

	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*j = img
	return nil
}

// Value return json value, implement driver.Valuer interface
func (j *Image) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

func GetImageFromURL(appCtx appctx.AppContext, url string) Image {
	img := Image{}

	if url[:len(appCtx.GetServerHost())] == appCtx.GetServerHost() {
		img.Url = url[len(appCtx.GetServerHost()):]
		img.CloudName = "local"
	} else {
		img.Url = url
		img.CloudName = "cloudinary"
	}

	return img
}
