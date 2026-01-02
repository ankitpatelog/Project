package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		start := time.Now()
		next.ServeHTTP(w,r)
		fmt.Println(r.Method,r.URL.Path,time.Since(start))
	})
}

func Recover(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := recover()
		if err!=nil {
			http.Error(w,"Internal sever error",http.StatusInternalServerError)
			return
		}
	})
}