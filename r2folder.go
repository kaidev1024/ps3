package ps3

type R2Folder string

type ImageSize string

const (
	ImageSizeSm ImageSize = "sm"
	ImageSizeMd ImageSize = "md"
	ImageSizeLg ImageSize = "lg"
)

const (
	R2FolderCollection  R2Folder = "collection"
	R2FolderPost        R2Folder = "post"
	R2FolderGame        R2Folder = "game"
	R2FolderUoi         R2Folder = "uoi"
	R2FolderUser        R2Folder = "user"
	R2FolderTeam        R2Folder = "team"
	R2FolderLeaderboard R2Folder = "leaderboard"
	R2FolderArena       R2Folder = "arena"
	R2FolderEvent       R2Folder = "event"
	R2FolderComment     R2Folder = "comment"
)
