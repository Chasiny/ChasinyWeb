package controller

import (
	"net/http"
	"../data"
	"encoding/json"
	"fmt"
	"../errorcode"
	"../controller/session"
)

func AddHeader(w http.ResponseWriter, ) {
	w.Header().Set("Access-Control-Allow-Origin", data.WebHost)
	w.Header().Add("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
}

func Respon(w http.ResponseWriter, code int, data interface{}) {
	if w == nil {
		return
	}
	fb := errorcode.FeedBack{
		Code: code,
		Data: data,
	}
	buf, _ := json.Marshal(fb)
	fmt.Fprint(w, string(buf))
}


func GateWay(protectedPage http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AddHeader(w)
		_,err:=session.CheckSession(r)
		if(err!=nil){
			Respon(w,errorcode.SessionErr,nil)
			return
		}
		protectedPage(w, r)
	})
}