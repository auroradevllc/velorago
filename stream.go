package velora

import "time"

// UserStream is a stream retrieved by the user stream endpoint
// /api/streams/user/<username>
type UserStream struct {
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
	Channel Channel    `json:"channel"`
	Stream  UserStream `json:"stream"`
}

type Category struct {
	ID               string `json:"id"`
	Slug             string `json:"slug"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	ImageURL         string `json:"imageUrl"`
	ExternalImageURL string `json:"externalImageUrl"`
	AssetPath        string `json:"assetPath"`
	Source           string `json:"source"`
	CategoryType     string `json:"categoryType"`
	Viewers          int    `json:"viewers"`
	Streamers        int    `json:"streamers"`
}

type ListStreamResponse struct {
	Streams []ListStream `json:"streams"`
}

// ListStream is a response type for the listing endpoints, which return different structures
// than the UserStream endpoint.
type ListStream struct {
	ID           string `json:"id"`
	UserID       string `json:"userId"`
	Username     string `json:"username"`
	DisplayName  string `json:"displayName"`
	Title        string `json:"title"`
	Category     string `json:"category"`
	CategorySlug string `json:"categorySlug"`
	ViewerCount  int    `json:"viewerCount"`
	Thumbnail    string `json:"thumbnail"`
	// Avatar has an alternate avatarUrl which is populated with the same data
	Avatar string `json:"avatar"`
	// Banner has an alternate bannerUrl which is populated the same
	Banner string `json:"banner"`

	Tags          []string      `json:"tags"`
	SessionTags   []string      `json:"sessionTags"`
	IdentityTags  []string      `json:"identityTags"`
	Language      string        `json:"language"`
	Languages     []string      `json:"languages"`
	StartedAt     time.Time     `json:"startedAt"`
	FollowerCount int           `json:"followerCount"`
	AccentColor   *string       `json:"accentColor"`
	PlaylistURL   string        `json:"playlistUrl"`
	IsFirstStream bool          `json:"isFirstStream"`
	ProfileTheme  *ProfileTheme `json:"profileTheme"`
	IngestSource  string        `json:"ingestSource"`
}
