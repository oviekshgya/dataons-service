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
