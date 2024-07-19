package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Oriseer/goapi/api"
	"github.com/Oriseer/internal/tools"
	log "github.com/sirupsen/logrus"
)

var unAuthorizedErrorMsg = errors.New(fmt.Sprintf("Invalid Username"))

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")

		if username != "" || token != "" {
			api.RequstErrorHandler(w, unAuthorizedErrorMsg)
		}

		var database *tools.DatabaseInterface
		database, err := tools.NewDatabase()
		if err != nil {
			log.Error(err)
			api.InternalErrorHandler(w)
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || token != *(loginDetails).Auth_token {
			log.Error(unAuthorizedErrorMsg)
			api.RequstErrorHandler(w, unAuthorizedErrorMsg)
		}

		next.ServeHTTP(w, r)

	})
}
