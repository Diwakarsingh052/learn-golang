package middleware

import (
	"auth/model"
	"fmt"
	"net/http"
)

type RequireUser struct {
	model.UserServiceInterface
}

func (mw *RequireUser) ApplyFn(next http.HandlerFunc) http.HandlerFunc { // it can accept anything which has rw , req in signature

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil { // cookie is not there
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		user, err := mw.SearchToken(cookie.Value)

		if err != nil { // cookie or token is not valid
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		fmt.Println(user)
		next(w, r)

	})
}
