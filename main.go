	package main

import (
	"fmt"
	"goswagger/config"
	"goswagger/controllers"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	dbsqlx := config.ConnectDBSqlx()
	hsqlx := controllers.NewBaseHandlerSqlx(dbsqlx)

	r.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// documentation for developers
	 opts := middleware.SwaggerUIOpts{SpecURL: "/swagger.yaml"}
	 sh := middleware.SwaggerUI(opts, nil)
	 r.Handle("/docs", sh)

	// documentation for share
	// opts1 := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	// sh1 := middleware.Redoc(opts1, nil)
	// r.Handle("/docs", sh1)
	// opts2 := middleware.RapiDocOpts{SpecURL: "/swagger.yaml"}
	// sh2 := middleware.RapiDoc(opts2, nil)
	// r.Handle("/docs", sh2)

	company := r.PathPrefix("/admin/company").Subrouter()
	company.HandleFunc("/", hsqlx.PostCompanySqlx).Methods("POST")
	company.HandleFunc("/", hsqlx.GetCompaniesSqlx).Methods("GET")
	company.HandleFunc("/{id}", hsqlx.EditCompany).Methods("PUT")
	company.HandleFunc("/{id}", hsqlx.DeleteCompany).Methods("DELETE")

	http.Handle("/", r)
	s := &http.Server{
		Addr: fmt.Sprintf("%s:%s", "localhost", "5000"),
	}
	s.ListenAndServe()
}
