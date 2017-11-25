package database

import "errors"

//获取用户等级
func CheckUser(name string,pwd string)(int64,error){
	if(name==""||pwd==""){
		return 0,errors.New("parm nil")
	}
	sql := "select level from chyuser where name = $1 and pwd =$2;"
	row := Db.QueryRow(sql, name,pwd)
	var level NullInt
	err := row.Scan(&level.Nullint64)
	if err != nil {
		return 0,err
	}
	if !level.Nullint64.Valid {
		return 0, errors.New("no level")
	}
	return level.Nullint64.Int64,nil
}

func GetUserId(name string)(int64,error){
	if(name==""){
		return 0,errors.New("parm nil")
	}
	sql := "select id from chyuser where name = $1;"
	row := Db.QueryRow(sql, name)
	var id NullInt
	err := row.Scan(&id.Nullint64)
	if err != nil {
		return 0,err
	}
	if !id.Nullint64.Valid {
		return 0, errors.New("no id")
	}
	return id.Nullint64.Int64,nil
}
