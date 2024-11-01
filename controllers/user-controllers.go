package controllers

import (
	"dataons-service/pkg/response"
	"dataons-service/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

var UserController = &userController{}

type userController struct {
}

func (u *userController) MasterDataCompany(c *gin.Context) {
	app := response.Gin{C: c}

	result, err := repositories.StaticUserRepo(c).MasterDataCompany(c)
	if err != nil {
		app.Response(http.StatusBadRequest, "", "cannot generate access", nil)
		return
	}

	app.Response(http.StatusOK, "Succes", "", result)
	return
}
