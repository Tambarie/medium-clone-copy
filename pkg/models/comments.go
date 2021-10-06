package models

type Comments struct {
	commentId int
	Content string
	Author string
	Post *Post
}

