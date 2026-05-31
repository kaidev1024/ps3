package ps3

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

type AvatarR2Folder uint8

const (
	AvatarR2FolderCollection AvatarR2Folder = iota
	AvatarR2FolderGame
	AvatarR2FolderUser
	AvatarR2FolderTeam
	AvatarR2FolderLeaderboard
	AvatarR2FolderArena
	AvatarR2FolderEvent
)

var AllAvatarR2Folder = []AvatarR2Folder{
	AvatarR2FolderCollection,
	AvatarR2FolderGame,
	AvatarR2FolderUser,
	AvatarR2FolderTeam,
	AvatarR2FolderLeaderboard,
	AvatarR2FolderArena,
	AvatarR2FolderEvent,
}

func (f AvatarR2Folder) IsValid() bool {
	switch f {
	case AvatarR2FolderCollection,
		AvatarR2FolderGame,
		AvatarR2FolderUser,
		AvatarR2FolderTeam,
		AvatarR2FolderLeaderboard,
		AvatarR2FolderArena,
		AvatarR2FolderEvent:
		return true
	}
	return false
}

var avatarR2FolderToString = map[AvatarR2Folder]string{
	AvatarR2FolderCollection:  "collection",
	AvatarR2FolderGame:        "game",
	AvatarR2FolderUser:        "user",
	AvatarR2FolderTeam:        "team",
	AvatarR2FolderLeaderboard: "leaderboard",
	AvatarR2FolderArena:       "arena",
	AvatarR2FolderEvent:       "event",
}

var avatarR2FolderFromString = map[string]AvatarR2Folder{
	"collection":  AvatarR2FolderCollection,
	"game":        AvatarR2FolderGame,
	"user":        AvatarR2FolderUser,
	"team":        AvatarR2FolderTeam,
	"leaderboard": AvatarR2FolderLeaderboard,
	"arena":       AvatarR2FolderArena,
	"event":       AvatarR2FolderEvent,
}

func (f AvatarR2Folder) String() string {
	if value, ok := avatarR2FolderToString[f]; ok {
		return value
	}
	return strconv.Itoa(int(f))
}

func (f *AvatarR2Folder) UnmarshalGQL(v any) error {
	switch val := v.(type) {
	case string:
		if parsed, ok := avatarR2FolderFromString[val]; ok {
			*f = parsed
			return nil
		}
		return fmt.Errorf("%q is not a valid AvatarR2Folder", val)
	case json.Number:
		num, err := val.Int64()
		if err != nil {
			return fmt.Errorf("invalid AvatarR2Folder number: %w", err)
		}
		*f = AvatarR2Folder(num)
		if !f.IsValid() {
			return fmt.Errorf("%d is not a valid AvatarR2Folder", num)
		}
		return nil
	default:
		return fmt.Errorf("AvatarR2Folder must be a string or number")
	}
}

func (f AvatarR2Folder) MarshalGQL(w io.Writer) {
	fmt.Fprintf(w, "%q", f.String())
}
