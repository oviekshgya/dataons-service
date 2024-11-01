package division

import (
	"dataons-service/pkg"
	"time"
)

type Division struct {
	IdDivision   int       `gorm:"column:idDivision;primaryKey;autoIncrement;size:11" json:"id_division"`
	IdDepartment int       `gorm:"column:idDepartment;foreignKey;size:11" json:"idDepartment"`
	NameDivision string    `gorm:"column:nameDivision;size:200" json:"nameDivision"`
	IsActive     int       `gorm:"column:isActive;size:1" json:"isActive"`
	CreatedAt    time.Time `gorm:"column:createdAt;type:datetime;autoCreateTime" json:"createdAt"`
	UpdateAt     time.Time `gorm:"column:updatedAt;type:datetime;autoUpdateTime" json:"updatedAt"`
}

func (Division) TableName() string {
	return pkg.DIVISION
}
