package ps3

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

var c *s3.Client
var bucketName string

const cdnBaseURL = "image.buddiesnearby.com"

var allowedContentTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/webp": true,
	"image/gif":  true,
}

func Init(accessKeyID, secretAccessKey, accountID, _bucketName string) {
	endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("auto"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
	)
	if err != nil {
		panic(err)
	}
	c = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})
	bucketName = _bucketName
	log.Println("r2 client initialized")
}

func GetKey(folder, filename string) string {
	return fmt.Sprintf("%s/%s", folder, filename)
}

func GetCDNURL(key string) string {
	return fmt.Sprintf("%s/%s", cdnBaseURL, key)
}

// GetImageUploadURL generates a presigned upload URL and the final CDN URL for an image.
// folder is the destination folder (e.g. "avatars"). contentType must be an allowed image type.
func GetImageUploadURL(folder, contentType string) (uploadURL, cdnURL string, err error) {
	if !allowedContentTypes[contentType] {
		return "", "", fmt.Errorf("unsupported content type: %s", contentType)
	}
	key := GetKey(folder, uuid.New().String())
	uploadURL, err = presignUpload(key, contentType, 15*time.Minute)
	if err != nil {
		return "", "", err
	}
	return uploadURL, GetCDNURL(key), nil
}

func presignUpload(key, contentType string, expiry time.Duration) (string, error) {
	presigner := s3.NewPresignClient(c)
	result, err := presigner.PresignPutObject(context.Background(),
		&s3.PutObjectInput{
			Bucket:      &bucketName,
			Key:         &key,
			ContentType: &contentType,
		},
		s3.WithPresignExpires(expiry),
	)
	if err != nil {
		return "", err
	}
	return result.URL, nil
}
