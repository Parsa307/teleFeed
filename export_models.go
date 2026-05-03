package main

import "time"

// ExportPost is used for JSON export without empty caption
type ExportPost struct {
	ID          int64     `json:"id"`
	Message     string    `json:"message"`
	Date        time.Time `json:"date"`
	Edited      bool      `json:"edited"`
	EditDate    time.Time `json:"edit_date,omitempty"`
	Views       int       `json:"views"`
	Forwards    int       `json:"forwards"`
	Replies     int       `json:"replies"`
	SenderID    int64     `json:"sender_id"`
	SenderName  string    `json:"sender_name"`
	Media       []Media   `json:"media,omitempty"`
	Hashtags    []string  `json:"hashtags,omitempty"`
	Mentions    []string  `json:"mentions,omitempty"`
	Links       []string  `json:"links,omitempty"`
	
	// Caption only included if not empty
	Caption string `json:"caption,omitempty"`
}

// ExportChannelData is used for JSON export
type ExportChannelData struct {
	Info        ChannelInfo  `json:"info"`
	Posts       []ExportPost `json:"posts"`
	LastUpdated int64        `json:"last_updated"`
}

func toExportPost(post Post) ExportPost {
	exportPost := ExportPost{
		ID:         post.ID,
		Message:    post.Message,
		Date:       post.Date,
		Edited:     post.Edited,
		EditDate:   post.EditDate,
		Views:      post.Views,
		Forwards:   post.Forwards,
		Replies:    post.Replies,
		SenderID:   post.SenderID,
		SenderName: post.SenderName,
		Media:      post.Media,
		Hashtags:   post.Hashtags,
		Mentions:   post.Mentions,
		Links:      post.Links,
	}
	
	// Only include caption if not empty
	if post.Caption != "" {
		exportPost.Caption = post.Caption
	}
	
	return exportPost
}

func toExportChannelData(channelData ChannelData) ExportChannelData {
	exportPosts := make([]ExportPost, len(channelData.Posts))
	for i, post := range channelData.Posts {
		exportPosts[i] = toExportPost(post)
	}
	
	return ExportChannelData{
		Info:        channelData.Info,
		Posts:       exportPosts,
		LastUpdated: channelData.LastUpdated,
	}
}
