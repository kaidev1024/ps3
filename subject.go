package ps3

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

const subjectBucketName = "subject"

type SubjectImageUploadInput struct {
	ImageID      string
	CreationHour int32
}

func getSubjectImageKey(folder SubjectR2Folder, creationHour int32, imageID string, imageSize ImageSize) string {
	return fmt.Sprintf("%s/%d/%s/%s.webp", folder, creationHour, imageID, imageSize)
}

func CreateSubjectImageUploadURLs(folder SubjectR2Folder, inputs []SubjectImageUploadInput) ([]string, error) {
	urls := make([]string, len(inputs))
	for i, input := range inputs {
		url, err := presignUpload(getSubjectImageKey(folder, input.CreationHour, input.ImageID, ImageSizeLg), subjectBucketName)
		if err != nil {
			return nil, err
		}
		urls[i] = url
	}
	return urls, nil
}

func DownloadSubjectImage(ctx context.Context, folder SubjectR2Folder, creationHour int32, imageID string) ([]byte, string, error) {
	return downloadImageFromR2(ctx, getSubjectImageKey(folder, creationHour, imageID, ImageSizeLg), subjectBucketName)
}

func UploadSubjectImages(ctx context.Context, folder SubjectR2Folder, creationHour int32, imageID string, images [][]byte) error {
	if len(images) != IMAGES_LEN {
		return fmt.Errorf("expected 2 images (sm, md), got %d", len(images))
	}
	errs := make([]error, len(imageSizes))
	var wg sync.WaitGroup
	for i, size := range imageSizes {
		wg.Add(1)
		go func(i int, size ImageSize) {
			defer wg.Done()
			key := getSubjectImageKey(folder, creationHour, imageID, size)
			errs[i] = uploadImageToR2(ctx, key, subjectBucketName, images[i])
		}(i, size)
	}
	wg.Wait()
	return errors.Join(errs...)
}
