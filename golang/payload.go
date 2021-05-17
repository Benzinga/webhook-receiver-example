package bzWebhook

import (
	"time"
)

type Security struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	Primary  bool   `json:"primary"`
}

type Content struct {
	Id         int64 `json:"id"`
	RevisionId int    `json:"revision_id"`
	Type       string `json:"type"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`

	Title      string     `json:"title"`
	Body       string     `json:"body"`
	Authors    []string   `json:"authors"`
	Teaser     string     `json:"teaser"`
	Url        string     `json:"url"`
	Tags       []string   `json:"tags"`
	Securities []Security `json:"securities"`
	Channels   []string   `json:"channels"`
}

type Data struct {
	Action    string    `json:"action"`
	Id        int64     `json:"id"`
	Content   Content   `json:"content"`
	Timestamp time.Time `json:"timestamp"`
}

type Body struct {
	Id         string `json:"id"`
	ApiVersion string `json:"api_version"`
	Kind       string `json:"kind"`
	Data       Data   `json:"data"`
}
