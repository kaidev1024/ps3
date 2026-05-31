package ps3

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

type SubjectR2Folder uint8

const (
	SubjectR2FolderPost SubjectR2Folder = iota
	SubjectR2FolderComment
	SubjectR2FolderReply
	SubjectR2FolderRating
	SubjectR2FolderBroadcast
)

var AllSubjectR2Folder = []SubjectR2Folder{
	SubjectR2FolderPost,
	SubjectR2FolderComment,
	SubjectR2FolderReply,
	SubjectR2FolderRating,
	SubjectR2FolderBroadcast,
}

func (f SubjectR2Folder) IsValid() bool {
	switch f {
	case SubjectR2FolderPost,
		SubjectR2FolderComment,
		SubjectR2FolderReply,
		SubjectR2FolderRating,
		SubjectR2FolderBroadcast:
		return true
	}
	return false
}

var subjectR2FolderToString = map[SubjectR2Folder]string{
	SubjectR2FolderPost:      "post",
	SubjectR2FolderComment:   "comment",
	SubjectR2FolderReply:     "reply",
	SubjectR2FolderRating:    "rating",
	SubjectR2FolderBroadcast: "broadcast",
}

var subjectR2FolderFromString = map[string]SubjectR2Folder{
	"post":      SubjectR2FolderPost,
	"comment":   SubjectR2FolderComment,
	"reply":     SubjectR2FolderReply,
	"rating":    SubjectR2FolderRating,
	"broadcast": SubjectR2FolderBroadcast,
}

func (f SubjectR2Folder) String() string {
	if value, ok := subjectR2FolderToString[f]; ok {
		return value
	}
	return strconv.Itoa(int(f))
}

func (f *SubjectR2Folder) UnmarshalGQL(v any) error {
	switch val := v.(type) {
	case string:
		if parsed, ok := subjectR2FolderFromString[val]; ok {
			*f = parsed
			return nil
		}
		return fmt.Errorf("%q is not a valid SubjectR2Folder", val)
	case json.Number:
		num, err := val.Int64()
		if err != nil {
			return fmt.Errorf("invalid SubjectR2Folder number: %w", err)
		}
		*f = SubjectR2Folder(num)
		if !f.IsValid() {
			return fmt.Errorf("%d is not a valid SubjectR2Folder", num)
		}
		return nil
	default:
		return fmt.Errorf("SubjectR2Folder must be a string or number")
	}
}

func (f SubjectR2Folder) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, "%q", f.String())
}
