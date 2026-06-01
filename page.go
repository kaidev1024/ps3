package ps3

import (
	"context"
	"fmt"
)

const pageBucketName = "page"

type PageImageUploadInput struct {
	ImageID     string
	PageID      string
	ContentType string
}

func getPageImageKey(folder PageR2Folder, pageID, imageID string) string {
	return fmt.Sprintf("%s/%s/%s", folder, pageID, imageID)
}

func CreatePageImageUploadURLs(folder PageR2Folder, inputs []PageImageUploadInput) ([]string, error) {
	urls := make([]string, len(inputs))
	for i, input := range inputs {
		url, err := presignUpload(getPageImageKey(folder, input.PageID, input.ImageID), input.ContentType, pageBucketName)
		if err != nil {
			return nil, err
		}
		urls[i] = url
	}
	return urls, nil
}

func DownloadPageImage(ctx context.Context, folder PageR2Folder, pageID, imageID string) ([]byte, string, error) {
	return downloadImageFromR2(ctx, getPageImageKey(folder, pageID, imageID), pageBucketName)
}

func UploadPageImage(ctx context.Context, folder PageR2Folder, pageID, imageID string, data []byte) error {
	return uploadImageToR2(ctx, getPageImageKey(folder, pageID, imageID), pageBucketName, data)
}
