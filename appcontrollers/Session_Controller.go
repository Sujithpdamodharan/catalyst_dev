package appcontrollers

import (
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
)

type Session_Controller struct {
	BaseController
}
type SessionValues struct {
	LoginId string
}

var cookieToken = securecookie.New(securecookie.GenerateRandomKey(64), securecookie.GenerateRandomKey(32))

func SetSession(w http.ResponseWriter, sessionValues SessionValues) {
	value := make(map[string]string)
	value["loginId"] = sessionValues.LoginId
	if encoded, err := cookieToken.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "http://onediwebapp.technovialabs.com/",
		}
		http.SetCookie(w, cookie)
		log.Println("Session is Set!")
		cookieToken.MaxAge(1800)
	}
}
func ReadSession(w http.ResponseWriter, r *http.Request) SessionValues {
	sessionValues := SessionValues{}
	if cookie, err := r.Cookie("session"); err == nil {
		log.Println(err)
		value := make(map[string]string)
		if err = cookieToken.Decode("session", cookie.Value, &value); err == nil {
			sessionValues.LoginId = value["loginId"]
		}
	}
	return sessionValues
}
func ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "http://onediwebapp.technovialabs.com/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	log.Println("Logged out Successfully!")

}
