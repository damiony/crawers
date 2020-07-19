package data

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMysql(username, password, host string) (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/zhenaiwang?charset=utf8", username, password, host)
	DB, err = sql.Open("mysql", dsn)
	return DB, err
}
