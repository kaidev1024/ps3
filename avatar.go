package ps3

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

const avatarBucketName = "avatar"

var avatarImageSizes = []ImageSize{ImageSizeSm, ImageSizeMd}

func getAvatarImageKey(folder AvatarR2Folder, pageID, imageID string, imageSize ImageSize) string {
	return fmt.Sprintf("%s/%s/%s/%s.webp", folder, pageID, imageID, imageSize)
}

func CreateAvatarUploadURL(folder AvatarR2Folder, pageID, imageID string) (string, error) {
	return presignUpload(getAvatarImageKey(folder, pageID, imageID, ImageSizeLg), avatarBucketName)
}

func DownloadAvatarImage(ctx context.Context, folder AvatarR2Folder, pageID, imageID string) ([]byte, string, error) {
	return downloadImageFromR2(ctx, getAvatarImageKey(folder, pageID, imageID, ImageSizeLg), avatarBucketName)
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
			key := getAvatarImageKey(folder, pageID, imageID, size)
			errs[i] = uploadImageToR2(ctx, key, avatarBucketName, images[i])
		}(i, size)
	}
	wg.Wait()
	return errors.Join(errs...)
}
