package model

type Post struct {
	PostId  int    `json:"postId"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
type ResponseWithData struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

func (p *Post) IsEmpty() bool {
	return p.Title == "" || p.Content == "" || p.Author == ""
}
