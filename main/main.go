// Author @Sujith P D
package main

import (
	"app/CatalystAdmin/controllers"

	"app/CatalystAdmin/appcontrollers"

	"github.com/astaxie/beego"
)

//func main() {
func main() {
	//Admin Panel
	beego.Router("/", &controllers.LoginController{}, "*:Login")
	beego.Router("/logout", &controllers.LoginController{}, "*:LogOut")
	beego.Router("/dashboard", &controllers.DashBoardController{}, "*:DashboardView")
	beego.Router("/userToken", &controllers.UserTokenController{}, "*:ViewUserTokens")
	beego.Router("/userToken/add", &controllers.UserTokenController{}, "*:AddUserTokensToDb")
	beego.Router("/userToken/recall/:userTokenID", &controllers.UserTokenController{}, "*:RecallUserTokens")

	//for application
	beego.Router("/checkLogin/:data", &appcontrollers.Login_Controller{}, "get:CheckLogin")

	beego.Run()
}
