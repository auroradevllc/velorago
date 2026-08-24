package velora

import "time"

type User struct {
	ID                string            `json:"id"`
	Username          string            `json:"username"`
	DisplayName       string            `json:"displayName"`
	Bio               string            `json:"bio"`
	AvatarURL         string            `json:"avatarUrl"`
	Role              string            `json:"role"`
	Status            string            `json:"status"`
	FollowerCount     int               `json:"followerCount"`
	TotalViews        int               `json:"totalViews"`
	TotalStreamCount  int               `json:"totalStreamCount"`
	TotalStreamHours  float64           `json:"totalStreamHours"`
	CreatedAt         time.Time         `json:"createdAt"`
	YoutubeConnected  bool              `json:"youtubeConnected"`
	YoutubeChannelID  *string           `json:"youtubeChannelId"`
	StaffBadgeType    *string           `json:"staffBadgeType"`
	StaffBadgeVisible bool              `json:"staffBadgeVisible"`
	ActiveTeamBadgeID *string           `json:"activeTeamBadgeId"`
	CanMonetize       bool              `json:"canMonetize"`
	Pronouns          string            `json:"pronouns"`
	PronounsShort     string            `json:"pronounsShort"`
	PronounsInChat    bool              `json:"pronounsInChat"`
	StreamTags        []string          `json:"streamTags"`
	StreamInfo        *UserStream       `json:"streamInfo"`
	IsLive            bool              `json:"isLive"`
	AccentColor       string            `json:"accentColor"`
	AllowEmbed        bool              `json:"allowEmbed"`
	LastLiveAt        time.Time         `json:"lastLiveAt"`
	ProfileTheme      *ProfileTheme     `json:"profileTheme"`
	ProfileLinks      []ProfileLink     `json:"profileLinks"`
	ProfileFeatured   []ProfileFeatured `json:"profileFeatured"`
	ProfileCanvas     *ProfileCanvas    `json:"profileCanvas"`
	CanvasStudioData  interface{}       `json:"canvasStudioData"`
	Badges            []interface{}     `json:"badges"`
	TeamMemberships   []interface{}     `json:"teamMemberships"`
}

type ProfileTheme struct {
	AccentColor                     *string `json:"accentColor"`
	DerivedAccentColor              *string `json:"derivedAccentColor"`
	TextureKey                      *string `json:"textureKey"`
	TextureIntensity                int     `json:"textureIntensity"`
	BackgroundEnabled               bool    `json:"backgroundEnabled"`
	CommunityLabelSingular          string  `json:"communityLabelSingular"`
	CommunityLabelPlural            string  `json:"communityLabelPlural"`
	CommunityLabelSingularCustom    *string `json:"communityLabelSingularCustom"`
	CommunityLabelPluralCustom      *string `json:"communityLabelPluralCustom"`
	SubscriberLabelSingular         string  `json:"subscriberLabelSingular"`
	SubscriberLabelPlural           string  `json:"subscriberLabelPlural"`
	SubscriberLabelSingularCustom   *string `json:"subscriberLabelSingularCustom"`
	SubscriberLabelPluralCustom     *string `json:"subscriberLabelPluralCustom"`
	DisplayFontFamily               *string `json:"displayFontFamily"`
	DisplayFontVariant              *string `json:"displayFontVariant"`
	DisplayFontHref                 *string `json:"displayFontHref"`
	DisplayFontSizeDelta            int     `json:"displayFontSizeDelta"`
	ChannelPointsNameSingular       string  `json:"channelPointsNameSingular"`
	ChannelPointsNamePlural         string  `json:"channelPointsNamePlural"`
	ChannelPointsNameSingularCustom *string `json:"channelPointsNameSingularCustom"`
	ChannelPointsNamePluralCustom   *string `json:"channelPointsNamePluralCustom"`
	ChannelPointsIcon               *string `json:"channelPointsIcon"`
	ChannelPointsIconCategory       *string `json:"channelPointsIconCategory"`
}

type ProfileLink struct {
	Type     string `json:"type"`
	URL      string `json:"url"`
	Title    string `json:"title,omitempty"`
	Visible  bool   `json:"visible"`
	IconType string `json:"iconType"`
}

type ProfileFeatured struct {
	ChannelID   string                 `json:"channelId"`
	Headline    string                 `json:"headline"`
	Tagline     string                 `json:"tagline"`
	AccentColor *string                `json:"accentColor"`
	Stream      *ProfileFeaturedStream `json:"stream"`
}

type ProfileFeaturedStream struct {
	ID                   string `json:"id"`
	Username             string `json:"username"`
	DisplayName          string `json:"displayName"`
	Title                string `json:"title"`
	Category             string `json:"category"`
	ViewerCount          int    `json:"viewerCount"`
	Avatar               string `json:"avatar"`
	FollowerCount        int    `json:"followerCount"`
	CommunityLabelPlural string `json:"communityLabelPlural"`
	IsLive               bool   `json:"isLive"`
}

type ProfileCanvas struct {
	Orientation     string        `json:"orientation"`
	Height          int           `json:"height"`
	BackgroundColor *string       `json:"backgroundColor"`
	Blocks          []interface{} `json:"blocks"`
}
