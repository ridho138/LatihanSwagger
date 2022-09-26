package controllers

import (
	"net/http"
	"sesi8-latihan/server/models"
	"sesi8-latihan/server/views"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PersonController struct {
	db *gorm.DB
}

func NewPersonController(db *gorm.DB) *PersonController {
	return &PersonController{
		db: db,
	}
}

// GetPeople godoc
// @Summary Get all people
// @Decription Get list people
// @Tags person
// @Accept json
// @Produce json
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /people [get]
func (p *PersonController) GetPeople(ctx *gin.Context) {
	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: people,
	})
}

// AddPeople godoc
// @Summary Add new people
// @Decription Add new people
// @Tags person
// @Accept json
// @Produce json
// @Param data body models.Person true "Add New People"
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /people [post]
func (p *PersonController) AddPeople(ctx *gin.Context) {
	var people models.Person

	err := ctx.ShouldBindJSON(&people)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = p.db.Create(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: people,
	})
}

// EditPeople godoc
// @Summary Edit existing people
// @Decription Edit existing people
// @Tags person
// @Accept json
// @Produce json
// @Param data body models.Person true "Add New People"
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /people [put]
func (p *PersonController) EditPeople(ctx *gin.Context) {
	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: people,
	})
}

// DeletePeople godoc
// @Summary Delete existing people
// @Decription Delete existing people
// @Tags person
// @Accept json
// @Produce json
// @Param id path int true "People ID"
// @Success 200 {object} views.GetAllPeopleSwagger
// @Router /people/{id} [delete]
func (p *PersonController) DeletePeople(ctx *gin.Context) {
	var people []models.Person

	err := p.db.Find(&people).Error
	if err != nil {
		WriteJsonResponse(ctx, &views.Response{
			Status:  http.StatusInternalServerError,
			Message: "GET_PEOPLE_FAIL",
			Error:   err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, &views.Response{
		Status:  http.StatusOK,
		Message: "GET_PEOPLE_SUCCESS",
		Payload: people,
	})
}
