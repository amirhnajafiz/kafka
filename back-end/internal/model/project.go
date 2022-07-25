package model

type Project struct {
	Id          int    `bson:"id,omitempty"`
	Title       string `bson:"title,omitempty"`
	Description string `bson:"description,omitempty"`
	Link        string `bson:"link,omitempty"`
	Visible     bool   `bson:"visible,omitempty"`
}
