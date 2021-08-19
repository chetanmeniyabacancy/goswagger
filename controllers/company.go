package controllers

import (
	"encoding/json"
	"fmt"
	"goswagger/lang"
	"goswagger/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

// BaseHandler will hold everything that controller needs
type BaseHandlerSqlx struct {
	db *sqlx.DB
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandlerSqlx(db *sqlx.DB) *BaseHandlerSqlx {
	return &BaseHandlerSqlx{
		db: db,
	}
}

// swagger:model CommonError
type CommonError struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
}

// swagger:model CommonSuccess
type CommonSuccess struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
}

// swagger:model GetCompanies
type GetCompanies struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string            `json:"message"`
	Data    *models.Companies `json:"data"`
}

// swagger:model GetCompany
type GetCompany struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string `json:"message"`
	// Companies for this user
	Data *models.Company `json:"data"`
}

// ErrHandler returns error message response
func ErrHandler(errmessage string) *CommonError {
	errresponse := CommonError{}
	errresponse.Status = 0
	errresponse.Message = errmessage
	return &errresponse
}

// swagger:route GET /admin/company/ admin listCompany
// Get companies list
//
// security:
// - apiKey: []
// responses:
//  401: CommonError
//  200: GetCompanies
func (h *BaseHandlerSqlx) GetCompaniesSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetCompanies{}
	companies := models.GetCompaniesSqlx(h.db)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = companies

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route POST /admin/company/ admin addCompany
// Create a new company
//
// security:
// - apiKey: []
// responses:
//  401: CommonError
//  200: GetCompany
func (h *BaseHandlerSqlx) PostCompanySqlx(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	response := GetCompany{}

	decoder := json.NewDecoder(r.Body)
	var reqcompany *models.ReqCompany
	err := decoder.Decode(&reqcompany)
	fmt.Println(err)

	if err != nil {
		json.NewEncoder(w).Encode(ErrHandler(lang.Get("invalid_requuest")))
		return
	}

	company, errmessage := models.PostCompanySqlx(h.db, reqcompany)
	if errmessage != "" {
		json.NewEncoder(w).Encode(ErrHandler(errmessage))
		return
	}

	response.Status = 1
	response.Message = lang.Get("insert_success")
	response.Data = company
	json.NewEncoder(w).Encode(response)
}

// swagger:route  PUT /admin/company/{id} admin editCompany
// Edit a company
//
// consumes:
//         - application/x-www-form-urlencoded
// security:
// - apiKey: []
// responses:
//  401: CommonError
//  200: GetCompany
func (h *BaseHandlerSqlx) EditCompany(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Header().Set("content-type", "application/json")
	vars := mux.Vars(r)
	response := GetCompany{}
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		json.NewEncoder(w).Encode(ErrHandler(lang.Get("invalid_requuest")))
		return
	}

	var reqcompany models.ReqCompany
	reqcompany.Status, err = strconv.ParseInt(r.FormValue("status"), 10, 64)
	reqcompany.Name = r.FormValue("name")

	if err != nil {
		json.NewEncoder(w).Encode(ErrHandler(lang.Get("invalid_requuest")))
		return
	}

	company, errmessage := models.EditCompany(h.db, &reqcompany, id)
	if errmessage != "" {
		json.NewEncoder(w).Encode(ErrHandler(errmessage))
		return
	}

	response.Status = 1
	response.Message = lang.Get("update_success")
	response.Data = company
	json.NewEncoder(w).Encode(response)
}

// swagger:route DELETE /admin/company/{id} admin deleteCompany
// Delete company
//
// security:
// - apiKey: []
// responses:
//  401: CommonError
//  200: CommonSuccess
// Create handles Delete get company
func (h *BaseHandlerSqlx) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	errmessage := models.DeleteCompany(h.db, vars["id"])

	if errmessage != "" {
		json.NewEncoder(w).Encode(ErrHandler(errmessage))
		return
	}

	successresponse := CommonSuccess{}
	successresponse.Status = 1
	successresponse.Message = lang.Get("delete_success")

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(successresponse)
}
