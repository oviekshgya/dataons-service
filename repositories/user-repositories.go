package repositories

import (
	"dataons-service/models"
	"dataons-service/models/queryScopes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type UserRepository struct {
	Mysql *gorm.DB
}

func StaticUserRepo(c *gin.Context) *UserRepository {
	return &UserRepository{
		Mysql: c.MustGet("mysql").(*gorm.DB),
	}
}

func (service UserRepository) MasterDataCompany(c *gin.Context) (interface{}, error) {

	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
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

	service.Mysql.Scopes(queryScopes.JOINMasterData(), queryScopes.SELECTMasterData()).Count(&totalData)

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
