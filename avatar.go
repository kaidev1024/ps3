package ps3

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

const avatarBucketName = "avatar"

var avatarImageSizes = []ImageSize{ImageSizeSm, ImageSizeMd}

func getAvatarImageKey(folder AvatarR2Folder, pageID, imageID string) string {
	return fmt.Sprintf("%s/%s/%s", folder, pageID, imageID)
}

func CreateAvatarUploadURL(folder AvatarR2Folder, pageID, imageID, contentType string) (string, error) {
	return presignUpload(getAvatarImageKey(folder, pageID, imageID), contentType, avatarBucketName)
}

func DownloadAvatarImage(ctx context.Context, folder AvatarR2Folder, pageID, imageID string) ([]byte, string, error) {
	return downloadImageFromR2(ctx, getAvatarImageKey(folder, pageID, imageID), avatarBucketName)
}

func UploadAvatarImages(ctx context.Context, folder AvatarR2Folder, pageID, imageID string, images [][]byte) error {
	if len(images) != 2 {
		return fmt.Errorf("expected 2 images (sm, md), got %d", len(images))
	}
	errs := make([]error, len(avatarImageSizes))
	var wg sync.WaitGroup
	for i, size := range avatarImageSizes {
		wg.Add(1)
		go func(i int, size ImageSize) {
			defer wg.Done()
			key := getAvatarImageKey(folder, pageID, appendImageSize(imageID, size))
			errs[i] = uploadImageToR2(ctx, key, avatarBucketName, images[i])
		}(i, size)
	}
	wg.Wait()
	return errors.Join(errs...)
}
