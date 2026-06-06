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

func Init(accessKeyID, secretAccessKey, accountID string) {
	endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountID)
	cfg, err := config.LoadDefaultConfig(context.Background(),
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

func presignUpload(key, bucketName string) (string, error) {
	presigner := s3.NewPresignClient(c)
	result, err := presigner.PresignPutObject(context.Background(),
		&s3.PutObjectInput{
			Bucket:      &bucketName,
			Key:         &key,
			ContentType: aws.String("image/webp"),
		},
		s3.WithPresignExpires(15*time.Minute),
	)
	if err != nil {
		return "", err
	}
	return result.URL, nil
}
