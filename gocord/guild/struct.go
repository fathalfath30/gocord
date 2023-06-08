/*
//
//  ______    _   _           _  __      _   _     ____   ___
// |  ____|  | | | |         | |/ _|    | | | |   |___ \ / _ \
// | |__ __ _| |_| |__   __ _| | |_ __ _| |_| |__   __) | | | |
// |  __/ _` | __| '_ \ / _` | |  _/ _` | __| '_ \ |__ <| | | |
// | | | (_| | |_| | | | (_| | | || (_| | |_| | | |___) | |_| |
// |_|  \__,_|\__|_| |_|\__,_|_|_| \__,_|\__|_| |_|____/ \___/
//
// Written by Fathalfath30.
// Email : fathalfath30@gmail.com
// Follow me on:
//  Github : https://github.com/fathalfath30
//  Gitlab : https://gitlab.com/Fathalfath30
//
*/

package guild

type (
	Roles struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		Permission  string `json:"permission"`
		Position    uint32 `json:"position"`
		Color       uint32 `json:"color"`
		Hoist       bool   `json:"hoist"`
		Managed     bool   `json:"managed"`
		Mentionable bool   `json:"mentionable"`
	}

	Emoji struct {
		Id            string   `json:"id"`
		Name          string   `json:"name"`
		Roles         []string `json:"roles"`
		RequireColons bool     `json:"require_colons"`
		Managed       bool     `json:"managed"`
		Animated      bool     `json:"animated"`
		Available     bool     `json:"available"`
	}

	Guild struct {
		Id                         string   `json:"id"`
		Name                       string   `json:"name"`
		Icon                       string   `json:"icon"`
		Description                *string  `json:"description"`
		Splash                     string   `json:"splash"`
		DiscoverySplash            *string  `json:"discovery_splash"`
		ApproximateMemberCount     uint64   `json:"approximate_member_count"`
		ApproximatePresenceCount   uint64   `json:"approximate_presence_count"`
		Features                   []string `json:"features"`
		Emojis                     []*Emoji `json:"emojis"`
		Banner                     string   `json:"banner"`
		OwnerId                    string   `json:"owner_id"`
		ApplicationId              *string  `json:"application_id"`
		Region                     *string  `json:"region"`
		AfkChannelId               *string  `json:"afk_channel_id"`
		AfkTimeout                 uint64   `json:"afk_timeout"`
		SystemChannelId            *string  `json:"system_channel_id"`
		WidgetEnabled              bool     `json:"widget_enabled"`
		WidgetChannelId            string   `json:"widget_channel_id"`
		VerificationLevel          uint32   `json:"verification_level"`
		Roles                      []*Roles `json:"roles"`
		DefaultMessageNotification uint64   `json:"default_message_notification"`
		MfaLevel                   uint32   `json:"mfa_level"`
		ExplicitContentFilter      uint32   `json:"explicit_content_filter"`
		MaxPresence                *string  `json:"max_presence"`
		MaxMembers                 uint64   `json:"max_members"`
		MaxVideoChannelUser        uint64   `json:"max_video_channel_user"`
		VanityUrlCode              string   `json:"vanity_url_code"`
		PremiumTier                uint32   `json:"premium_tier"`
		PremiumSubscriptionCount   uint32   `json:"premium_subscription_count"`
		SystemChannelFlag          uint32   `json:"system_channel_flag"`
		PreferredLocale            string   `json:"preferred_locale"`
		RulesChannelId             *string  `json:"rules_channel_id"`
		PublicUpdateChannelId      *string  `json:"public_update_channel_id"`
		SafetyAlertChannelId       *string  `json:"safety_alert_channel_id"`
	}
)
