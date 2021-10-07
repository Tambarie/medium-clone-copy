package models

import "database/sql"

type Comments struct {
	commentId string
	Content string
	AuthorId string
	PostId *Post
}

type IComments interface {
	CreateComments(post *Post) error
	GetAComments(post *Post,id string)*Post
	GetAllPComments() (posts []Post, err error)
	DeleteComments(id string) (err error)
	UpdateComments(post *Post)(err error)
	GetAllCommentsByaUser(post *Post,id string)(posts []Post, err error)
	//SharePostLink()
}

type CommentModel struct {
	Db *sql.DB
}




func (c *CommentModel) CreateComments(post *Post) error {
	panic("implement me")
}

func (c *CommentModel) GetAComments(post *Post,id string)*Post{
	panic("implement me")
}

func (c *CommentModel) GetAllPComments() (posts []Post, err error) {
	panic("implement me")
}

func (c CommentModel) DeleteComments(id string) (err error) {
	panic("implement me")
}

func (c CommentModel) UpdateComments(post *Post)(err error){
	panic("implement me")
}

func (c CommentModel) GetAllCommentsByaUser(post *Post,id string)(posts []Post, err error) {
	panic("implement me")
}

