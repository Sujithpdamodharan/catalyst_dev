// CreatedBy Sujith P D

package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type UserToken struct {
	UserTokenID      int
	UserName         string
	UserMobileNumber string
	UserAuthToken    string
	TokenStatus      string
	TokenExpiryDate  string
	CreatedBy        int
	CreatedAt        string
	UpdatedBy        int
	UpdatedAt        string
}

func (m *UserToken) AddUserTokenToDb(loginType string) string {
	db := dbConn()
	tx, _ := db.Begin()
	stmt, err := tx.Prepare("INSERT INTO user_token (user_name, user_mobile_number, user_auth_token, token_status, token_expiry_date, created_at, created_by) VALUES(?,?,?,?,?,?,?)")
	log.Println("errorrr111", err)
	_, err = stmt.Exec(m.UserName, m.UserMobileNumber, m.UserAuthToken, m.TokenStatus, m.TokenExpiryDate, m.CreatedAt, m.CreatedBy)
	if err != nil {
		log.Println("login error", err)
		return "false"
	}
	stmt.Close()
	tx.Commit()
	db.Close()
	return "true"
}

func GetAllUserTokenDetails(loginType string) (string, []UserToken) {
	db := dbConn()
	userTokenStruct := []UserToken{}
	if loginType == "1" {
		userTokenDataById, _ := db.Query("SELECT  user_token_id, user_name, user_mobile_number, user_auth_token, token_status, token_expiry_date FROM user_token ORDER BY user_token_id DESC")
		for userTokenDataById.Next() {
			var userTokenStructData UserToken
			err := userTokenDataById.Scan(&userTokenStructData.UserTokenID, &userTokenStructData.UserName, &userTokenStructData.UserMobileNumber, &userTokenStructData.UserAuthToken, &userTokenStructData.TokenStatus, &userTokenStructData.TokenExpiryDate)
			log.Println("err", err)
			userTokenStruct = append(userTokenStruct, userTokenStructData)

			log.Println("iam here11", userTokenStruct)
		}
		userTokenDataById.Close()
	}
	db.Close()
	return "true", userTokenStruct
}

func RecallUserTokenID(userTokenID string) string {
	db := dbConn()
	stmt, _ := db.Prepare("update user_token set token_status = ? WHERE user_token_id = ?")
	res, err := stmt.Exec(2, userTokenID)
	log.Println("err2", err)
	affect, _ := res.RowsAffected()
	stmt.Close()
	log.Println("affect", affect)
	db.Close()
	return "true"
}
