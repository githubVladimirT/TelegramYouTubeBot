package youtube

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const YOUTUBE_SEARCH_URL = "https://www.googleapis.com/youtube/v3/search"
const YOUTUBE_API_TOKEN = "AIzaSyAnp0mE7jOsUgrYdC8YvcWMIK1_B4gwMoE"
const YOUTUBE_VIDEO_URL = "https://www.youtube.com/watch?v="

// GET https://youtube.googleapis.com/youtube/v3/search?part=id&channelId=[id]&maxResults=1&order=date&key=[YOUR_API_KEY] HTTP/1.1

// Authorization: Bearer [YOUR_ACCESS_TOKEN]
// Accept: application/json

func GetLastVideo(channelUrl string, maxResults int) (string, error) {
	items, err := retrieveVideos(channelUrl, maxResults)
	if err != nil {
		return "", err
	}

	if len(items) < 1 {
		return "", errors.New("Error! No videos found!")
	}

	return YOUTUBE_VIDEO_URL + items[0].Id.VideoId, nil
}

func retrieveVideos(channelUrl string, maxResults int) ([]Item, error) {
	req, err := makeRequest(channelUrl, maxResults)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Printf("rsp %s\n", body)
	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}

	fmt.Printf("rsp items %#+v\n", restResponse)
	return restResponse.Items, nil
}

func makeRequest(channelUrl string, maxResults int) (*http.Request, error) {
	lastSlashIndex := strings.LastIndex(channelUrl, "/")
	channelId := channelUrl[lastSlashIndex+1:]

	req, err := http.NewRequest("GET", YOUTUBE_SEARCH_URL, nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("part", "id")
	query.Add("channelId", channelId)
	query.Add("maxResults", strconv.Itoa(maxResults))
	query.Add("order", "date")
	query.Add("key", YOUTUBE_API_TOKEN)

	req.URL.RawQuery = query.Encode()

	fmt.Println(req.URL.String())

	return req, nil
}
