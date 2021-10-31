## Сompound Telegram bot and YouTube API v3

### This is app for check a last video from channel url

In: url to channel **With id of channel** example: www.youtube.com/channel/[id]
  
Out: the last video on www.youtube.com/channel/[id] channel.

Of checking "This video is last?" per function GetLastVideo:
  ```go
	func GetLastVideo(channelUrl string) (string, error) {
  		items, err := retrieveVideos(channelUrl)
		if err != nil {
			return "", err
		}
	
		if len(items) < 1 {
			return "", errors.New("Error! No videos found!")
		}
	
		return YOUTUBE_VIDEO_URL + items[0].Id.VideoId, nil
	}
  ```
