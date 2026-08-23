package velora

import "time"

type Stream struct {
	ID             string    `json:"id"`
	UserID         string    `json:"userId"`
	Username       string    `json:"username"`
	Title          string    `json:"title"`
	Description    string    `json:"description"` // HTML content
	Category       string    `json:"category"`
	CategorySlug   string    `json:"categorySlug"`
	ThumbnailURL   string    `json:"thumbnailUrl"`
	IsLive         bool      `json:"isLive"`
	ViewerCount    int       `json:"viewerCount"`
	StartedAt      time.Time `json:"startedAt"`
	HLSURL         string    `json:"hlsUrl"`
	WSSignalingURL string    `json:"wsSignalingUrl"`
	Language       string    `json:"language"`
	IngestOrigin   string    `json:"ingestOrigin"`
	IngestSource   string    `json:"ingestSource"`
}

type Channel struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
}

type StreamData struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Category    Category  `json:"category"`
	StartedAt   time.Time `json:"startedAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Duration    int       `json:"duration"`
	PeakViewers int       `json:"peakViewers"`
}

type StreamEvent struct {
	Channel Channel `json:"channel"`
	Stream  Stream  `json:"stream"`
}

type Category struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type ListStream struct {
	Stream

	Tags      []string `json:"tags"`
	Languages []string `json:"languages"`
}
