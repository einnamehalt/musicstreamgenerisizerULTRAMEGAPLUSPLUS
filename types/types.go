package types

type SongInformation struct {
	Title, Artist, Album string
}

type StreamingProvider interface {
	IsThisProvider(string) bool
	GetSongInformation(string) SongInformation
	GetStreamURL(SongInformation) string
}
