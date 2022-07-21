package appcontrollers

import (
	models "app/CatalystAdmin/appmodels"
	"encoding/json"
	"log"
	"time"

	"github.com/nubo/jwt"
)

type Login_Controller struct {
	BaseController
}

/*
CheckLogin
Input mobile number and password and type mac address
if type 301 checks in login table and if verified update msg
status to active.
if type 303 usual login verification
200 success 201 error
After successful login insert mac address along with
login_id

*/
func (c *Login_Controller) CheckLogin() {

	w := c.Ctx.ResponseWriter
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var data jwt.ClaimSet
	var code int32

	userTokenID := c.Ctx.Input.Param(":data")
	log.Println("Values", userTokenID)
	currentDate := time.Now().Local()
	ExpiryDate := currentDate.Format("2006-01-02 15:04:05")

	status, logindetails := models.ValidateLogin(userTokenID, ExpiryDate)
	log.Println("logindetails", logindetails)
	if status == "true" {
		code = 200
		message := "Login Successful"
		data = jwt.ClaimSet{
			"code":    code,
			"message": message,
			"data":    logindetails,
		}
	} else {
		code = 201
		errData := jwt.ClaimSet{
			"userTokenID": userTokenID,
		}
		data = jwt.ClaimSet{
			"code":    code,
			"message": "Retry",
			"data":    errData,
		}
	}
	slice := []interface{}{data}
	sliceToClient, _ := json.Marshal(slice[0])
	w.Write(sliceToClient)
}
