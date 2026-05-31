package ps3

import (
	"fmt"
)

const avatarBucketName = "avatar"

type ImageUploadInput struct {
	ImageID     string
	PageID      string
	ContentType string
}

func CreateAvatarUploadURL(pageID, imageID, contentType string) (string, error) {
	imageKey := fmt.Sprintf("%s/%s", pageID, imageID)
	return presignUpload(imageKey, contentType, avatarBucketName)
}
