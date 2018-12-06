package article

import (
	"github.com/ycyxuehan/bingo"
	"time"
	
)

const TABLE = "article"

type Article struct {
	bingo.Model
	ID int64 `json:"id"`
	CategoryID int `json:"categoryid"`
	Title string `json:"title"`
	Outline string `json:"outline"`
	CreateDate int64 `json:"createdate"`
	Content string `json:"content"`
	Deleted bool `json:"deleted"`
}

func New(id int64)*Article{
	a := Article{
		Deleted: false,
		CreateDate: time.Now().Unix(),
	}
	if id == 0 {
		a.GenerateID()
	} else {
		a.ID = id
	}
	a.SetThis(&a)
	return &a
}

func (a *Article)Filter()map[string]interface{}{
	filter := make(map[string]interface{})
	filter["id"] = a.ID
	return filter
}

func (a *Article)Columns()([]string, []interface{}){
	return []string{"id", "categoryid", "title", "outline", "createdate", "content", "deleted"}, []interface{}{
		&a.ID, &a.CategoryID, &a.Title, &a.Outline, &a.CreateDate, &a.Content, &a.Deleted}
}

func (a *Article)Table()string{
	return TABLE
}

func (a *Article)GenerateID(){
	a.ID = time.Now().Unix()
}