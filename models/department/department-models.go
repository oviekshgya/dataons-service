package department

import (
	"dataons-service/pkg"
	"time"
)

type Department struct {
	IdDepartment   int       `gorm:"column:idDepartment;primaryKey;autoIncrement;size:11" json:"idDepartment"`
	IdCompany      int       `gorm:"column:idCompany;foreignKey;size:11" json:"idCompany"`
	NameDepartment string    `gorm:"column:nameDepartment;size:200" json:"nameDepartment"`
	IsActive       int       `gorm:"column:isActive;size:1" json:"isActive"`
	CreatedAt      time.Time `gorm:"column:createdAt;type:datetime;autoCreateTime" json:"createdAt"`
	UpdateAt       time.Time `gorm:"column:updatedAt;type:datetime;autoUpdateTime" json:"updatedAt"`
}

func (Department) TableName() string {
	return pkg.DEPARTEMEN
}
