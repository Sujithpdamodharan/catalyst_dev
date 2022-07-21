package controllers

import (
	"net/http"
)

type DashBoardController struct {
	BaseController
}

func (c *DashBoardController) DashboardView() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r)
	if storedSession.UserName != "" {
		if storedSession.LoginType == "1" {
			c.Layout = "layout/layout.html"
			c.TplName = "template/dashboard.html"
		}
	} else {
		http.Redirect(w, r, "/", 302)
	}
}
