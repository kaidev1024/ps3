package ps3

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

const pageBucketName = "page"

type PageImageUploadInput struct {
	ImageID string
	PageID  string
}

func getPageImageKey(folder PageR2Folder, pageID, imageID string, imageSize ImageSize) string {
	return fmt.Sprintf("%s/%s/%s/%s.webp", folder, pageID, imageID, imageSize)
}

func CreatePageImageUploadURLs(folder PageR2Folder, inputs []PageImageUploadInput) ([]string, error) {
	urls := make([]string, len(inputs))
	for i, input := range inputs {
		url, err := presignUpload(getPageImageKey(folder, input.PageID, input.ImageID, ImageSizeLg), pageBucketName)
		if err != nil {
			return nil, err
		}
		urls[i] = url
	}
	return urls, nil
}

func DownloadPageImage(ctx context.Context, folder PageR2Folder, pageID, imageID string) ([]byte, string, error) {
	return downloadImageFromR2(ctx, getPageImageKey(folder, pageID, imageID, ImageSizeLg), pageBucketName)
}

func UploadPageImages(ctx context.Context, folder PageR2Folder, pageID, imageID string, images [][]byte) error {
	if len(images) != IMAGES_LEN {
		return fmt.Errorf("expected 2 images (sm, md), got %d", len(images))
	}
	errs := make([]error, len(imageSizes))
	var wg sync.WaitGroup
	for i, size := range imageSizes {
		wg.Add(1)
		go func(i int, size ImageSize) {
			defer wg.Done()
			key := getPageImageKey(folder, pageID, imageID, size)
			errs[i] = uploadImageToR2(ctx, key, pageBucketName, images[i])
		}(i, size)
	}
	wg.Wait()
	return errors.Join(errs...)
}
