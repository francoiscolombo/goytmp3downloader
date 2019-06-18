package utube

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/schollz/progressbar"
)

// NbLinksToRetrieve is the maximum number of link to display when we search for a video
const NbLinksToRetrieve = 50

// URLMetaData is the format string used to get metadata regarding a video
const URLMetaData = "https://www.youtube.com/get_video_info?video_id=%s"

// VideoFormatToDownload is the prefered video formats list, the first in the list have priority
const VideoFormatToDownload = "video/mp4"

// ThumbnailValue is part of Thumbnails
type ThumbnailValue struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Thumbnails is part of Snippet
type Thumbnails struct {
	DefaultSize ThumbnailValue `json:"default"`
	MediumSize  ThumbnailValue `json:"medium"`
	HighSize    ThumbnailValue `json:"high"`
}

// Snippet is part of Item
type Snippet struct {
	PublishedAt          string     `json:"publishedAt"`
	ChannelID            string     `json:"channelId"`
	Title                string     `json:"title"`
	Description          string     `json:"description"`
	Thumbnails           Thumbnails `json:"thumbnails"`
	ChannelTitle         string     `json:"channelTitle"`
	LiveBroadcastContent string     `json:"liveBroadcastContent"`
}

// IDValue is part of Item
type IDValue struct {
	Kind    string `json:"kind"`
	VideoID string `json:"videoId"`
}

// Item is part of Results
type Item struct {
	Kind    string  `json:"kind"`
	Etag    string  `json:"etag"`
	ID      IDValue `json:"id"`
	Snippet Snippet `json:"snippet"`
}

// PageInfo is part of Results
type PageInfo struct {
	TotalResults   int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

// SearchResults is the response receive from youtube API
type SearchResults struct {
	Kind          string   `json:"kind"`
	Etag          string   `json:"etag"`
	NextPageToken string   `json:"nextPageToken"`
	RegionCode    string   `json:"regionCode"`
	PageInfo      PageInfo `json:"pageInfo"`
	Items         []Item   `json:"items"`
}

// VideoFormat is part of VideoMetaData
type VideoFormat struct {
	Itag      int
	VideoType string
	Quality   string
	URL       string
}

// VideoMetaData contains all the datas needed to download the video
type VideoMetaData struct {
	ID            string
	Date          string
	Title         string
	Description   string
	Author        string
	Keywords      string
	ThumbnailURL  string
	AvgRating     float32
	ViewCount     int
	LengthSeconds int
	Format        VideoFormat
}

// SearchVideos is used to search a list of videos containing the search string that you pass.
func SearchVideos(query, key string) (results []VideoMetaData, err error) {
	searchResults := SearchResults{}
	urlSearch := fmt.Sprintf("https://www.googleapis.com/youtube/v3/search?part=snippet&maxResults=%d&order=relevance&q=%s&safeSearch=strict&type=video&videoSyndicated=true&key=%s", NbLinksToRetrieve, url.QueryEscape(query), key)
	req, _ := http.NewRequest("GET", urlSearch, nil)
	req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = fmt.Errorf("status code error while trying to search for videos from %s: %d %s", urlSearch, res.StatusCode, res.Status)
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &searchResults)
	bar := progressbar.NewOptions(len(searchResults.Items))
	bar.RenderBlank()
	for _, item := range searchResults.Items {
		bar.Add(1)
		metadata, e := GetVideoMetaData(item.ID.VideoID)
		if e == nil {
			itemTitle, _ := url.QueryUnescape(item.Snippet.Title)
			t, _ := time.Parse(time.RFC3339, item.Snippet.PublishedAt)
			metadata.ID = item.ID.VideoID
			metadata.Title = itemTitle
			metadata.Date = t.Format("Monday, 02-Jan-06")
			metadata.Description = item.Snippet.Description
			results = append(results, metadata)
		}
	}
	fmt.Println()
	return
}

// GetVideoMetaData download metadatas for the video identified by the id
func GetVideoMetaData(id string) (metadatas VideoMetaData, err error) {
	err = nil
	urlMetadata := fmt.Sprintf(URLMetaData, url.QueryEscape(id))
	req, _ := http.NewRequest("GET", urlMetadata, nil)
	req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		err = fmt.Errorf("status code error while trying to search for videos from %s: %d %s", urlMetadata, res.StatusCode, res.Status)
		return
	}
	body, _ := ioutil.ReadAll(res.Body)
	metadatasQueryString := string(body)
	u, _ := url.Parse("?" + metadatasQueryString)
	query := u.Query()
	formatParams := strings.Split(query.Get("url_encoded_fmt_stream_map"), ",")
	if query.Get("errorcode") != "" || query.Get("status") == "fail" {
		err = errors.New(query.Get("reason"))
		return
	}
	foundGoodFormat := false
	for _, formatParam := range formatParams {
		formatParamURL, _ := url.Parse("?" + formatParam)
		formatParamQuery := formatParamURL.Query()
		itag, _ := strconv.Atoi(formatParamQuery.Get("itag"))
		if strings.HasPrefix(formatParamQuery.Get("type"), VideoFormatToDownload) {
			rc, ec := http.Head(formatParamQuery.Get("url"))
			if ec == nil {
				if rc.StatusCode != 403 {
					foundGoodFormat = true
					metadatas.Format.Itag = itag
					metadatas.Format.VideoType = formatParamQuery.Get("type")
					metadatas.Format.Quality = formatParamQuery.Get("quality")
					metadatas.Format.URL = formatParamQuery.Get("url")

				}
			}
		}
	}
	if foundGoodFormat == false {
		err = errors.New("No format found for this video, we have to skip it")
		return
	}
	viewCount, _ := strconv.Atoi(query.Get("view_count"))
	avgRating, _ := strconv.ParseFloat(query.Get("avg_rating"), 32)
	lengthSec, _ := strconv.Atoi(query.Get("length_seconds"))
	metadatas.ID = id
	metadatas.Title = query.Get("title")
	metadatas.Author = query.Get("author")
	metadatas.Keywords = query.Get("keywords")
	metadatas.ThumbnailURL = query.Get("thumbnail_url")
	metadatas.ViewCount = viewCount
	metadatas.AvgRating = float32(avgRating)
	metadatas.LengthSeconds = lengthSec
	return
}
