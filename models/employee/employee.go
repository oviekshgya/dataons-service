package employee

import (
	"dataons-service/pkg"
	"time"
)

type Employee struct {
	IdEmployee   int       `gorm:"column:idEmployee;primaryKey;autoIncrement;size:11" json:"idDepartment"`
	IdDivision   int       `gorm:"column:idDivision;foreignKey;size:11" json:"id_division"`
	NameEmployee string    `gorm:"column:nameEmployee;size:200" json:"nameEmployee"`
	NPK          string    `gorm:"column:npK;size:20" json:"nPK"`
	IsActive     int       `gorm:"column:isActive;size:1" json:"isActive"`
	CreatedAt    time.Time `gorm:"column:createdAt;type:datetime;autoCreateTime" json:"createdAt"`
	UpdateAt     time.Time `gorm:"column:updatedAt;type:datetime;autoUpdateTime" json:"updatedAt"`
}

func (Employee) TableName() string {
	return pkg.Employee
}
