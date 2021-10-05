package models

type Login struct {
	email string `json:"email"`
	password string`json:"password"`
}


type ILogin interface {
	Login(u *User) error
}

func (db *DbModel) Login(u *User)error  {
	return nil
}




