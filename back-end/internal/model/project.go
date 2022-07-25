package model

type Project struct {
	Title       string `bson:"title,omitempty"`
	Description string `bson:"description,omitempty"`
	Link        string `bson:"link,omitempty"`
	Visible     bool   `bson:"visible,omitempty"`
}
