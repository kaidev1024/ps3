package ps3

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

type PageR2Folder uint8

const (
	PageR2FolderCollection PageR2Folder = iota
	PageR2FolderGame
	PageR2FolderUoi
	PageR2FolderUser
	PageR2FolderTeam
	PageR2FolderLeaderboard
	PageR2FolderArena
	PageR2FolderEvent
	PageR2FolderLoi
	PageR2FolderLocation
)

var AllPageR2Folder = []PageR2Folder{
	PageR2FolderCollection,
	PageR2FolderGame,
	PageR2FolderUoi,
	PageR2FolderUser,
	PageR2FolderTeam,
	PageR2FolderLeaderboard,
	PageR2FolderArena,
	PageR2FolderEvent,
	PageR2FolderLoi,
	PageR2FolderLocation,
}

func (f PageR2Folder) IsValid() bool {
	switch f {
	case PageR2FolderCollection,
		PageR2FolderGame,
		PageR2FolderUoi,
		PageR2FolderUser,
		PageR2FolderTeam,
		PageR2FolderLeaderboard,
		PageR2FolderArena,
		PageR2FolderEvent,
		PageR2FolderLoi,
		PageR2FolderLocation:
		return true
	}
	return false
}

var pageR2FolderToString = map[PageR2Folder]string{
	PageR2FolderCollection:  "collection",
	PageR2FolderGame:        "game",
	PageR2FolderUoi:         "uoi",
	PageR2FolderUser:        "user",
	PageR2FolderTeam:        "team",
	PageR2FolderLeaderboard: "leaderboard",
	PageR2FolderArena:       "arena",
	PageR2FolderEvent:       "event",
	PageR2FolderLoi:         "loi",
	PageR2FolderLocation:    "location",
}

var pageR2FolderFromString = map[string]PageR2Folder{
	"collection":  PageR2FolderCollection,
	"game":        PageR2FolderGame,
	"uoi":         PageR2FolderUoi,
	"user":        PageR2FolderUser,
	"team":        PageR2FolderTeam,
	"leaderboard": PageR2FolderLeaderboard,
	"arena":       PageR2FolderArena,
	"event":       PageR2FolderEvent,
	"loi":         PageR2FolderLoi,
	"location":    PageR2FolderLocation,
}

func (f PageR2Folder) String() string {
	if value, ok := pageR2FolderToString[f]; ok {
		return value
	}
	return strconv.Itoa(int(f))
}

func (f *PageR2Folder) UnmarshalGQL(v any) error {
	switch val := v.(type) {
	case string:
		if parsed, ok := pageR2FolderFromString[val]; ok {
			*f = parsed
			return nil
		}
		return fmt.Errorf("%q is not a valid PageR2Folder", val)
	case json.Number:
		num, err := val.Int64()
		if err != nil {
			return fmt.Errorf("invalid PageR2Folder number: %w", err)
		}
		*f = PageR2Folder(num)
		if !f.IsValid() {
			return fmt.Errorf("%d is not a valid PageR2Folder", num)
		}
		return nil
	default:
		return fmt.Errorf("PageR2Folder must be a string or number")
	}
}

func (f PageR2Folder) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, "%q", f.String())
}
