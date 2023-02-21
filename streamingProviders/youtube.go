package streamingProviders

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/einnamehalt/musicstreamgenerisizerULTRAMEGAPLUSPLUS/types"
)

var youtubeApiKey string = "<API KEY>"

type APIResponseSnippet struct {
	Title string `json:"title"`
}

type APIResponseId struct {
	VideoId string `json:"videoId"`
}

type APIResponseItem struct {
	Snippet APIResponseSnippet `json:"snippet"`
	Id      APIResponseId      `json:"id"`
}

type APIResponse struct {
	Items []APIResponseItem `json:"items"`
}

type youtubeMusic struct {
	hostname string
}

func (youtubeMusic youtubeMusic) IsThisProvider(url string) bool {
	return strings.Contains(url, youtubeMusic.hostname)
}

func (youtubeMusic youtubeMusic) GetSongInformation(inputUrl string) types.SongInformation {

	u, err := url.Parse(inputUrl)
	if err != nil {
		log.Fatal(err)
	}

	var videoId string = strings.SplitAfter(strings.SplitAfter(u.RawQuery, "=")[1], "&")[0]

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://youtube.googleapis.com/youtube/v3/videos?part=snippet&id="+videoId+"&key="+youtubeApiKey+"&maxResults=1", nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var apiResponse APIResponse
	json.Unmarshal(bodyBytes, &apiResponse)

	return types.SongInformation{Title: apiResponse.Items[0].Snippet.Title, Artist: "", Album: ""}
}

func (youtubeMusic youtubeMusic) GetStreamURL(songInformation types.SongInformation) string {

	client := &http.Client{}
	var url string = "https://youtube.googleapis.com/youtube/v3/search?q=" + url.QueryEscape(songInformation.Title) + "&key=" + youtubeApiKey + "&maxResults=1&part=snippet"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject APIResponse
	json.Unmarshal(bodyBytes, &responseObject)

	return "https://" + youtubeMusic.hostname + "/watch?v=" + responseObject.Items[0].Id.VideoId
}

var YoutubeMusic = youtubeMusic{
	hostname: "music.youtube.com",
}
