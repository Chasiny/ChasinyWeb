package database

import (
	"../../data"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

var Db *sql.DB

func Init() {
	var err error
	connstr := "host=" + data.DbHost + " port=" + strconv.FormatInt(data.DbPort, 10) + " user=" + data.DbUser + " password=" + data.DbPwd + " dbname=" + data.DbName + "  sslmode=disable"
	Db, err = sql.Open("postgres", connstr)
	if Db == nil {
		log.Println("Db is nil")
		return
	}
	if err != nil {
		log.Println(err.Error())
		return
	}
	err = Db.Ping()
	if err != nil {
		log.Println(err.Error())
		return
	}

	//初始化表
	initTable()

	log.Println("database init success.")
}

//初始化表
func initTable() {
	err := ExeSQL(`CREATE TABLE IF NOT EXISTS ChyUser (id SERIAL NOT NULL,name varchar(255) NOT NULL UNIQUE,pwd varchar(255) NOT NULL ,level int8  NOT NULL,PRIMARY KEY ("id"));`)
	if err != nil {
		log.Println("ChyUser表创建失败! " + err.Error())
		return
	}
	err = ExeSQL(`insert into chyuser(name,level,pwd) select $1,$2,$3 where not exists(select * from chyuser where name = $4);`,"chy",1,"888888","chy")
	if err != nil {
		log.Println("chy帐号创建失败! " + err.Error())
		return
	}
	err = ExeSQL(`insert into chyuser(name,level,pwd) select $1,$2,$3 where not exists(select * from chyuser where name = $4);`,"hjj",1,"888888","hjj")
	if err != nil {
		log.Println("hjj帐号创建失败! " + err.Error())
		return
	}
	err = ExeSQL(`CREATE TABLE IF NOT EXISTS file (id SERIAL NOT NULL PRIMARY KEY,sha256 VARCHAR(70) ,filename text NOT NULL,filebody bytea NOT NULL)`)
	if err != nil {
		log.Println("file表创建失败! " + err.Error())
		return
	}
	err = ExeSQL(`CREATE TABLE IF NOT EXISTS boke (id SERIAL NOT NULL PRIMARY KEY,userid int8 NOT NULL references chyuser(id),title text not null,body text,status int8 not null default 0,createtime timestamp NOT NULL default now());`)
	if err != nil {
		log.Println("boke表创建失败! " + err.Error())
		return
	}
}

//执行sql,不获取返回值
func ExeSQL(command string, args ...interface{}) error {
	if Db == nil {
		log.Println("db is nil")
		return fmt.Errorf("db is nil")
	}
	stmt, err := Db.Prepare(command)
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		return err
	}
	if args != nil {
		_, err = stmt.Exec(args...)
	} else {
		_, err = stmt.Exec()
	}

	if err != nil {
		return err
	}
	return nil
}

//获取数据，返回数据列表
func ExeSQLforResult(command string, args ...interface{}) (*sql.Rows, error) {
	if Db == nil {
		log.Println("db is nil")
		return nil, fmt.Errorf("db is nil")
	}
	stmt, err := Db.Prepare(command)
	if err != nil {
		return nil, err
	}
	var rows *sql.Rows
	if args != nil {
		rows, err = stmt.Query(args...)
	} else {
		rows, err = stmt.Query()
	}

	if err != nil {
		return nil, err
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	return rows, nil
}
