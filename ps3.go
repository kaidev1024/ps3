package ps3

import (
	"context"
	"fmt"
	"io"
	"log"

	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type (
	Bucket string
	File   struct {
		Bucket Bucket
		Folder string
		Name   string
		Data   []byte
	}
	client struct {
		client *s3.Client
		bucket string
	}
)

func (b Bucket) String() string {
	return string(b)
}

var clients = make(map[Bucket]*client)

func Init(accessKey, secretKey string, buckets []Bucket) {
	for _, bucket := range buckets {
		client, err := newClient(bucket, accessKey, secretKey)
		if err != nil {
			panic(err)
		}
		clients[bucket] = client
	}
	log.Println("s3Clients initialized")
}

func newClient(bucket Bucket, accessKey, secretKey string) (*client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-west-2"))
	if err != nil {
		return nil, err
	}
	cfg.Credentials = aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""))

	return &client{
		client: s3.NewFromConfig(cfg),
		bucket: string(bucket),
	}, nil
}

func (sc *client) UploadFile(file io.Reader, key string) error {
	_, err := sc.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Body:   file,
		Bucket: &sc.bucket,
		Key:    &key,
	})
	return err
}

func (sc *client) GetPresignedUrl(key string) (string, error) {
	presignClient := s3.NewPresignClient(sc.client)
	presignedUrl, err := presignClient.PresignGetObject(context.Background(),
		&s3.GetObjectInput{
			Bucket: &sc.bucket,
			Key:    &key,
		},
		s3.WithPresignExpires(time.Hour)) // TODO: need to adjust the expiration time

	if err != nil {
		return "", err
	}
	return presignedUrl.URL, nil
}

func GetS3Key(folder, filename string) string {
	return fmt.Sprintf("%s/%s", folder, filename)
}

func UploadFile(bucket Bucket, folder, filename string, file io.Reader) (string, error) {
	client := clients[bucket]
	key := GetS3Key(folder, filename)
	if err := client.UploadFile(file, key); err != nil {
		return "", err
	}
	return key, nil
}

func Upload(bucket Bucket, key *string, file io.Reader) error {
	client := clients[bucket]
	_, err := client.client.PutObject(context.TODO(), &s3.PutObjectInput{
		Body:   file,
		Bucket: &client.bucket,
		Key:    key,
	})
	return err
}

func GetPresignedUrl(bucket Bucket, key string) (string, error) {
	return clients[bucket].GetPresignedUrl(key)
}
