package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DbWorker struct {
	Dsn  string
	Db   *sql.DB
	User user
}

type user struct {
	Id         int32
	Username   sql.NullString
	Password   sql.NullString
	Nickname   sql.NullString
	Email      sql.NullString
	Telephone  sql.NullString
	RealName   sql.NullString
	Enabled    sql.NullBool
	CreateTime time.Time
}

func main() {
	var err error
	dbw := DbWorker{
		Dsn: "root:Xm123456.@tcp(192.168.1.220:3306)/maynrent?charset=utf8mb4&parseTime=true",
	}
	dbw.Db, err = sql.Open("mysql", dbw.Dsn)
	if err != nil {
		panic(err)
		return
	}
	defer dbw.Db.Close()

	dbw.QueryData()
}

func (dbw *DbWorker) queryDataPre() {
	dbw.User = user{}
}

func (dbw *DbWorker) QueryData() {
	stmt, _ := dbw.Db.Prepare(`select id, username, password, nickname, email, telephone, real_name, enabled, create_time from user`)

	defer stmt.Close()

	dbw.queryDataPre()
	rows, err := stmt.Query()
	defer  rows.Close()
	if err != nil {
		fmt.Printf("insert data error:%v\n", err)
		return
	}

	for rows.Next() {
		err := rows.Scan(&dbw.User.Id, &dbw.User.Username, &dbw.User.Password, &dbw.User.Nickname, &dbw.User.Email, &dbw.User.Telephone, &dbw.User.RealName, &dbw.User.Enabled, &dbw.User.CreateTime)

		if err != nil {
			fmt.Printf(err.Error())
			continue
		}

		if !dbw.User.Nickname.Valid {
			dbw.User.Nickname.String = ""
		}

		fmt.Println("get data: id: ", dbw.User.Id, " username: ", dbw.User.Username, " nickname: ", dbw.User.Nickname, " telephone: ", dbw.User.Telephone, " realName: ", dbw.User.RealName, " enabled: ", dbw.User.Enabled, " createTime: ", dbw.User.CreateTime)
	}

	err = rows.Err()

	if err != nil {
		fmt.Println(err.Error())
	}


}





