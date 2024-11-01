package queryScopes

import (
	"dataons-service/pkg"
	"gorm.io/gorm"
)

func JOINMasterData() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Table(pkg.COMPANY + " as a").Joins("INNER JOIN " + pkg.DEPARTEMEN + " as b ON b.idCompany = a.idCompany").Joins("INNER JOIN " + pkg.DIVISION + " as c ON c.idDepartment = b.idDepartment")
	}
}

func WHEREMasterData(isActive int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("a.isActive = ?", 1)
	}
}

func SELECTMasterData() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select("a.nameCompany, b.nameDepartment, c.nameDivision")
	}
}
