package youtube

type RestResponse struct {
	Items []Item    `json:"item"`
}

type Item struct {
	Id ItemInfo     `json:"id"`
}

type ItemInfo struct {
	VideoId string  `json:"video_id"`
}