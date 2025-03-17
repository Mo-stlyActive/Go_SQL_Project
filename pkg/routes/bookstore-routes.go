package routes

import (
	"github.com/gorilla/mux"
	//use absolute paths rather the relative paths like you would in node
	"github.com/Mo-stlyActive/Go_SQL_Project/pkg/controllers"
)

var RegisterBookStore = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
}
