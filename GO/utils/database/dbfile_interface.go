package database

import (
	"errors"
	"fmt"
)

func UploadDBFile(filename string,filedata []byte)error{
	if(filename==""||filedata==nil){
		return errors.New("parms error")
	}
	sql := "insert into file(filename,filebody) values($1,$2);"
	_,err:=Db.Exec(sql,filename,filedata)
	return err
}


func GetDBFile(id int)(filename string,filebody []byte,err error){
	if(id<0){
		return "",nil,errors.New("parm error")
	}
	sql := "select filename,filebody from file where id =$1;"
	row:=Db.QueryRow(sql,id)
	var name string
	var body []byte
	err=row.Scan(&name,&body)
	if(err!=nil){
		return filename,filebody,err
	}
	fmt.Println("ok")
	return name,body,nil
}
