package streamingProviders

import (
	"strings"

	"github.com/einnamehalt/musicstreamgenerisizerULTRAMEGAPLUSPLUS/types"
)

var spotifyApiKey string = "<API KEY>"

type spotify struct {
	hostname string
}

func (spotify spotify) IsThisProvider(url string) bool {
	return strings.Contains(url, spotify.hostname)
}

func (spotify spotify) GetSongInformation(inputUrl string) types.SongInformation {
	return types.SongInformation{Title: "", Artist: "", Album: ""}
}

func (spotfiy spotify) GetStreamURL(songInformation types.SongInformation) string {

	return "https://"
}

var Spotify = spotify{
	hostname: "spotify.com",
}
