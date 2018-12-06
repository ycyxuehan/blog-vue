package user

import (
	"github.com/ycyxuehan/bingo"
)

const TABLE = "user"

type User struct{
	bingo.Model
	Name string `json:"name"`
	Password string `json:"password"`
	Enabled bool `json:"enabled"`
}

func New(name string, passwd string)*User{
	u := User{
		Model:bingo.Model{},
		Name: name,
		Password: passwd,
		Enabled:true,
	}
	u.SetThis(&u)
	return &u
}

//Filter database query filter
func (u *User)Filter()map[string]interface{}{
	filter := make(map[string]interface{})
	filter["name"] = &u.Name
	return filter
}

//Table return database table
func (u *User)Table()string{
	return TABLE
}

//Columns return columns map
func (u *User)Columns()([]string, []interface{}){
	return []string{"name", "password", "enabled"}, []interface{}{&u.Name, &u.Password, &u.Enabled}
}

//