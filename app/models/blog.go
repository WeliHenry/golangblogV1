package models

// Blog model that holds informations of a blog post.
type Blog struct {
	ID      string `json:"id,omitempty "bson:"_id,omitempty"`
	Title   string `json:"title,omitempty" bson:"title,omitempty"`
	Content string `json:"content,omitempty" bson:"content,omitempty"`
	Author  string `json:"author,omitempty" bson:"author,omitempty"`
}
