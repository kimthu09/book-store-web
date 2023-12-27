package uploadprovider

import (
	"book-store-management-backend/common"
	"cloud.google.com/go/storage"
	"context"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	"log"
)

type UploadProvider interface {
	UploadImage(ctx context.Context, data []byte, dst string) (*common.Image, error)
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}

type firebaseStorageUploadProvider struct {
	storageBucketUri string
	keyFilePath      string
	firebaseApp      *firebase.App
}

func NewFirebaseStorageUploadProvider(bucketUri string, keyFilePath string) *firebaseStorageUploadProvider {
	provider := &firebaseStorageUploadProvider{
		storageBucketUri: bucketUri,
		keyFilePath:      keyFilePath,
	}

	config := &firebase.Config{
		StorageBucket: bucketUri,
	}
	opt := option.WithCredentialsFile(keyFilePath)

	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalln(err)
	}
	provider.firebaseApp = app

	return provider
}

func (provider *firebaseStorageUploadProvider) UploadImage(ctx context.Context, data []byte, dst string) (string, error) {
	app := provider.firebaseApp

	client, err := app.Storage(ctx)
	if err != nil {
		return "", err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return "", err
	}

	// Upload an object with storage.Writer.
	wc := bucket.Object(dst).NewWriter(ctx)
	if _, err = wc.Write(data); err != nil {
		return "", err
	}

	if err := wc.Close(); err != nil {
		return "", err
	}

	// Provide a public read ACL so the image can be accessed by anyone
	bucket.Object(dst).ACL().Set(ctx, storage.AllUsers, storage.RoleReader)

	attrs, err := bucket.Object(dst).Attrs(ctx)
	if err != nil {
		return "", err
	}
	return attrs.MediaLink, nil
}
