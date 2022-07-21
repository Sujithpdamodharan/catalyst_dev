package models

import (
	"database/sql"
	"log"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := beego.AppConfig.String("DbUser")
	dbPass := beego.AppConfig.String("DbPasword")
	dbName := beego.AppConfig.String("DbName")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func ValidateLogin(userTokenID string, ExpiryDate string) (string, string) {

	var UserAuthToken string
	log.Println("here")
	count := 0
	db := dbConn()
	values, err := db.Query("SELECT user_auth_token from user_token where token_status=? and user_auth_token=? and token_expiry_date > ?", 1, userTokenID, ExpiryDate)
	log.Println(err)
	for values.Next() {
		count++
		err = values.Scan(&UserAuthToken)
		log.Println(err)
	}
	log.Println("count", count)
	db.Close()
	if count > 0 {
		return "true", UserAuthToken
	} else {
		return "false", UserAuthToken
	}
}
