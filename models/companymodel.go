package models

import (
	"goswagger/lang"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

// swagger:model Company
type Company struct {
	// Id of the company
	// in: int64
	Id int64 `json:"id"`
	// Name of the company
	// in: string
	Name string `json:"name"`
	// Status of the company
	// in: int64
	Status int64 `json:"status"`
}

type Companies []Company

// swagger:parameters admin deleteCompany
type ReqDeleteCompany struct {
	// name: id
	// in: path
	// type: integer
	// required: true
	Id string `json:"id"`
}

type ReqAddCompany struct {
	// Name of the company
	// in: string
	Name string `json:"name" validate:"required,min=2,max=100,alpha_space"`
	// Status of the company
	// in: int64
	Status int64 `json:"status" validate:"required"`
}

// swagger:parameters admin editCompany
type ReqCompany struct {
	// name: id
	// in: path
	// type: integer
	// required: true
	Id int64 `json:"id"`
	// name: Name
	// in: formData
	// type: string
	// required: true
	Name string `json:"name" validate:"required,min=2,max=100,alpha_space"`
	// name: Status
	// in: formData
	// type: int64
	// required: true
	Status int64 `json:"status" validate:"required"`
}

// swagger:parameters admin addCompany
type ReqCompanyBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/ReqAddCompany"
	//  required: true
	Body ReqAddCompany `json:"body"`
}

// ErrHandler returns error message bassed on env debug
func ErrHandler(err error) string {
	var errmessage string
	if os.Getenv("DEBUG") == "true" {
		errmessage = err.Error()
	} else {
		errmessage = lang.Get("something_went_wrong")
	}
	return errmessage
}

// GetCompaniesSqlx get company
func GetCompaniesSqlx(db *sqlx.DB) *Companies {
	companies := Companies{}

	err := db.Select(&companies, "SELECT id,name,status FROM companies")

	if err != nil {
		log.Fatal(err)
	}
	return &companies
}

// PostCompanySqlx insert company
func PostCompanySqlx(db *sqlx.DB, reqcompany *ReqCompany) (*Company, string) {
	name := reqcompany.Name
	status := reqcompany.Status

	var company Company

	stmt, err := db.Prepare("INSERT INTO companies(name,status) VALUES(?,?)")
	if err != nil {
		return &company, ErrHandler(err)
	}
	result, err := stmt.Exec(name, status)
	if err != nil {
		return &company, ErrHandler(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return &company, ErrHandler(err)
	}

	err = db.Get(&company, "SELECT id,name,status FROM companies where id = ?", id)
	if err != nil {
		return &company, lang.Get("no_result")
	}
	return &company, ""
}

// PostCompanySqlx insert company
func EditCompany(db *sqlx.DB, reqcompany *ReqCompany, id int64) (*Company, string) {
	name := reqcompany.Name
	status := reqcompany.Status

	var company Company

	stmt, err := db.Prepare("Update companies set name=?, status=? where id = ?")
	if err != nil {
		return &company, ErrHandler(err)
	}
	_, err = stmt.Exec(name, status, id)
	if err != nil {
		return &company, ErrHandler(err)
	}

	err = db.Get(&company, "SELECT id,name,status FROM companies where id = ?", id)
	if err != nil {
		return &company, lang.Get("no_result")
	}
	return &company, ""
}

// DeleteCompany get company
func DeleteCompany(db *sqlx.DB, id string) string {
	stmt, err := db.Prepare("DELETE FROM companies where id = ?")
	if err != nil {
		return ErrHandler(err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return ErrHandler(err)
	}
	return ""
}
