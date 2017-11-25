package session

import (
	"net/http"
	"fmt"
	"errors"
	"time"
	"strconv"
	"crypto/md5"
	"encoding/hex"
	"math/rand"
)

//需要测试并发锁

var sesseionManager *SessionManager

func Init(){
	sessionMap=make(map[string]UserSession)
	sesseionManager=&SessionManager{
		Key:"chy",
		SessionTime:60*60*2,
		Count:0,
	}

	ticker := time.NewTicker(3600 * time.Second)
	go func() {
		for {
			fmt.Println("clean session")
			<-ticker.C
			cleanSession()
		}
	}()
}


func MakeSession(w http.ResponseWriter, r *http.Request,name string,level int64)(sessionid string,err error) {
	sessionid,err=CheckSession(r)
	if(err!=nil){
		return SetSession(w,name,level)
	}
	fmt.Println("session still ok")
	return sessionid,nil
}

func CheckSession(r *http.Request)(sid string,err error){
	session,err:=getSession(r)
	if(err!=nil){
		return "",err
	}
	if(time.Now().Unix()-session.CreateTime>sesseionManager.SessionTime){
		return "",errors.New("session is out of date")
	}
	return session.Sid,nil
}

func getSession(r *http.Request)(UserSession,error){
	var session UserSession
	cookie, err:=r.Cookie(sesseionManager.Key)
	if(err!=nil){
		fmt.Println("no cookie")
		return session,errors.New("no cookie")
	}
	fmt.Println(cookie.Value)
	session,ok:=sessionMap[cookie.Value]
	if(!ok){
		return session,errors.New("session not in map")
	}
	return session,nil
}

func GetName(r *http.Request)(string,error){
	s,err:=getSession(r)
	if(err!=nil){
		return "",err
	}
	return s.Name,nil
}


func SetSession(w http.ResponseWriter,name string,level int64)(sessionid string,err error){
	//sesseionManager.SessionLock.Lock()
	//defer sesseionManager.SessionLock.Unlock()
	sessionid=createSessionID()
	userSession:=UserSession{
		Name:name,
		Level:level,
		Sid:sessionid,
		CreateTime:time.Now().Unix(),
	}
	sessionMap[sessionid]=userSession
	newcookie := http.Cookie{
		Name:     sesseionManager.Key,
		Value:    sessionid,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   3600 * 24 * 10,
	}
	http.SetCookie(w,&newcookie)
	return sessionid,nil
}

func createSessionID() string {
	for{
		nowTime := time.Now().UnixNano()
		rand.Seed(nowTime)
		randNum := rand.Int63()
		sesseionManager.Count++;
		result := strconv.FormatInt(nowTime, 10) + strconv.FormatInt(randNum, 10) + strconv.FormatInt(sesseionManager.Count, 10)
		m := md5.New()
		m.Write([]byte(result))
		sid:=hex.EncodeToString(m.Sum(nil))
		fmt.Println("md5:" + sid)
		_,ok:=sessionMap[sid]
		if(!ok){
			return sid
		}
	}
}

func cleanSession(){
	//sesseionManager.SessionLock.Lock()
	//defer sesseionManager.SessionLock.Unlock()
	for key,data:=range sessionMap{
		if time.Now().Unix()-data.CreateTime>sesseionManager.SessionTime{
			delete(sessionMap, key)
		}
	}
}