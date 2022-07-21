package controllers

import (
	"app/CatalystAdmin/models"
	"app/CatalystAdmin/viewmodels"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type UserTokenController struct {
	BaseController
}

func (c *UserTokenController) AddUserTokensToDb() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r)
	if storedSession.UserName != "" {
		if r.Method == "POST" {
			log.Println("cp2")
			userTokenData := models.UserToken{}
			userTokenData.UserName = c.GetString("addUserName")
			userTokenData.UserMobileNumber = c.GetString("addMobile")
			userTokenData.TokenStatus = "1"
			expiryDateTime := time.Now().AddDate(0, 0, 7)
			userTokenData.TokenExpiryDate = expiryDateTime.Format("2006-01-02 15:04:05")
			currentDate := time.Now().Local()
			userTokenData.CreatedAt = currentDate.Format("2006-01-02 15:04:05")
			userTokenData.CreatedBy = storedSession.LoginId

			rand.Seed(time.Now().UnixNano())
			charset := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNPQRSTUVWXYZ")

			random := make([]rune, 10)
			for i := range random {
				random[i] = charset[rand.Intn(len(charset))]
			}
			userTokenData.UserAuthToken = string(random)
			dbStatus := userTokenData.AddUserTokenToDb(storedSession.LoginType)
			if dbStatus == "true" {
				slices := []interface{}{dbStatus, userTokenData.UserAuthToken}
				sliceToClient, _ := json.Marshal(slices)
				w.Write([]byte(sliceToClient))

			} else {
				w.Write([]byte("false"))
			}
		} else {
			viewmodel := viewmodels.UserViewModel{}
			viewmodel.LoginType = storedSession.LoginType
			viewmodel.UserName = storedSession.UserName
			log.Println("viewmodel", viewmodel)
			c.Data["vm"] = viewmodel
			if storedSession.LoginType == "1" {
				c.Layout = "layout/layout.html"
				c.TplName = "template/add_user_token.html"
			} else {
				c.Layout = "layout/layout.html"
				c.TplName = "template/error403.html"
			}
		}
	} else {
		http.Redirect(w, r, "/", 302)
	}
}

func (c *UserTokenController) ViewUserTokens() {
	log.Println("iam here")
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r)
	viewModel := viewmodels.UserDetailsViewModel{}
	dbStatus, userTokenData := models.GetAllUserTokenDetails(storedSession.LoginType)
	if dbStatus == "true" {
		for i := 0; i < len(userTokenData); i++ {
			var tempValueSlice []string
			//str := strconv.Itoa(fareData[i].Status)
			tempValueSlice = append(tempValueSlice, userTokenData[i].UserName)
			tempValueSlice = append(tempValueSlice, userTokenData[i].UserMobileNumber)
			tempValueSlice = append(tempValueSlice, userTokenData[i].UserAuthToken)
			tempValueSlice = append(tempValueSlice, userTokenData[i].TokenExpiryDate)
			if userTokenData[i].TokenStatus == "1" {
				tempValueSlice = append(tempValueSlice, "Active")
			} else {
				tempValueSlice = append(tempValueSlice, "Inactive")
			}
			viewModel.Values = append(viewModel.Values, tempValueSlice)
			keyStr := strconv.Itoa(userTokenData[i].UserTokenID)
			viewModel.Keys = append(viewModel.Keys, keyStr)

		}
	} else {
		log.Println("error")
	}
	viewModel.LoginType = storedSession.LoginType
	viewModel.UserName = storedSession.UserName

	c.Data["vm"] = viewModel
	if viewModel.LoginType == "1" {
		c.Layout = "layout/layout.html"
		c.TplName = "template/view_user_token.html"
	} else {
		c.Layout = "layout/layout.html"
		c.TplName = "template/error403.html"
	}
}

func (c *UserTokenController) RecallUserTokens() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r)
	log.Println("stored setion", storedSession)
	userTokenID := c.Ctx.Input.Param(":userTokenID")
	dbStatus := models.RecallUserTokenID(userTokenID)
	if dbStatus == "true" {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}

}
