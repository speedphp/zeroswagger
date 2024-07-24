package zeroswagger

import (
	"net/http"
)

func Middleware(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		next(w, r)
	}
}
