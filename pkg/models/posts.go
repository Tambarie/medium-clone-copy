package models

type Post struct {
	Id string `json:"id"`
	Content string `json:"content"`
	Title string `json:"title"`
	Comments []Comments
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type IPost interface {
	CreatePost(post Post)
	GetAPost(id string)*Post
	GetAllPosts() (posts []Post, err error)
	DeletePost(id string) (err error)
	UpdatePost()(err error)
	//SharePostLink()
}

func (db *DbModel) CreatePost(post Post)  {

}

