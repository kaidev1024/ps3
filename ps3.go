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
var bucketName string

const cdnBaseURL = "image.buddiesnearby.com"

func Init(accessKey, secretKey, accountID, _bucketName string) {
	endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID)
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("auto"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
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

// GetCDNURL returns the public CDN URL for a given key.
func GetCDNURL(key string) string {
	return fmt.Sprintf("%s/%s", cdnBaseURL, key)
}

func GetPresignedUploadURL(key string, expiry time.Duration) (string, error) {
	presigner := s3.NewPresignClient(c)
	result, err := presigner.PresignPutObject(context.Background(),
		&s3.PutObjectInput{
			Bucket: &bucketName,
			Key:    &key,
		},
		s3.WithPresignExpires(expiry),
	)
	if err != nil {
		return "", err
	}
	return result.URL, nil
}
