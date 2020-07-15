package api

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func isLoggedIn(r *http.Request) bool {
	jwtCookie, err := r.Cookie("jwt")
	if err != nil {
		log.Error(err.Error())
		return false
	}
	log.Debug(jwtCookie)
	return false
}

