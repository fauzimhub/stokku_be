package controller

import (
	"net/http"
	"strconv"

	"github.com/fauzimhub/stokku/domain"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryUsecase domain.CategoryUsecase
}

func (cc *CategoryController) Create(c *gin.Context) {
	var category domain.Category

	err := c.ShouldBind(&category)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = cc.CategoryUsecase.Create(c, &category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Category created successfully",
	})
}

func (cc *CategoryController) Fetch(c *gin.Context) {

	categories, err := cc.CategoryUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (cc *CategoryController) GetByID(c *gin.Context) {
	categoryId := c.Param("id")
	id, err := strconv.Atoi(categoryId)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	category, err := cc.CategoryUsecase.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (cc *CategoryController) GetByName(c *gin.Context) {
	name := c.Param("name")

	category, err := cc.CategoryUsecase.GetByName(c, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}
