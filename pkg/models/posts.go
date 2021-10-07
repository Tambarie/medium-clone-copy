package models

import (
	"database/sql"
	"fmt"
)

type Post struct {
	Id string `json:"id"`
	Content string `json:"content"`
	Title string `json:"title"`
	Comments []Comments
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Author string `json:"author"`
}

type IPost interface {
	CreatePost(post *Post) error
	GetAPost(post *Post,id string)*Post
	GetAllPosts() (posts []Post, err error)
	DeletePost(id string) (err error)
	UpdatePost(post *Post)(err error)
	//SharePostLink()
}

type PostModel struct {
	Db *sql.DB
}



func (db *PostModel) CreatePost(post *Post) error {
	statement := `INSERT INTO posts(id,content,createdat,title,updatedat,authorid) VALUES($1,$2,$3,$4,$5,$6)`
	_, err := db.Db.Exec(statement,post.Id,post.Content,post.CreatedAt,post.Title,post.UpdatedAt,post.Author)
	return err

}

func (db *PostModel) GetAllPosts() (posts []Post, err error) {
	rows, err := db.Db.Query(`SELECT id, content,createdat, title,updatedat,authorid FROM posts`)
	if err != nil{
		return nil, err
	}
	for rows.Next(){
		p := Post{}
		err = rows.Scan(&p.Id,&p.Content,&p.CreatedAt,&p.Title,&p.UpdatedAt,&p.Author)
		if err != nil{
			return nil, err
		}
		posts = append(posts,p)
	}
	rows.Close()
	return posts,nil
}


func (db *PostModel) GetAPost(post *Post,id string) *Post {
	row := db.Db.QueryRow(`SELECT id, title, content , authorid FROM posts WHERE id = $1`,id)
	err := row.Scan(&post.Id, &post.Title,&post.Content,&post.Author)
	if err != nil{
		panic(err)
	}
	return post
}

func (db *PostModel) UpdatePost(post *Post) (err error) {
	posts := "posts"
	stmt, err := db.Db.Prepare(fmt.Sprintf("UPDATE %s SET title = $1, content = $2 WHERE id = $3",posts))
	if err != nil{
		panic(err)
	}
	_, err = stmt.Exec(post.Title,post.Content, post.Id)
	return
}

func (db *PostModel) DeletePost(id string) (err error) {
	_, err = db.Db.Query(`DELETE FROM posts where id = $1`, id)
	return err
}










func (db *PostModel) SignUp(u User) error {
	panic("implement me")
}

func (db *PostModel) IfEmailExists(u *User, email string) (*User, bool) {
	panic("implement me")
}

func (db *PostModel) IfPasswordExists(u *User, password string) (*User, bool) {
	panic("implement me")
}

func (db *PostModel) GetUserId(u *User, email string) (*User, error) {
	panic("implement me")
}

func (db *PostModel) GetUser(u User) (users []User, err error) {
	panic("implement me")
}

