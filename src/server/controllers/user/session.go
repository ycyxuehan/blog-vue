package user

import (
	"server/models/user"
	"github.com/ycyxuehan/bingo"

)

type SessionController struct{
	bingo.Controller
}

func (s *SessionController)Get(){
	loginID := s.Ctx.Param("loginid")
	bingo.Logger.Info("loginid:%s", loginID)
	if loginID == "" {
		bingo.Logger.Info("login id is empty")
		s.Ctx.ServeJSON(&bingo.Result{ResultCode:10, ResultString:"auth failed", ResultData:nil})
		return
	}
	loginSession := user.NewSession("")
	loginSession.LoginID = loginID
	err := loginSession.Init(bingo.DBConnection)
	if err != nil {
		bingo.Logger.Info("get session for login id :%s error %s", loginID, err)
		s.Ctx.ServeJSON(&bingo.Result{ResultCode:11, ResultString:"auth failed", ResultData:nil})
		return
	}
	bingo.Logger.Info("get session successed")
	s.Ctx.ServeJSON(&bingo.Result{ResultCode:0, ResultString:"auth successed", ResultData:loginSession})

}