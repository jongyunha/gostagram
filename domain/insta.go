package domain

type Instagram struct {
	Caption      string   `json:"caption"`
	Id           string   `json:"id"`
	MediaType    string   `json:"media_type"`
	MediaURL     string   `json:"media_url"`
	Permalink    string   `json:"permalink"`
	ThumbnailURL string   `json:"thumbnail_url"`
	Timestamp    string   `json:"timestamp"`
	Username     string   `json:"username"`
	Children     Children `json:"children,omitempty"`
}

type InstagramData struct {
	Data []Instagram `json:"data"`
}

type Children struct {
	Data []Child `json:"data"`
}

type Child struct {
	Id           string `json:"id"`
	MediaType    string `json:"media_type"`
	MediaURL     string `json:"media_url"`
	Permalink    string `json:"permalink"`
	ThumbnailURL string `json:"thumbnail_url"`
}
