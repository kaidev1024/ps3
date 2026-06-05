package ps3

import (
	"context"
	"errors"
	"fmt"
	"sync"
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

func UploadPageImages(ctx context.Context, folder PageR2Folder, pageID, imageID string, images [][]byte) error {
	if len(images) != 3 {
		return fmt.Errorf("expected 3 images (sm, md, lg), got %d", len(images))
	}
	errs := make([]error, len(imageSizes))
	var wg sync.WaitGroup
	for i, size := range imageSizes {
		wg.Add(1)
		go func(i int, size ImageSize) {
			defer wg.Done()
			key := getPageImageKey(folder, pageID, appendImageSize(imageID, size))
			errs[i] = uploadImageToR2(ctx, key, pageBucketName, images[i])
		}(i, size)
	}
	wg.Wait()
	return errors.Join(errs...)
}
