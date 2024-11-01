package controllers

import (
	"dataons-service/models"
	"dataons-service/pkg/response"
	"dataons-service/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var UserController = &userController{}

type userController struct {
}

func (u *userController) MasterDataCompany(c *gin.Context) {
	app := response.Gin{C: c}

	result, err := repositories.StaticUserRepo(c).MasterDataCompany(c)
	if err != nil {
		app.Response(http.StatusBadRequest, "", err.Error(), nil)
		return
	}

	app.Response(http.StatusOK, "Succes", "", result)
	return
}

func (u *userController) MasterCompanyInheritance(c *gin.Context) {
	app := response.Gin{C: c}

	result, err := repositories.StaticUserRepo(c).MasterCompanyInheritance(c)
	if err != nil {
		app.Response(http.StatusBadRequest, "", err.Error(), nil)
		return
	}

	app.Response(http.StatusOK, "Succes", "", result)
	return
}

func (u *userController) CreateUpdateCompany(c *gin.Context) {
	app := response.Gin{C: c}

	var input models.JSONCreateUpdate
	if err := c.BindJSON(&input); err != nil {
		app.Response(http.StatusUnprocessableEntity, "", err.Error(), nil)
		return
	}

	result, err := repositories.StaticUserRepo(c).CreateUpdateCompany(input, c)
	if err != nil {
		app.Response(http.StatusBadRequest, "", "cannot generate access", nil)
		return
	}

	app.Response(http.StatusOK, "Succes", "", result)
	return
}

func (u *userController) DeleteCompany(c *gin.Context) {
	app := response.Gin{C: c}

	id, _ := strconv.Atoi(c.Query("id"))
	result, err := repositories.StaticUserRepo(c).DeleteCompany(id, c)
	if err != nil {
		app.Response(http.StatusBadRequest, "", "cannot generate access", nil)
		return
	}

	app.Response(http.StatusOK, "Succes", "", result)
	return
}

func (controller userController) GenerateCodeAccess(c *gin.Context) {
	app := response.Gin{C: c}
	result, err := repositories.StaticUserRepo(c).GenerateCodeAccess()
	if err != nil {
		app.Response(http.StatusBadRequest, "", "cannot generate access", nil)
		return
	}

	app.Response(http.StatusOK, "Succes", "", result)
	return
}
