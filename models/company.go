package models

import (
	"fmt"

	"github.com/fatih/structs"
	"github.com/go-playground/validator/v10"
)

type CompanyType string

const (
	Corporation        CompanyType = "Corporation"
	NonProfit          CompanyType = "NonProfit"
	Cooperative        CompanyType = "Cooperative"
	SoleProprietorship CompanyType = "Sole Proprietorship"
)

type Company struct {
	ID           string      `json:"id,omitempty" bson:"_id"`
	Name         string      `json:"name" bson:"company_name" binding:"required,min=3,max=15"`
	Description  string      `json:"description,omitempty" bson:"description" binding:"max=3000"`
	EmpCount     uint16      `json:"total_employees" bson:"total_employees" binding:"required"`
	IsRegistered bool        `json:"registered" bson:"registered" binding:"required"`
	Type         CompanyType `json:"type" bson:"type" binding:"required,companyType"`
}

type CompanyUpdate struct {
	Name         *string      `json:"name,omitempty" binding:"omitempty,min=3,max=15"`
	Description  *string      `json:"description,omitempty" binding:"omitempty,max=3000"`
	EmpCount     *uint16      `json:"total_employees,omitempty"`
	IsRegistered *bool        `json:"registered,omitempty"`
	Type         *CompanyType `json:"type,omitempty" binding:"omitempty,companyType"`
}

// Map function returns map values
func (co *Company) Map() map[string]interface{} {
	return structs.Map(co)
}

// Names function returns field names
func (co *Company) Names() []string {
	fields := structs.Fields(co)
	names := make([]string, len(fields))
	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}
	return names
}

func (ct CompanyType) IsValid() bool {
	switch ct {
	case Corporation, NonProfit, Cooperative, SoleProprietorship:
		return true
	}
	return false
}

func validateCompanyType(fl validator.FieldLevel) bool {
	value := fl.Field().Interface()
	switch v := value.(type) {
	case CompanyType:
		switch v {
		case Corporation, NonProfit, Cooperative, SoleProprietorship:
			return true
		}
	case *CompanyType:
		if v != nil {
			switch *v {
			case Corporation, NonProfit, Cooperative, SoleProprietorship:
				return true
			}
		}
	}
	return false
}

func ErrValidationSlice(err error) ([]string, bool) {
	errs := []string{}
	ve, ok := err.(validator.ValidationErrors)
	if ok {
		for _, fe := range ve {
			errStr := errMsgFromValidator(fe.Field(), fe.Tag(), fe.Param())
			errs = append(errs, errStr)
		}
	}

	return errs, ok
}

func errMsgFromValidator(field, tag, value string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s field is required", field)
	case "min":
		return fmt.Sprintf("minimum %s characters required for %s", value, field)
	case "max":
		return fmt.Sprintf("maximum %s characters allowed for %s", value, field)
	case "email":
		return fmt.Sprint("email format is not valid")
	}
	return tag
}
