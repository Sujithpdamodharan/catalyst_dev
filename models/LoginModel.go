package models

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"log"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

/* function for db connection*/
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	// local server
	dbUser := beego.AppConfig.String("DbUser")
	dbPass := beego.AppConfig.String("DbPasword")
	dbName := beego.AppConfig.String("DbName")

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

/* fuction for check login details*/
func CheckLogin(username string, password string) (string, string, int, string, int) {
	hasher := md5.New()
	hasher.Write([]byte(password))
	encryptPassword := hex.EncodeToString(hasher.Sum(nil))
	db := dbConn()
	count := 0

	selDB, err := db.Query("SELECT  login_id, username, status, login_type FROM login  WHERE username = ? AND password = ?", username, encryptPassword)
	if err != nil {
		panic(err.Error())
	}
	var usernameFromDb, loginType string
	var loginID, loginStatus int

	for selDB.Next() {
		err = selDB.Scan(&loginID, &usernameFromDb, &loginStatus, &loginType)
		log.Println("error", err)

		if loginStatus == 1 {
			count = count + 1
		}
	}
	selDB.Close()
	db.Close()
	if count == 1 {
		return "true", usernameFromDb, loginStatus, loginType, loginID
	} else {

		return "false", usernameFromDb, loginStatus, loginType, loginID
	}

}
