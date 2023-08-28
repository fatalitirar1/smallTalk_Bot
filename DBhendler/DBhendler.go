package dbhendler

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type dbcon interface {
	getLineCon() string
	openConn()
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "smallTalk_bot"
)

func getLineCon() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}

func IsChatExists(chatcode int64) bool {
	db := openConn()
	defer db.Close()

	rows := db.QueryRow("SELECT id AS sqlId,chatcode AS sqlChatcode FROM chats WHERE chatcode = $1 LIMIT 1", chatcode)
	var sqlId, sqlChatcode int64
	switch err := rows.Scan(&sqlId, &sqlChatcode); err {
	case nil:
		return sqlId != 0
	case sql.ErrNoRows:
		return false
	}
	return false
}

func openConn() *sql.DB {
	db, err := sql.Open("postgres", getLineCon())
	if err != nil {
		panic(err)
	}
	return db
}

func CreateUser(chatcode int64) {
	db := openConn()
	defer db.Close()
	_, err := db.Exec("INSERT INTO CHATS(chatcode) VALUES($1)", chatcode)
	if err != nil {
		fmt.Println(err)
	}

}
