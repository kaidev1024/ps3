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
)

var c *s3.Client

var bucketName = "image"

const cdnBaseURL = "image.buddiesnearby.com"

func Init(accessKeyID, secretAccessKey, accountID string) {
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
	log.Println("r2 client initialized")
}

func GetCdnUrl(folderName R2Folder, creationHour int32, imageID string, size ImageSize) string {
	return fmt.Sprintf("https://%s/%s/%s/%d/%s_%s.webp", cdnBaseURL, bucketName, folderName, creationHour, imageID, size)
}

func GetKey(folderName R2Folder, creationHour int32, imageID string) string {
	return fmt.Sprintf("%s/%d/%s", folderName, creationHour, imageID)
}

type ImageUploadInput struct {
	ImageID      string
	ContentType  string
	CreationHour int32
}

func CreateImageUploadURL(folderName R2Folder, creationHour int32, imageID, contentType string) (string, error) {
	return presignUpload(GetKey(folderName, creationHour, imageID), contentType, 15*time.Minute)
}

func CreateImageUploadURLs(folderName R2Folder, inputs []ImageUploadInput) ([]string, error) {
	urls := make([]string, len(inputs))
	for i, input := range inputs {
		url, err := CreateImageUploadURL(folderName, input.CreationHour, input.ImageID, input.ContentType)
		if err != nil {
			return nil, err
		}
		urls[i] = url
	}
	return urls, nil
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
