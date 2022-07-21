package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/securecookie"
)

type SessionController struct {
	BaseController
}

type SessionValues struct {
	UserName     string
	Password     string
	LoginType    string
	LoginStatus  int
	LoginId      int
	ProjectID    int
	LoginProject string
}

var cookieToken = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func SetSession(w http.ResponseWriter, sessionValues SessionValues) {
	loginIdstr := strconv.Itoa(sessionValues.LoginId)
	loginStatusStr := strconv.Itoa(sessionValues.LoginStatus)
	projectStr := strconv.Itoa(sessionValues.ProjectID)

	value := make(map[string]string)
	value["project"] = projectStr
	value["loginProject"] = sessionValues.LoginProject
	value["userName"] = sessionValues.UserName
	value["password"] = sessionValues.Password
	value["loginType"] = sessionValues.LoginType
	value["loginStatus"] = loginStatusStr
	value["loginId"] = loginIdstr

	if encoded, err := cookieToken.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		log.Println("Session is Set!", value)
		cookieToken.MaxAge(28800)
	}

}

func ReadSession(w http.ResponseWriter, r *http.Request) SessionValues {
	sessionValues := SessionValues{}
	if cookie, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		if err = cookieToken.Decode("session", cookie.Value, &value); err == nil {
			i, _ := strconv.Atoi(value["loginId"])
			j, _ := strconv.Atoi(value["loginStatus"])
			project, _ := strconv.Atoi(value["project"])

			sessionValues.LoginId = i
			sessionValues.LoginStatus = j
			sessionValues.LoginType = value["loginType"]
			sessionValues.UserName = value["userName"]
			sessionValues.ProjectID = project
			sessionValues.LoginProject = value["loginProject"]

		} else {

			http.Redirect(w, r, "/", 302)
			log.Println("Access Denied! You are not logged in!")
		}
	} else {
		http.Redirect(w, r, "/", 302)
		log.Println("Access Denied! You are not logged in!")
	}
	return sessionValues
}

func ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	log.Println("Logged out Successfully!")

}

func (c *SessionController) ProjectSectionChanges() {
	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	storedSession := ReadSession(w, r)
	if r.Method == "POST" {
		projectSelection := c.GetString("projectSelection")
		projectSelectionName := c.GetString("projectSelectionName")
		project, _ := strconv.Atoi(projectSelection)

		sessionValues := SessionValues{}
		sessionValues.LoginStatus = 1
		sessionValues.LoginType = storedSession.LoginType
		sessionValues.UserName = storedSession.UserName
		sessionValues.LoginId = storedSession.LoginId
		sessionValues.ProjectID = project
		sessionValues.LoginProject = projectSelectionName

		SetSession(w, sessionValues)
		if sessionValues.LoginType == "service" || sessionValues.LoginType == "paycraft" {
			slices := []interface{}{"service"}
			sliceToClient, _ := json.Marshal(slices)
			w.Write([]byte(sliceToClient))
		} else if sessionValues.LoginType == "association" {
			slices := []interface{}{"association"}
			sliceToClient, _ := json.Marshal(slices)
			w.Write([]byte(sliceToClient))
		} else {
			slices := []interface{}{"true"}
			sliceToClient, _ := json.Marshal(slices)
			w.Write([]byte(sliceToClient))
		}

	}
}
