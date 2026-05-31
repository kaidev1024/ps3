package ps3

const avatarBucketName = "avatar"

func CreateAvatarUploadURL(pageID, imageID, contentType string) (string, error) {
	return presignUpload(getImageKey(pageID, imageID), contentType, avatarBucketName)
}
