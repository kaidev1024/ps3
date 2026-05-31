package ps3

import (
	"fmt"
)

const subjectBucketName = "subject"

type SubjectImageUploadInput struct {
	ImageID      string
	ContentType  string
	CreationHour int32
}

func CreateSubjectImageUploadURL(folder SubjectR2Folder, creationHour int32, imageID, contentType string) (string, error) {
	return presignUpload(getSubjectImageKey(folder, creationHour, imageID), contentType, subjectBucketName)
}

func GetSubjectImageCdnUrl(folder SubjectR2Folder, creationHour int32, imageID string, size ImageSize) string {
	return fmt.Sprintf("https://%s/%s/%s/%d/%s_%s.webp", cdnBaseURL, subjectBucketName, folder, creationHour, imageID, size)
}

func CreateSubjectImageUploadURLs(folder SubjectR2Folder, inputs []SubjectImageUploadInput) ([]string, error) {
	urls := make([]string, len(inputs))
	for i, input := range inputs {
		url, err := CreateSubjectImageUploadURL(folder, input.CreationHour, input.ImageID, input.ContentType)
		if err != nil {
			return nil, err
		}
		urls[i] = url
	}
	return urls, nil
}

func getSubjectImageKey(folder SubjectR2Folder, creationHour int32, imageID string) string {
	return fmt.Sprintf("%s/%d/%s", folder, creationHour, imageID)
}
