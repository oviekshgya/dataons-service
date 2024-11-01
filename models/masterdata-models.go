package models

type MasterData struct {
	NameCompany    string `gorm:"column:nameCompany" json:"name_company"`
	NameDepartment string `gorm:"column:nameDepartment" json:"name_department"`
	NameDivision   string `gorm:"column:nameDivision" json:"name_division"`
}
