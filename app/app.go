package app

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/vipul-08/student-api/domain"
	"github.com/vipul-08/student-api/service"
	"log"
	"net/http"
	"os"
)

func StartRoutes() {
	router := mux.NewRouter()

	handler := StudentHandlers{service.NewStudentService(domain.NewStudentRepositoryDb())}

	router.HandleFunc("/students", handler.getAllStudents).Methods(http.MethodGet)
	router.HandleFunc("/students/{id:[0-9]+}", handler.getStudentById).Methods(http.MethodGet)
	router.HandleFunc("/students", handler.insertStudent).Methods(http.MethodPost)
	router.HandleFunc("/students/{id:[0-9]+}", handler.deleteStudent).Methods(http.MethodDelete)
	router.HandleFunc("/students", handler.updateStudent).Methods(http.MethodPut)

	ops := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(ops, nil)
	router.Handle("/docs", sh)
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
