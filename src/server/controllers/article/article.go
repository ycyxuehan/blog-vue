package article

import (
	"encoding/json"
	"github.com/ycyxuehan/bingo"
	"fmt"
	"strconv"
	marticle "server/models/article"
)

type ArticleController struct{
	bingo.Controller
	Test string
}


func (a *ArticleController)Get(){
	id := a.Ctx.Param("id")
	if id == "list" {
		bingo.Logger.Info("list articles")
		filter := make(map[string]interface{})
		filter["deleted"] = false
		limit, err := strconv.Atoi(a.Ctx.Param("limit"))
		if err == nil {
			limit = 0
		}
		articles, err := bingo.List(bingo.DBConnection, filter, marticle.New(0), limit)
		if err != nil {
			bingo.Logger.Error("get article list error: %s", err)
			a.Ctx.ServeJSON(bingo.Result{ResultCode:3,ResultString:fmt.Sprintf("id error, please check it: %s", err), ResultData:nil })
			return 
		}
		for _, art := range articles{
			if a, ok := art.(marticle.Article); ok {
				a.Content = ""
			}
		}
		a.Ctx.ServeJSON(bingo.Result{
			ResultCode:0,
			ResultString:"sucess",
			ResultData: articles,
		})
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt == 0  {
		bingo.Logger.Error("id error, please check it: %s", err)
		a.Ctx.ServeJSON(bingo.Result{ResultCode:2,ResultString:fmt.Sprintf("id error, please check it: %s", err), ResultData:nil })
		return 
	}
	ar := marticle.New(int64(idInt))
	err = ar.Init(bingo.DBConnection)
	if err != nil {
		bingo.Logger.Error("id error, please check it: %s", err)
		a.Ctx.ServeJSON(bingo.Result{ResultCode:2,ResultString:fmt.Sprintf("article init  error: %s", err), ResultData:nil })
		return 
	}
	bingo.Logger.Info("article init successed")
	a.Ctx.ServeJSON(bingo.Result{
		ResultCode:0,
		ResultString:"sucess",
		ResultData: ar,
	})
}
//Put motify
func (a *ArticleController)Put(){
	paramID := a.Ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil  || id == 0 {
		bingo.Logger.Error("id error, please check it: %s", err)
		a.Ctx.ServeJSON(bingo.Result{ResultCode:2,ResultString:fmt.Sprintf("id error, please check it: %s", err), ResultData:nil })
		return 
	}
	arNew := marticle.New(0)
	ar := marticle.New(int64(id))
	err = json.Unmarshal(a.Ctx.Input.RequestBody, arNew)
	if err != nil {
		bingo.Logger.Error("unmarshal data error: %s", err)
		a.Ctx.ServeJSON(bingo.Result{ResultCode:3, ResultString:fmt.Sprintf("unmarshal data error: %s", err), ResultData:nil})
		return
	}
	if arNew.Title != ""{
		ar.Title = arNew.Title
	}
	if arNew.Content != "" {
		ar.Content = arNew.Content
	}
	if arNew.Outline != "" {
		ar.Outline = arNew.Outline
	}
	if arNew.CategoryID >= 0 {
		ar.CategoryID = arNew.CategoryID
	}
	res, err := ar.Save(bingo.DBConnection)
	if err != nil {
		bingo.Logger.Error("unmarshal data error: %s", err)
		a.Ctx.ServeJSON(bingo.Result{ResultCode:0, ResultString:fmt.Sprintf("add article error: %s", err), ResultData:nil})
		return
	}
	bingo.Logger.Info("%s", res)
}

//Post add
func (a *ArticleController)Post(){
	id := a.Ctx.Param("id")
	if id != "-1" {
		a.NotFound()
		return
	}
	ar := marticle.New(0)
	err := json.Unmarshal(a.Ctx.Input.RequestBody, ar)
	if err != nil {
		bingo.Logger.Error("unmarshal data error: %s", err)
		a.Ctx.ServeJSON(bingo.Result{ResultCode:3, ResultString:fmt.Sprintf("unmarshal data error: %s", err), ResultData:nil})
		return
	}
	ar.GenerateID()
	res, err := ar.Save(bingo.DBConnection)
	if err != nil {
		bingo.Logger.Error("save article error: %s", err)
		a.Ctx.ServeJSON(bingo.Result{ResultCode:3, ResultString:fmt.Sprintf("add article error: %s", err), ResultData:nil})
		return
	}

	bingo.Logger.Info("%s", res)
	a.Ctx.ServeJSON(bingo.Result{ResultCode:0, ResultString:"successed", ResultData:nil})
}

func (a *ArticleController)Delete(){
	paramID := a.Ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil  || id == 0 {
		bingo.Logger.Error("id error, please check it: %s", err)
		a.Ctx.ServeJSON(bingo.Result{ResultCode:2,ResultString:fmt.Sprintf("id error, please check it: %s", err), ResultData:nil })
		return 
	}
	ar := marticle.New(int64(id))
	err = ar.Init(bingo.DBConnection)
	if err != nil {
		bingo.Logger.Error("id error, please check it: %s", err)
		a.Ctx.ServeJSON(bingo.Result{ResultCode:2,ResultString:fmt.Sprintf("article init  error: %s", err), ResultData:nil })
		return 
	}
	ar.Deleted = true
	res, err := ar.Save(bingo.DBConnection)
	if err != nil {
		bingo.Logger.Error("unmarshal data error: %s", err)
		a.Ctx.ServeJSON(bingo.Result{ResultCode:3, ResultString:fmt.Sprintf("add article error: %s", err), ResultData:nil})
		return
	}

	bingo.Logger.Info("%s", res)
	a.Ctx.ServeJSON(bingo.Result{ResultCode:0, ResultString:"successed", ResultData:nil})
	
}

func (a *ArticleController)List(){

}