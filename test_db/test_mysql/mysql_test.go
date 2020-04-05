package test_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

type DbWorker struct {
	Dsn      string
	Db       *sql.DB
	UserInfo userTB
}
type userTB struct {
	Id   int
	Name string
	Age  uint64
}

func (dbw *DbWorker) insertData() {
	stmt, err := dbw.Db.Prepare(`INSERT INTO user (name, age) VALUES (?, ?)`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	ret, err := stmt.Exec("xys", 25)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}

func (dbw *DbWorker) QueryDataPre() {
	dbw.UserInfo = userTB{}
}

func (dbw *DbWorker) queryData() {
	stmt, err := dbw.Db.Prepare(`SELECT * From user where age >= ? AND age < ?`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	dbw.QueryDataPre()

	rows, err := stmt.Query(20, 30)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&dbw.UserInfo.Id, &dbw.UserInfo.Name, &dbw.UserInfo.Age)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		fmt.Printf("--- data:%+v\n", dbw.UserInfo)
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func Test_001(t *testing.T) {
	var err error
	dbw := DbWorker{
		Dsn: "root:123456@tcp(127.0.0.1:6306)/testdb1?charset=utf8mb4",
	}
	dbw.Db, err = sql.Open("mysql", dbw.Dsn)
	if err != nil {
		panic(err)
	}
	defer dbw.Db.Close()

	dbw.insertData()
	dbw.queryData()
}

/*
// 1. 创建 testdb1 数据库
create database testdb1;

// 2. 创建 user 表
CREATE TABLE IF NOT EXISTS `user`(
   `id` INT UNSIGNED AUTO_INCREMENT,
   `name` VARCHAR(100) NOT NULL,
   `age` INT UNSIGNED NOT NULL,
   PRIMARY KEY ( `id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

/*
使 唯一 id 从某个值开始自增
1. insert 一个
*/
