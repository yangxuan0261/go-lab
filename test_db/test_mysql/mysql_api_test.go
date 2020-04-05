package test_mysql

import (
	"database/sql"
	"fmt"
	"testing"
)

var dbIns *sql.DB

func init() {
	var err error

	addr := "root:123456@tcp(192.168.2.233:6306)/testdb1?charset=utf8mb4"
	dbIns, err = sql.Open("mysql", addr)
	if err != nil {
		panic(err)
		return
	}
	//defer dbIns.Close()
}

func Test_Insert(t *testing.T) {
	stmt, err := dbIns.Prepare(`INSERT INTO user (name, age) VALUES (?, ?)`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i := 1; i <= 5; i++ {
		ret, err := stmt.Exec(fmt.Sprintf("wilker-%d", i), 100+i)
		if err != nil {
			panic(err)
		}
		if LastInsertId, err := ret.LastInsertId(); nil == err {
			fmt.Println("--- LastInsertId:", LastInsertId)
		}
		if RowsAffected, err := ret.RowsAffected(); nil == err {
			fmt.Println("--- RowsAffected:", RowsAffected)
		}
	}
}

func Test_Delete(t *testing.T) {
	stmt, err := dbIns.Prepare(`DELETE FROM user WHERE id=?`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	ret, err := stmt.Exec(1)
	if err != nil {
		panic(err)
	}

	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("--- RowsAffected:", RowsAffected)
	}
}

func Test_Upate(t *testing.T) {
	stmt, err := dbIns.Prepare(`UPDATE user SET name=? WHERE id<=?`)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec("wolegequ", 2)
	if err != nil {
		panic(err)
	}

	if RowsAffected, err := result.RowsAffected(); nil == err {
		fmt.Println("--- RowsAffected:", RowsAffected)
	}
}

func Test_QuerSingle(t *testing.T) {
	data := userTB{}
	//err := dbIns.QueryRow("SELECT `name`,`age` FROM user WHERE id= ? LIMIT 1", 2).Scan(&data.Name, &data.Age)
	err := dbIns.QueryRow("SELECT * FROM user WHERE id= ? LIMIT 1", 2).Scan(&data.Id, &data.Name, &data.Age) // 映射顺序必须与数据库一致
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- data:%+v\n", data)
}

func Test_QuerMulti(t *testing.T) {
	stmt, err := dbIns.Prepare(`SELECT * From user where age >= ? AND age < ?`) // 映射顺序必须与数据库一致
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(20, 500)
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		data := userTB{}
		rows.Scan(&data.Id, &data.Name, &data.Age)
		if err != nil {
			fmt.Printf("--- rows.Next err:%+v\n", err)
			continue
		}
		fmt.Printf("--- data:%+v\n", data)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}
}

func Test_Transaction(t *testing.T) {
	tx, err := dbIns.Begin()
	if err != nil {
		panic(err)
	}

	sqlStr := `UPDATE user SET name=? WHERE id=?`

	res1, err1 := tx.Exec(sqlStr, "wangwu111", 4)
	if err1 != nil {
		fmt.Printf("--- exec 1 err1:%+v\n", err1)
		err2 := tx.Rollback()
		if err2 != nil {
			fmt.Printf("--- Rollback err2:%+v\n", err2)
			return
		}
		return
	} else {
		line, _ := res1.RowsAffected()
		if line <= 0 {
			fmt.Printf("--- RowsAffected 1 error, line:%+v\n", line)

			err3 := tx.Rollback()
			if err3 != nil {
				fmt.Printf("--- Rollback err3:%+v\n", err3)
				return
			}
			return
		}
	}

	res1, err1 = tx.Exec(sqlStr, "wangwu222", 5)
	if err1 != nil {
		fmt.Printf("--- exec 2 err1:%+v\n", err1)
		err2 := tx.Rollback()
		if err2 != nil {
			fmt.Printf("--- Rollback err2:%+v\n", err2)
			return
		}
		return
	} else {
		line, _ := res1.RowsAffected()
		if line <= 0 {
			fmt.Printf("--- RowsAffected 2 error, line:%+v\n", line)

			err3 := tx.Rollback()
			if err3 != nil {
				fmt.Printf("--- Rollback err3:%+v\n", err3)
			}
			return
		}
	}

	res1, err1 = tx.Exec(sqlStr, "wangwu333", "asdasd") // 传入错误的数据测试回滚
	if err1 != nil {
		fmt.Printf("--- exec 3 err1:%+v\n", err1)
		err2 := tx.Rollback()
		if err2 != nil {
			fmt.Printf("--- Rollback err2:%+v\n", err2)
			return
		}
		return
	} else {
		line, _ := res1.RowsAffected()
		if line <= 0 {
			fmt.Printf("--- RowsAffected 3 error, line:%+v\n", line)

			err3 := tx.Rollback()
			if err3 != nil {
				fmt.Printf("--- Rollback err3:%+v\n", err3)
				return
			}
			return
		}
	}

	err = tx.Commit() // commit 之后才生效
	if err != nil {
		fmt.Printf("--- Commit err:%+v\n", err)

		err3 := tx.Rollback()
		if err3 != nil {
			fmt.Printf("--- Commit Rollback err3:%+v\n", err3)
		}
	}
	fmt.Printf("--- all success\n")
}
