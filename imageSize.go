package ps3

import "fmt"

type ImageSize string

const (
	ImageSizeSm ImageSize = "sm"
	ImageSizeMd ImageSize = "md"
	ImageSizeLg ImageSize = "lg"
)

var imageSizes []ImageSize = []ImageSize{
	ImageSizeSm,
	ImageSizeMd,
	ImageSizeLg,
}

func appendImageSize(imageID string, imageSize ImageSize) string {
	return fmt.Sprintf("%s/%s", imageID, imageSize)
}
