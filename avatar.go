package ps3

import (
	"context"
	"fmt"
)

const avatarBucketName = "avatar"

func getAvatarImageKey(folder AvatarR2Folder, pageID, imageID string) string {
	return fmt.Sprintf("%s/%s/%s", folder, pageID, imageID)
}

func CreateAvatarUploadURL(folder AvatarR2Folder, pageID, imageID, contentType string) (string, error) {
	return presignUpload(getAvatarImageKey(folder, pageID, imageID), contentType, avatarBucketName)
}

func DownloadAvatarImage(ctx context.Context, folder AvatarR2Folder, pageID, imageID string) ([]byte, string, error) {
	return downloadImageFromR2(ctx, getAvatarImageKey(folder, pageID, imageID), avatarBucketName)
}

func UploadAvatarImage(ctx context.Context, folder AvatarR2Folder, pageID, imageID string, data []byte) error {
	return uploadImageToR2(ctx, getAvatarImageKey(folder, pageID, imageID), avatarBucketName, data)
}
