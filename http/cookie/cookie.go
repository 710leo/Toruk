package cookie

import (
	"net/http"

	"github.com/gorilla/securecookie"
)

var SecureCookie *securecookie.SecureCookie

func Init() {
	var hashKey = []byte(nil)
	var blockKey = []byte(nil)
	SecureCookie = securecookie.New(hashKey, blockKey)
}

func ReadUser(r *http.Request) (example string) {
	if cookie, err := r.Cookie("u"); err == nil {
		value := make(map[string]interface{})
		if err = SecureCookie.Decode("u", cookie.Value, &value); err == nil {
			example = value["example"].(string)
		}
	}
	return
}

func WriteUser(w http.ResponseWriter, example string) error {
	value := make(map[string]interface{})
	value["example"] = example
	encoded, err := SecureCookie.Encode("u", value)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     "u",
		Value:    encoded,
		Path:     "/",
		MaxAge:   3600 * 24 * 7,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)

	return nil
}

func RemoveUser(w http.ResponseWriter) error {
	value := make(map[string]interface{})
	value["example"] = ""
	encoded, err := SecureCookie.Encode("u", value)
	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:   "u",
		Value:  encoded,
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	return nil
}
