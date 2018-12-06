package article

import (
	"encoding/json"
	"github.com/ycyxuehan/bingo"
	"fmt"
	"strconv"
	marticle "server/models/article"
)

type HeatController struct{
	bingo.Controller
}

func (h *HeatController)Get(){
	id := h.Ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil || idInt == 0 {
		bingo.Logger.Error("id error, please check it: %s", err)
		h.Ctx.ServeJSON(bingo.Result{ResultCode:2,ResultString:fmt.Sprintf("id error, please check it: %s", err), ResultData:nil })
		return 
	}
	arh := marticle.NewHeat(int64(idInt))
	err = arh.Init(bingo.DBConnection)
	if err != nil {
		bingo.Logger.Error("id error, please check it: %s", err)
		h.Ctx.ServeJSON(bingo.Result{ResultCode:2,ResultString:fmt.Sprintf("article init  error: %s", err), ResultData:nil })
		return 
	}
	bingo.Logger.Info("article init successed")
	h.Ctx.ServeJSON(bingo.Result{
		ResultCode:0,
		ResultString:"sucess",
		ResultData: arh,
	})
}

func (h *HeatController)Put(){
	id := h.Ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil  || idInt == 0 {
		bingo.Logger.Error("id error, please check it: %s", err)
		h.Ctx.ServeJSON(bingo.Result{ResultCode:2,ResultString:fmt.Sprintf("id error, please check it: %s", err), ResultData:nil })
		return 
	}
	arh := marticle.NewHeat(int64(idInt))
	err = json.Unmarshal(h.Ctx.Input.RequestBody, arh)
	if err != nil {
		bingo.Logger.Error("unmarshal data error: %s", err)
		h.Ctx.ServeJSON(bingo.Result{ResultCode:3, ResultString:fmt.Sprintf("unmarshal data error: %s", err), ResultData:nil})
		return
	}
	bingo.Logger.Info("article init successed")
	h.Ctx.ServeJSON(bingo.Result{
		ResultCode:0,
		ResultString:"sucess",
		ResultData: arh,
	})
}