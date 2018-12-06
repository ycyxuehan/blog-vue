package category

import (
	"encoding/json"
	"fmt"
	"server/models/category"
	"strconv"
	"github.com/ycyxuehan/bingo"

)

type CategoryController struct{
	bingo.Controller
}

func (c *CategoryController)Get(){
	bingo.Logger.Info("list articles")
	filter := make(map[string]interface{})
	filter["enabled"] = true
	limit, err := strconv.Atoi(c.Ctx.Param("limit"))
	if err == nil {
		limit = 0
	}
	categories, err := bingo.List(bingo.DBConnection, filter, &category.Category{}, limit)
	if err != nil {
		bingo.Logger.Error("get article list error: %s", err)
		c.Ctx.ServeJSON(bingo.Result{ResultCode:3,ResultString:fmt.Sprintf("id error, please check it: %s", err), ResultData:nil })
		return 
	}
	c.Ctx.ServeJSON(bingo.Result{
		ResultCode:0,
		ResultString:"sucess",
		ResultData: categories,
	})
}

func (c *CategoryController)Put(){

}

func (c *CategoryController)Post(){
	cate := category.New("", "")
	err := json.Unmarshal(c.Ctx.Input.RequestBody, cate)
	if err != nil {
		bingo.Logger.Error("unmarshal data error: %s", err)
		c.Ctx.ServeJSON(bingo.Result{ResultCode:3, ResultString:fmt.Sprintf("unmarshal data error: %s", err), ResultData:nil})
		return
	}
	res, err := cate.Save(bingo.DBConnection)
	if err != nil {
		bingo.Logger.Error("unmarshal data error: %s", err)
		c.Ctx.ServeJSON(bingo.Result{ResultCode:3, ResultString:fmt.Sprintf("unmarshal data error: %s", err), ResultData:nil})
		return
	}
	c.Ctx.ServeJSON(bingo.Result{ResultCode:0, ResultString:"successed", ResultData:res})
}