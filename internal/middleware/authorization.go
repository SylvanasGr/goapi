package middleware

import (
	"errors"
	"net/http"

	"github.com/SylvanasGr/goapi/api"
	"github.com/SylvanasGr/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var ErrUnauthorized = errors.New("Invalid username or Token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")

		if username == "" || token == "" {
			log.Error(ErrUnauthorized)
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		database, err := tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		loginDetails := database.GetUserLoginDetails(username)

		if loginDetails == nil || token != loginDetails.AuthToken {
			log.Error(ErrUnauthorized)
			api.RequestErrorHandler(w, ErrUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
