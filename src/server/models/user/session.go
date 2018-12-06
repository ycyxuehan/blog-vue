package user

import (
	"strconv"
	"math/rand"
	"time"
	"fmt"
	"io"
	"crypto/md5"
	"github.com/ycyxuehan/bingo"

)

const (
	SESSIONTABLE = "session"
	MAX_EFFECTIVE = 600
)

type Session struct{
	bingo.Model
	User string `json:"user"`
	Session string `json:"session"`
	LoginID string `json:"loginid"`
	CreateTime int64 `json:"createtime"`
	ExpireTime int64 `json:"expiretime"`
	Expired bool `json:"expired"`
}

func NewSession(user string)*Session{
	s := Session{
		User:user,
		Session: GenerateSessionID(""),
		CreateTime: time.Now().Unix(),
	}
	s.LoginID = GenerateSessionID(user)
	s.ExpireTime = s.CreateTime + MAX_EFFECTIVE
	s.Expired = false
	s.SetThis(&s)
	return &s
}

func (s *Session)Table()string{
	return SESSIONTABLE
}

func (s *Session)Filter()map[string]interface{}{
	filter := make(map[string]interface{})
	filter["loginid"] = s.LoginID
	return filter
}

func (s *Session)Columns()([]string, []interface{}){
	return []string{"user", "loginid", "session", "createtime", "expiretime", "expired"}, []interface{}{&s.User, &s.LoginID, &s.Session, &s.CreateTime, &s.ExpireTime, &s.Expired}
}

func Md5(text string) string {
    hashMd5 := md5.New()
    io.WriteString(hashMd5, text)
    return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func GenerateSessionID(user string)string{
	nano := time.Now().UnixNano()
    rand.Seed(nano)
    rndNum := rand.Int63()
	sessionID := Md5(user + Md5(strconv.FormatInt(nano, 10))+Md5(strconv.FormatInt(rndNum, 10)))
	return sessionID
}