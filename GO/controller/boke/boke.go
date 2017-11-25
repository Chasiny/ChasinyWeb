package boke

import (
	"net/http"
	"../../controller"
	"strconv"
	"fmt"
	"../../utils/database"
	"../../errorcode"
	"../session"
)


func GetAdminBokeList(w http.ResponseWriter, r *http.Request) {
	controller.AddHeader(w)
	list,err:=database.GetBoKeList(1)
	if(err!=nil){
		fmt.Println("getBokeList error:",err.Error())
		controller.Respon(w,errorcode.SysErr,nil)
		return
	}
	fmt.Println(list)
	controller.Respon(w,errorcode.SuccessCode,list)
}


func GetAdminBokeDetail(w http.ResponseWriter, r *http.Request) {
	controller.AddHeader(w)
	id:=r.FormValue("bokeid")
	if(id==""){
		controller.Respon(w,errorcode.ParmErr,nil)
		return
	}
	idInt,err:=strconv.Atoi(id)
	if(err!=nil){
		fmt.Println("GetBokeDetail:strconv.Atoi(id) error:"+err.Error())
		controller.Respon(w,errorcode.ParmErr,nil)
		return
	}
	boke,err:=database.GetBoKe(int64(idInt))
	if(err!=nil){
		fmt.Println("GetBokeDetail error:"+err.Error())
		controller.Respon(w,errorcode.SysErr,nil)
		return
	}
	fmt.Println(boke)
	controller.Respon(w,errorcode.SuccessCode,boke)
}


func AddBoke() http.HandlerFunc {
	return controller.GateWay(addBoke)
}

func addBoke(w http.ResponseWriter, r *http.Request) {
	title:=r.FormValue("title")
	body:=r.FormValue("body")
	if(title==""||body==""){
		fmt.Println("title or body nil")
		controller.Respon(w,errorcode.ParmErr,nil)
		return
	}
	status:=r.FormValue("status")
	statusInt,err:=strconv.Atoi(status)
	if(err!=nil){
		fmt.Println("statusStr:"+err.Error())
		controller.Respon(w,errorcode.ParmErr,nil)
		return
	}
	name,err:=session.GetName(r)
	if(err!=nil){
		fmt.Println(err)
		controller.Respon(w,errorcode.SysErr,nil)
		return
	}
	id,err:=database.GetUserId(name)
	if(err!=nil){
		fmt.Println(err)
		controller.Respon(w,errorcode.SysErr,nil)
		return
	}
	err=database.InsertBoKe(id,title,body,int64(statusInt))
	fmt.Println("id:"+strconv.FormatInt(id, 10))
	if(err!=nil){
		fmt.Println(err)
		controller.Respon(w,errorcode.SysErr,nil)
		return
	}
	controller.Respon(w,errorcode.SuccessCode,nil)
	return
}


func EditBoke() http.HandlerFunc {
	return controller.GateWay(editBoke)
}

func editBoke(w http.ResponseWriter, r *http.Request) {
	bokeid:=r.FormValue("bokeid")
	title:=r.FormValue("title")
	body:=r.FormValue("body")
	status:=r.FormValue("status")
	if(title==""||body==""||bokeid==""||status==""){
		fmt.Println("title or body or bokeid or status nil")
		controller.Respon(w,errorcode.ParmErr,nil)
		return
	}
	statusInt,err:=strconv.Atoi(status)
	if(err!=nil){
		fmt.Println("statusStr:"+err.Error())
		controller.Respon(w,errorcode.ParmErr,nil)
		return
	}
	bokeidInt,err:=strconv.Atoi(bokeid)
	if(err!=nil){
		fmt.Println("bokeidInt:"+err.Error())
		controller.Respon(w,errorcode.ParmErr,nil)
		return
	}
	name,err:=session.GetName(r)
	if(err!=nil){
		fmt.Println(err)
		controller.Respon(w,errorcode.SysErr,nil)
		return
	}
	id,err:=database.GetUserId(name)
	if(err!=nil){
		fmt.Println(err)
		controller.Respon(w,errorcode.SysErr,nil)
		return
	}
	err=database.UpdateBoKe(id,int64(bokeidInt),title,body,int64(statusInt))
	if(err!=nil){
		fmt.Println(err)
		controller.Respon(w,errorcode.SysErr,nil)
		return
	}
	controller.Respon(w,errorcode.SuccessCode,nil)
	return
}



func GetBokeList() http.HandlerFunc {
	return controller.GateWay(getBokeList)
}

func getBokeList(w http.ResponseWriter, r *http.Request) {
	name,err:=session.GetName(r)
	if(err!=nil){
		controller.Respon(w,errorcode.SysErr,nil)
		return
	}
	id,err:=database.GetUserId(name)
	if(err!=nil){
		controller.Respon(w,errorcode.SysErr,nil)
		return
	}
	list,err:=database.GetBoKeList(id)
	if(err!=nil){
		controller.Respon(w,errorcode.SysErr,nil)
		return
	}
	fmt.Println(list)
	controller.Respon(w,errorcode.SuccessCode,list)
}

func DeleteBoke() http.HandlerFunc {
	return controller.GateWay(deleteBoke)
}

func deleteBoke(w http.ResponseWriter, r *http.Request) {
	name,err:=session.GetName(r)
	if(err!=nil){
		controller.Respon(w,errorcode.SysErr,nil)
		return
	}
	id,err:=database.GetUserId(name)
	if(err!=nil){
		controller.Respon(w,errorcode.SysErr,nil)
		return
	}
	bokeid:=r.FormValue("bokeid")
	if(bokeid==""){
		controller.Respon(w,errorcode.ParmErr,nil)
		return
	}
	bokeidInt,err:=strconv.Atoi(bokeid)
	if(err!=nil){
		controller.Respon(w,errorcode.SysErr,nil)
		return
	}
	err=database.DeleteBoKe(id,int64(bokeidInt))
	controller.Respon(w,errorcode.SuccessCode,nil)
}