package article

import (
	"github.com/ycyxuehan/bingo"
)

const CONTENT_TABLE = "article_heat"

type ArticleHeat struct{
	bingo.Model
	ArticleID int64 `json:"articleid"`
	Views int `json:"views"`
	Praises int `json:"praises"`
	Collects int `json:"collects"`
}

func NewHeat(articleId int64)*ArticleHeat{
	ah:= ArticleHeat{
		ArticleID: articleId,
	}
	ah.SetThis(&ah)
	return &ah
}

func (ah*ArticleHeat)Table()string{
	return CONTENT_TABLE
}

func (ah*ArticleHeat)Filter()map[string]interface{}{
	filter := make(map[string]interface{})
	filter["articleid"] = ah.ArticleID
	return filter
}

func (ah*ArticleHeat)Columns()([]string, []interface{}){
	return []string{"articleid", "views", "praises", "collects"}, []interface{}{&ah.ArticleID, &ah.Views, &ah.Praises, &ah.Collects}
}