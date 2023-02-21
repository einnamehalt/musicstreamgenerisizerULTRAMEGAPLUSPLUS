package main

import (
	"fmt"

	"github.com/einnamehalt/musicstreamgenerisizerULTRAMEGAPLUSPLUS/streamingProviders"
	"github.com/einnamehalt/musicstreamgenerisizerULTRAMEGAPLUSPLUS/types"
)

var providers []types.StreamingProvider = []types.StreamingProvider{
	streamingProviders.YoutubeMusic,
	streamingProviders.Spotify,
}

func main() {
	var url string = "https://music.youtube.com/watch?v=VNWO5N8Vxsc&list=RDAMVMqdtLCfEcPL4"

	currentProvider := findCurrentProvider(providers, url)

	var songInformation types.SongInformation = currentProvider.GetSongInformation(url)

	streamURLs := getAllStreamURLs(songInformation)

	fmt.Println(streamURLs)
}

func getAllStreamURLs(songInformation types.SongInformation) []string {
	var streamURLs []string
	for _, provider := range providers {
		streamURLs = append(streamURLs, provider.GetStreamURL(songInformation))
	}
	return streamURLs
}

func findCurrentProvider(providers []types.StreamingProvider, url string) types.StreamingProvider {
	var currentProvider types.StreamingProvider
	for _, provider := range providers {
		if provider.IsThisProvider(url) {
			currentProvider = provider
			break
		}
	}
	return currentProvider
}
