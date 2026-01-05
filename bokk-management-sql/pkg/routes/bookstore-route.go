package route

import "github.com/gorilla/mux"
import "bokk-management-sql/pkg/controllers"

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBok).Methods("DELETE")
}
