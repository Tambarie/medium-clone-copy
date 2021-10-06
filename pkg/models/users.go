package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type User struct {
	Id string `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type IUserCRUD interface {
	SignUp(u User) error
	GetUser(u User) (users []User, err error)
	IfEmailExists(u *User, email string) (*User, bool)
	IfPasswordExists(u *User, password string) (*User, bool)
	//UpdateUser(u *User)error
	//Delete(u *User)error


}

type DbModel struct {
	Db *sql.DB
}



func (db *DbModel) SignUp(u User) error {
	statement := `INSERT INTO users(id,firstName,lastName,email,password,createdAt,updatedAt) VALUES($1,$2,$3,$4,$5,$6,$7)`
	_, err := db.Db.Exec(statement,u.Id,u.FirstName,u.LastName,u.Email,u.Password,u.CreatedAt,u.UpdatedAt)
	return err
}

func (db DbModel) GetUser(u User) (users []User,err error) {
	row, err := db.Db.Query(`SELECT id, firstName,lastName,email,password,createdAt,updatedAt FROM users`)
	if err != nil{
		return
	}
	for row.Next(){
		user := User{}
		err = row.Scan(&u.Id,&u.Email,&u.FirstName,&u.LastName,&u.Password,&u.CreatedAt,&u.UpdatedAt)
		if err != nil{
			return
		}
		users = append(users,user)
	}
	row.Close()
	return
}

func (db *DbModel) IfEmailExists(u *User, email string) (*User, bool) {
	row := db.Db.QueryRow(`SELECT email from users where email = $1`,email)
	err := row.Scan(&u.Email)
	if err != nil{
		return nil, false
		//panic(err)
	}
	return u, true
}

func (db *DbModel) IfPasswordExists(u *User, email string) (*User, bool) {
	row := db.Db.QueryRow(`SELECT password from users where email = $1`,email)
	err := row.Scan(&u.Password)
	if err != nil{
		return nil, false
		//panic(err)
	}
	return u, true
}

