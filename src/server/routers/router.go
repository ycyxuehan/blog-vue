package routers

import (
	"server/controllers/user"
	"server/controllers/category"
	"github.com/ycyxuehan/bingo"
	"server/controllers/article"
	
	
)

func init(){
	bingo.RegisterRoute(
		bingo.NewHTTPEntry("/api/v1/article/{id}", &article.ArticleController{}),
		bingo.NewHTTPEntry("/api/v1/article/heat/{id}", &article.HeatController{}),
		bingo.NewHTTPEntry("/api/v1/category", &category.CategoryController{}),
		bingo.NewHTTPEntry("/api/v1/user/login", &user.UserController{}),
		bingo.NewHTTPEntry("/api/v1/user/session", &user.SessionController{}),
	)
}