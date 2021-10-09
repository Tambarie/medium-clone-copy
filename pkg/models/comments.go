package models

import "database/sql"

type Comments struct {
	CommentId string
	Content string
	CreatedAt string
	AuthorId string
	PostId string
}

type IComments interface {
	CreateComments(comment *Comments) error
	GetAComments(comment *Comments,id string)*Comments
	GetAllComments(id string) (comment []Comments, err error)
	DeleteComments(id string) (err error)
	UpdateComments(comment *Comments)(err error)
	GetAllCommentsByaUser(comment *Comments,id string)(posts []Post, err error)
	//SharePostLink()
}

type CommentModel struct {
	Db *sql.DB
}

func (c *CommentModel) CreateComments(comment *Comments) error {
	_, err := c.Db.Exec(`INSERT INTO comments (id, content,createdat,authorid,postid) VALUES ($1,$2,$3,$4,$5)`,
		&comment.CommentId,&comment.Content,&comment.CreatedAt,&comment.AuthorId,&comment.PostId )
	return err
}


func (c *CommentModel) GetAllComments(id string) (comment []Comments, err error) {
	rows, err := c.Db.Query(`SELECT content,createdat,authorid from comments where postid = $1`,id)
		if err != nil{
			return
		}
		for rows.Next(){
			c := Comments{}
			err = rows.Scan(&c.Content,&c.CreatedAt,&c.AuthorId)
			if err != nil{
				return
			}
			comment = append(comment,c)
		}
		rows.Close()
	return
}

func (c *CommentModel) GetAComments(post *Comments, id string) *Comments {
	panic("implement me")
}
func (c *CommentModel) DeleteComments(id string) (err error) {
	panic("implement me")
}

func (c *CommentModel) UpdateComments(comment *Comments) (err error) {
	panic("implement me")
}

func (c *CommentModel) GetAllCommentsByaUser(comment *Comments, id string) (posts []Post, err error) {
	panic("implement me")
}





