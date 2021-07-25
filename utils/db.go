package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMySQL() error {
	if db == nil {
		db0, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/productgo")
		if err != nil {
			return err
		}
		db = db0
		err = InitTableUser()
		if err != nil {
			return err
		}
	}
	return nil
}

func InitTableUser() error {
	sql := `CREATE TABLE IF NOT EXISTS user(
		id INT(10) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		createtime int(10));
        `
	if _, err := ExecSQL(sql); err != nil {
		return err
	}
	return nil
}

func ExecSQL(sql string, args ...interface{}) (int64, error) {
	if result, err := db.Exec(sql, args...); err != nil {
		return 0, err
	} else {
		cout, err := result.RowsAffected()
		if err != nil {
			return 0, err
		} else {
			return cout, nil
		}
	}
}

func QueryRow(sql string) *sql.Row {
	return db.QueryRow(sql)
}
