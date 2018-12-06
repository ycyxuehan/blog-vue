package user

import (
	"encoding/json"
	"strconv"
	"fmt"
	"server/models/user"
	"github.com/ycyxuehan/bingo"

)

type UserController struct{
	bingo.Controller
}

type LoginData struct{
	User string	`json:"username"`
	Password string `json:"password"`
}

func (u *UserController)Post(){
	loginData := LoginData{}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &loginData)
	if err != nil {
		bingo.Logger.Error("unmarshal data error: %s", err)
		u.Ctx.ServeJSON(bingo.Result{ResultCode:3, ResultString:fmt.Sprintf("unmarshal data error: %s", err), ResultData:nil})
		return
	}

	if loginData.User == "" {
		bingo.Logger.Error("user name is empty")
		u.Ctx.ServeJSON(bingo.Result{ResultCode:3,ResultString:"user name is empty", ResultData:nil })
		return 
	}
	us := user.New(loginData.User, "")
	err = us.Init(bingo.DBConnection)
	if err != nil {
		bingo.Logger.Error("get user from database error: %s", err)
		u.Ctx.ServeJSON(bingo.Result{ResultCode:4,ResultString:fmt.Sprintf("get user from database error: %s", err) , ResultData:nil })
		return 
	}
	if loginData.Password != us.Password {
		bingo.Logger.Error("password is not invalidate of user: %s", loginData.User)
		u.Ctx.ServeJSON(bingo.Result{ResultCode:4,ResultString:fmt.Sprintf("password is not invalidate of user: %s", loginData.User) , ResultData:nil })
		return 
	}
	session := user.NewSession(loginData.User)
	if i, e := strconv.Atoi(bingo.BingConf.Get("sessiontime")); e == nil && i>0 {
		session.ExpireTime = session.CreateTime + int64(i)
	}
	_, err = session.Save(bingo.DBConnection)
	if err != nil {
		bingo.Logger.Error("save session error: %s", err)
		u.Ctx.ServeJSON(bingo.Result{ResultCode:4,ResultString:fmt.Sprintf("save session error: %s", err) , ResultData:nil })
		return 
	}

	bingo.Logger.Info("login success")
	u.Ctx.ServeJSON(bingo.Result{ResultCode:0,ResultString:"success" , ResultData:session })
}