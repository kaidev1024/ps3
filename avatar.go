package ps3

import (
	"fmt"
)

const avatarBucketName = "avatar"

type AvatarUploadInput struct {
	ImageID     string
	PageID      string
	ContentType string
}

func CreateAvatarUploadURL(pageID, imageID, contentType string) (string, error) {
	return presignUpload(getAvatarKey(pageID, imageID), contentType, avatarBucketName)
}

func getAvatarKey(pageID, imageID string) string {
	return fmt.Sprintf("%s/%s", pageID, imageID)
}
