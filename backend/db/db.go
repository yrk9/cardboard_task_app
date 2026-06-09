package db

import (
	"database/sql"
	"time"
	"os"
	"github.com/go-sql-driver/mysql"
	"log"
)

func Connect() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	str := mysql.Config{
		DBName: dbName,
		User: dbUser,
		Passwd: dbPassword,
		Addr: dbHost + ":3306",
		Net: "tcp",
		ParseTime: true,
		AllowNativePasswords: true,
	}
	dbConn, err := sql.Open("mysql", str.FormatDSN())

	if err != nil {
		log.Panicf("failed to open database: %v", err)
		return nil, err
	}

	dbConn.SetMaxOpenConns(25)	// 同時に開ける接続の最大数
	dbConn.SetMaxIdleConns(25) // アイドル状態で保持する最大数
	dbConn.SetConnMaxLifetime(5 * time.Minute) // 接続の最大生存時間
	dbConn.SetConnMaxIdleTime(3 * time.Minute) // アイドル状態の接続を保持する最大時間

	if err := dbConn.Ping(); err != nil {
		log.Panicf("failed to connect to mysql: %v", err)
		return nil, err
	}
	log.Println("connected to mysql")
	return dbConn, nil
}