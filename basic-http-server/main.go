package main
import "net/http"
import "basic-http-server/handlers"
import "basic-http-server/middleware"

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/time",handlers.handleTime)
	mux.HandleFunc("/status",handlers.handlestatus)
	mux.HandleFunc("/health",handlers.handlehealth)

	loggedmux := middleware.Logger(mux)
	http.ListenAndServe(":8080",loggedmux)
}