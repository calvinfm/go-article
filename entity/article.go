package entity

type Article struct {
	Id      int    `json:"id"`
	Author  string `json:"author" form:"author"`
	Title   string `json:"title" form:"title"`
	Body    string `json:"body" form:"body"`
	Created string `json:"created"`
}

type AddArticle struct {
	Author string `json:"author" form:"author"`
	Title  string `json:"title" form:"title"`
	Body   string `json:"body" form:"body"`
}
