package company

import (
	"dataons-service/pkg"
	"time"
)

type Company struct {
	IdCompany   int       `gorm:"column:idCompany;primaryKey;autoIncrement;size:11" json:"idCompany"`
	NameCompany string    `gorm:"column:nameCompany;size:200" json:"nameCompany"`
	IsActive    int       `gorm:"column:isActive;size:1" json:"isActive"`
	Address     string    `gorm:"column:address;size:200" json:"address"`
	CreatedAt   time.Time `gorm:"column:createdAt;type:datetime;autoCreateTime" json:"createdAt"`
	UpdateAt    time.Time `gorm:"column:updatedAt;type:datetime;autoUpdateTime" json:"updatedAt"`
}

func (Company) TableName() string {
	return pkg.COMPANY
}

type MasterDepartment struct {
	IdDepartment   int       `gorm:"column:idDepartment;primaryKey;autoIncrement;size:11" json:"idDepartment"`
	IdCompany      int       `gorm:"column:idCompany;foreignKey;size:11" json:"idCompany"`
	NameDepartment string    `gorm:"column:nameDepartment;size:200" json:"nameDepartment"`
	IsActive       int       `gorm:"column:isActive;size:1" json:"isActive"`
	CreatedAt      time.Time `gorm:"column:createdAt;type:datetime;autoCreateTime" json:"createdAt"`
	UpdateAt       time.Time `gorm:"column:updatedAt;type:datetime;autoUpdateTime" json:"updatedAt"`
}

func (MasterDepartment) TableName() string {
	return pkg.DEPARTEMEN
}

type MasterDivision struct {
	IdDivision   int       `gorm:"column:idDivision;primaryKey;autoIncrement;size:11" json:"id_division"`
	IdDepartment int       `gorm:"column:idDepartment;foreignKey;size:11" json:"idDepartment"`
	NameDivision string    `gorm:"column:nameDivision;size:200" json:"nameDivision"`
	IsActive     int       `gorm:"column:isActive;size:1" json:"isActive"`
	CreatedAt    time.Time `gorm:"column:createdAt;type:datetime;autoCreateTime" json:"createdAt"`
	UpdateAt     time.Time `gorm:"column:updatedAt;type:datetime;autoUpdateTime" json:"updatedAt"`
}

func (MasterDivision) TableName() string {
	return pkg.DIVISION
}

type MasterCompany struct {
	IdCompany   int                `gorm:"column:idCompany;primaryKey;autoIncrement;size:11" json:"idCompany"`
	NameCompany string             `gorm:"column:nameCompany;size:200" json:"nameCompany"`
	IsActive    int                `gorm:"column:isActive;size:1" json:"isActive"`
	Address     string             `gorm:"column:address;size:200" json:"address"`
	Department  []MasterDepartment `gorm:"foreignKey:idCompany"`
}

func (MasterCompany) TableName() string {
	return pkg.COMPANY
}
