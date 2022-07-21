package controllers

import (
	"app/CatalystAdmin/models"
	"encoding/json"
	"log"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Login() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	if r.Method == "POST" {
		userName := c.GetString("username")
		password := c.GetString("password")
		dbStatus, actualUsername, LoginStatus, loginType, LoginID := models.CheckLogin(userName, password)
		switch dbStatus {
		case "true":
			sessionValues := SessionValues{}
			sessionValues.LoginStatus = LoginStatus
			sessionValues.LoginType = loginType
			sessionValues.UserName = actualUsername
			sessionValues.LoginId = LoginID

			log.Println("sessionValues", sessionValues)
			SetSession(w, sessionValues)

			if sessionValues.LoginType == "1" {
				slices := []interface{}{"true"}
				sliceToClient, _ := json.Marshal(slices)
				w.Write([]byte(sliceToClient))
			}
		case "false":
			slices := []interface{}{"false"}
			sliceToClient, _ := json.Marshal(slices)
			w.Write([]byte(sliceToClient))
		}
	} else {
		c.TplName = "template/main_login.html"
	}
}

func (c *LoginController) LogOut() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	ClearSession(w)
	http.Redirect(w, r, "/", 302)
}
