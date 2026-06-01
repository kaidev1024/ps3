package ps3

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// DownloadImageFromR2 fetches an image from R2 storage.
func downloadImageFromR2(ctx context.Context, key, bucketName string) ([]byte, string, error) {
	resp, err := c.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, "", fmt.Errorf("download image from r2: %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("download image from r2: read body: %w", err)
	}

	contentType := aws.ToString(resp.ContentType)
	return data, contentType, nil
}

// UploadImageToR2 uploads image bytes directly to R2.
func uploadImageToR2(ctx context.Context, key, bucketName string, data []byte) error {
	_, err := c.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("image/webp"),
	})
	if err != nil {
		return fmt.Errorf("upload image to r2: %w", err)
	}
	return nil
}
