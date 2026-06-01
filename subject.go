package ps3

import (
	"context"
	"fmt"
)

const subjectBucketName = "subject"

type SubjectImageUploadInput struct {
	ImageID      string
	ContentType  string
	CreationHour int32
}

func getSubjectImageKey(folder SubjectR2Folder, creationHour int32, imageID string) string {
	return fmt.Sprintf("%s/%d/%s", folder, creationHour, imageID)
}

func CreateSubjectImageUploadURLs(folder SubjectR2Folder, inputs []SubjectImageUploadInput) ([]string, error) {
	urls := make([]string, len(inputs))
	for i, input := range inputs {
		url, err := presignUpload(getSubjectImageKey(folder, input.CreationHour, input.ImageID), input.ContentType, subjectBucketName)
		if err != nil {
			return nil, err
		}
		urls[i] = url
	}
	return urls, nil
}

func DownloadSubjectImage(ctx context.Context, folder SubjectR2Folder, creationHour int32, imageID string) ([]byte, string, error) {
	return downloadImageFromR2(ctx, getSubjectImageKey(folder, creationHour, imageID), subjectBucketName)
}

func UploadSubjectImage(ctx context.Context, folder SubjectR2Folder, creationHour int32, imageID string, data []byte) error {
	return uploadImageToR2(ctx, getSubjectImageKey(folder, creationHour, imageID), subjectBucketName, data)
}
