package repositories

import (
	"dataons-service/models"
	"dataons-service/models/company"
	"dataons-service/models/queryScopes"
	"dataons-service/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type UserRepository struct {
	Mysql *gorm.DB
}

func StaticUserRepo(c *gin.Context) *UserRepository {
	return &UserRepository{
		Mysql: c.MustGet("mysql").(*gorm.DB),
	}
}

func DetailCompany(idCompany, idDepartment, idDivision int, db *gorm.DB) (interface{}, bool) {
	var masterCom []company.MasterCompany
	fmt.Println("id", idCompany, idDivision, idDivision)
	if idCompany != 0 && idDepartment == 0 && idDivision == 0 {
		db.Preload("Department").Preload("Department.Division").Where("idCompany = ?", idCompany).Find(&masterCom)
	} else if idCompany != 0 && idDepartment != 0 && idDivision == 0 {
		db.Preload("Department").Preload("Department", "idDepartment", idDepartment).Preload("Department.Division").Find(&masterCom)
	} else if idCompany != 0 && idDepartment != 0 && idDivision != 0 {
		db.Preload("Department", "idDepartment", idDepartment).Preload("Department.Division", "idDivision", idDivision).Preload("Department.Division.Employee").Find(&masterCom)
	}

	return masterCom, true
}

func (service UserRepository) MasterDataCompany(c *gin.Context) (interface{}, error) {

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	idCompany, _ := strconv.Atoi(c.Param("idCompany"))
	idDepartment, _ := strconv.Atoi(c.Param("idDepart"))
	idDivision, _ := strconv.Atoi(c.Param("idDivision"))

	if result, is := DetailCompany(idCompany, idDepartment, idDivision, service.Mysql); is {
		return result, nil
	}

	var totalData int64
	var data []models.MasterData
	var totalPage int

	switch {
	case pageSize > 100:
		pageSize = pageSize
	case pageSize <= 0:
		pageSize = 10
	}

	service.Mysql.Scopes(models.Paginate(pageSize, page), queryScopes.JOINMasterData(), queryScopes.SELECTMasterData()).Find(&data)

	service.Mysql.Scopes(queryScopes.JOINMasterData()).Count(&totalData)

	if int(totalData) < pageSize {
		totalPage = 1
	} else {
		totalPage = int(totalData) / pageSize
		if (int(totalData) % pageSize) != 0 {
			totalPage = totalPage + 1
		}
	}

	if page == 0 {
		page = 1
	}

	return models.Response{
		Data:       data,
		Page:       page,
		PageSize:   pageSize,
		TotalData:  totalData,
		TotalPages: int(totalData),
	}, nil
}

func (service UserRepository) MasterCompanyInheritance(c *gin.Context) (interface{}, error) {

	var masterCom []company.MasterCompany
	service.Mysql.Preload("Department").Preload("Department.Division").Preload("Department.Division.Employee").Find(&masterCom)

	return &masterCom, nil
}

func (service UserRepository) CreateUpdateCompany(input models.JSONCreateUpdate, c *gin.Context) (interface{}, error) {

	if input.IdCompany != 0 {
		tx := service.Mysql.Begin()
		if updated := tx.Table(pkg.COMPANY).Where("idCompany = ?", input.IdCompany).Updates(map[string]interface{}{
			"nameCompany": input.NameCompany,
			"isActive":    input.IsActive,
		}); updated.Error != nil {
			tx.Rollback()
			return map[string]interface{}{
				"craeted": false,
				"updated": false,
				"errors":  updated.Error,
			}, nil
		}
		tx.Commit()
		return map[string]interface{}{
			"craeted": false,
			"updated": true,
			"types":   "Company",
		}, nil
	}

	tx := service.Mysql.Begin()
	if craeted := tx.Create(&company.Company{
		NameCompany: input.NameCompany,
		IsActive:    1,
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
		Address:     input.Address,
	}); craeted.Error != nil {
		tx.Rollback()
		return map[string]interface{}{
			"craeted": false,
			"updated": false,
			"types":   craeted.Error,
		}, nil
	}
	tx.Commit()

	return map[string]interface{}{
		"craeted": true,
		"updated": false,
		"types":   "Company",
	}, nil
}

func (service UserRepository) DeleteCompany(id int, c *gin.Context) (interface{}, error) {

	tx := service.Mysql.Begin()
	tx.Delete(&company.Company{}, id)
	tx.Commit()

	return map[string]interface{}{
		"delete": true,
		"types":  "Company",
	}, nil
}
