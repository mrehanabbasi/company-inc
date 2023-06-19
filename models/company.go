package models

import "github.com/fatih/structs"

type CompanyType string

const (
	Corporation        CompanyType = "Corporation"
	NonProfit          CompanyType = "NonProfit"
	Cooperative        CompanyType = "Cooperative"
	SoleProprietorship CompanyType = "Sole Proprietorship"
)

type Company struct {
	ID           string      `json:"id" bson:"_id"`
	Name         string      `json:"name" bson:"company_name" binding:"required,min:3,max:15"`
	Description  string      `json:"description,omitempty" bson:"description" binding:"max=3000"`
	EmpCount     *uint16     `json:"total_employees" bson:"total_employees" binding:"required"`
	IsRegistered *bool       `json:"registered" bson:"registered" binding:"required"`
	Type         CompanyType `json:"type" bson:"type" binding:"required"`
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
