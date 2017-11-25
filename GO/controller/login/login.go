package login

import (
	"net/http"
	"../../controller"
	"../../controller/session"
	"fmt"
	"../../errorcode"
	"../../utils/database"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
	controller.AddHeader(w)
	name:=r.FormValue("name")
	pwd:=r.FormValue("pwd")
	fmt.Println(name)
	fmt.Println(pwd)
	if(name==""||pwd==""){
		controller.Respon(w,errorcode.ParmErr,nil)
		return
	}
	level,err:=database.CheckUser(name,pwd)
	if(err!=nil){
		fmt.Println("database.CheckUser:"+err.Error())
		controller.Respon(w,errorcode.PwdErr,nil)
		return
	}
	fmt.Println("level:"+strconv.FormatInt(level, 10))
	session.SetSession(w,name,level)
	controller.Respon(w,errorcode.SuccessCode,nil)
}