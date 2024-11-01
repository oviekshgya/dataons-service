package models

type JSONCreateUpdate struct {
	IdCompany      int    `json:"idComapny"`
	NameCompany    string `json:"nameCompany"`
	IdDepartment   int    `json:"idDepartment"`
	NameDepartment string `json:"nameDepartment"`
	IdDivison      int    `json:"idDivison"`
	NameDivison    string `json:"nameDivison"`
	IdEmployee     int    `json:"idEmployee"`
	NameEmployee   string `json:"nameEmployee"`
	Npk            string `json:"npk"`
	IsActive       int    `json:"isActive"`
	Address        string `json:"address"`
}
