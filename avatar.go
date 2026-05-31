package ps3

import "fmt"

const avatarBucketName = "avatar"

func CreateAvatarUploadURL(folder AvatarR2Folder, pageID, imageID, contentType string) (string, error) {
	imageKey := fmt.Sprintf("%s/%s/%s", folder, pageID, imageID)
	return presignUpload(imageKey, contentType, avatarBucketName)
}
